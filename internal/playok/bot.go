package playok

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lk16/dots/internal/othello"
	"github.com/pkg/errors"
)

const (
	playokURL          = "https://www.playok.com/"
	playokLoginURL     = "https://www.playok.com/en/login.phtml"
	playokReversiURL   = "https://www.playok.com/en/reversi/"
	playokWebsocketURL = "wss://x.playok.com:17003/ws/"
	userAgent          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
)

var (
	playokParsedURL *url.URL
	windowApRegex   = regexp.MustCompile("window.ap = (.*);")
	windowGeRegex   = regexp.MustCompile("window.ge = (.*);")
)

func init() {
	var err error
	playokParsedURL, err = url.Parse(playokURL)
	if err != nil {
		panic(err)
	}
}

// errorf returns a formated and stack-annotated error
// it is a replacement for errorf(), which does not annotate the stack trace
func errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

// Bot contains the state of an automated player on playok.com
type Bot struct {
	userName  string
	password  string
	windowAp  string
	windowGe  string
	browser   *http.Client
	websocket *websocket.Conn
	playok    *state
}

// NewBot initializes a new bot
func NewBot(userName, password string) *Bot {

	// do not follow redirects
	redirectHandler := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// skip checking error, this can't go wrong
	cookieJar, _ := cookiejar.New(nil)

	return &Bot{
		userName: userName,
		password: password,
		browser: &http.Client{
			CheckRedirect: redirectHandler,
			Jar:           cookieJar,
		},
		playok: newState(),
	}
}

func checkCookies(jar http.CookieJar, expectedCookies []string) error {
	cookies := jar.Cookies(playokParsedURL)

	for _, cookieName := range expectedCookies {
		var found bool
		for _, cookie := range cookies {
			if cookie.Name == cookieName {
				found = true
				break
			}
		}
		if !found {
			return errorf("missing cookie with name %s", cookieName)
		}
	}

	if len(cookies) > len(expectedCookies) {
		log.Printf("warning: received %d cookies while expecing %d cookies", len(cookies), len(expectedCookies))
	}

	return nil
}

func (bot *Bot) login() error {

	if bot.userName == "" || bot.password == "" {
		return errors.New("username and/or password are not set")
	}

	formValues := url.Values{
		"username": {bot.userName},
		"pw":       {bot.password},
		"cc":       {"0"},
	}
	request, err := http.NewRequest("POST", playokLoginURL,
		strings.NewReader(formValues.Encode()))

	if err != nil {
		return errors.Wrap(err, "building request failed")
	}
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := bot.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusFound {
		return errors.New("unexpected status code")
	}

	expectedCookies := []string{"ku", "ksession"}

	if err := checkCookies(bot.browser.Jar, expectedCookies); err != nil {
		return err
	}

	return nil
}

func (bot *Bot) visitReversiPage() error {
	request, err := http.NewRequest("GET", playokReversiURL, nil)
	if err != nil {
		return errors.Wrap(err, "building request failed")
	}
	request.Header.Add("User-Agent", userAgent)

	// add cookie with constant value
	newCookies := []*http.Cookie{
		&http.Cookie{
			Name:  "kbeta",
			Value: "rv",
		},
	}

	bot.browser.Jar.SetCookies(playokParsedURL, newCookies)

	response, err := bot.browser.Do(request)
	if err != nil {
		return errors.Wrap(err, "sending request failed")
	}

	defer response.Body.Close()

	expectedCookies := []string{"ku", "ksession", "kbeta", "kbexp", "kt"}

	if err := checkCookies(bot.browser.Jar, expectedCookies); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	windowAp := windowApRegex.FindSubmatch(body)
	windowGe := windowGeRegex.FindSubmatch(body)

	if windowAp == nil || windowGe == nil {
		return errors.New("failed to regex match js vars")
	}

	bot.windowAp = string(windowAp[1])
	bot.windowGe = string(windowGe[1])

	return nil
}

// takeAction gives the bot to take initiative, such as joining tables or doing moves
func (bot *Bot) takeAction() error {

	if bot.playok.currentTable.ID != 0 {
		log.Printf("Bot viewing table, no action")
		return nil
	}

	if len(bot.playok.tables) == 0 {
		log.Printf("Bot has not received list of tables (yet), no action")
		return nil
	}

	for tableID, table := range bot.playok.tables {
		if table.players[0] == "" || table.players[1] == "" {
			continue
		}

		if err := bot.sendJoinTableRequest(tableID); err != nil {
			return errors.Wrap(err, "joining table request failed")
		}
	}

	return nil
}

func (bot *Bot) connectWebSocket() error {

	headers := make(http.Header)
	headers.Add("Origin", "https://www.playok.com")
	headers.Add("User-Agent", userAgent)

	dialer := *websocket.DefaultDialer
	dialer.Jar = bot.browser.Jar
	conn, _, err := dialer.Dial(playokWebsocketURL, headers)

	if err != nil {
		return errors.Wrap(err, "connecting failed")
	}

	bot.websocket = conn
	return nil
}

func (bot *Bot) getInitMessage() Message {

	var kSessionCookie string
	for _, cookie := range bot.browser.Jar.Cookies(playokParsedURL) {
		if cookie.Name == "ksession" {
			kSessionCookie = cookie.Value
		}
	}

	if kSessionCookie == "" {
		log.Printf("warning: could not find ksession cookie")
	}

	splitKSession := strings.Split(kSessionCookie, ":")

	firstArg := fmt.Sprintf("%s+|%s|%s",
		splitKSession[0], // part of a cookie
		bot.windowAp,     // scraped JS value window.ap
		bot.windowGe,     // scraped JS value window.ge
	)

	nowMilli := time.Now().UnixNano() / 1000000

	return Message{
		I: []int{1713},
		S: []string{
			firstArg,                                 // see above
			"en",                                     // language
			"b",                                      // ???
			"",                                       // ???
			userAgent,                                // user-agent
			fmt.Sprintf("/%d/1", nowMilli),           // timestamp
			"w",                                      // ???
			"2560x1440 1",                            // screen size
			"ref:https://www.playok.com/en/reversi/", // referer
			"ver:233",                                // client version (TODO scrape this)
		},
	}
}

func (bot *Bot) loop() error {

	err := bot.sendMessage(bot.getInitMessage())
	if err != nil {
		return errors.Wrap(err, "sending init message failed")
	}

	printStateTicker := time.NewTicker(500 * time.Millisecond)
	keepAliveTicker := time.NewTicker(30 * time.Second)
	botActionTicker := time.NewTicker(time.Second)
	messageChan := make(chan []byte)

	go func() {
		for {
			messageType, messageBytes, err := bot.websocket.ReadMessage()
			if err != nil {
				panic("read message error:" + err.Error())
			}

			if messageType != websocket.TextMessage {
				log.Printf("RECV ignoring message with type %d: %s",
					messageType, string(messageBytes))
			}

			messageChan <- messageBytes
		}
	}()

	for {
		select {
		case <-printStateTicker.C:
			bot.printState()
		case <-keepAliveTicker.C:
			if err = bot.sendKeepAlive(); err != nil {
				return errors.Wrap(err, "sending keep alive failed")
			}
		case <-botActionTicker.C:
			if err = bot.takeAction(); err != nil {
				return errors.Wrap(err, "bot action failed")
			}
		case messageBytes := <-messageChan:
			if err = bot.readMessage(messageBytes); err != nil {
				return errors.Wrap(err, "message handling error")
			}
		}
	}
}

func (bot *Bot) handleConnect(message Message) error {
	bot.playok.userName = message.S[0]
	if bot.playok.userName != bot.userName {
		return errors.New("received unexpected username from server")
	}

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	return nil
}

func (bot *Bot) sendMessage(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = bot.websocket.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) sendKeepAlive() error {

	messages := []Message{
		Message{I: []int{2}},
		Message{I: []int{}},
	}

	for _, message := range messages {
		if err := bot.sendMessage(message); err != nil {
			return err
		}
	}

	return nil
}

func (bot *Bot) sendMove(moveID int) error {
	message := Message{
		I: []int{
			92,
			bot.playok.currentTable.ID,
			2, // hardcoded as 2 in JS. wtf?
			moveID,
			0, // TODO
		},
	}

	return bot.sendMessage(message)
}

func (bot *Bot) sendJoinTableRequest(tableID int) error {
	message := Message{
		I: []int{72, tableID},
	}
	return bot.sendMessage(message)
}

func (bot *Bot) handleRating(message Message) error {
	bot.playok.rating = message.I[1]
	return nil
}

func (bot *Bot) handleInitUsers(message Message) error {
	expectedIlength := len(message.S)*3 + 3
	if len(message.I) != expectedIlength {
		return errorf("expected len(message.I)==%d, got %d", expectedIlength, len(message.I))
	}

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	for i := range message.S {
		offset := 3 + 3*i
		if err := bot.upsertUser(message.I[offset:offset+3], message.S[i]); err != nil {
			return err
		}
	}

	return nil
}

func (bot *Bot) handleUpdateUser(message Message) error {
	return bot.upsertUser(message.I[1:], message.S[0])
}

func (bot *Bot) upsertUser(iSlice []int, s string) error {

	if len(iSlice) != 3 {
		return errorf("len(iSlice)==%d, expected 3", len(iSlice))
	}

	_ = iSlice[0] // TODO
	_ = iSlice[1] // TODO viewing table
	rating := iSlice[2]

	playerName := s

	bot.playok.players[playerName] = player{
		rating: rating,
	}

	return nil
}

func (bot *Bot) handleInitTables(message Message) error {
	roomCount := len(message.S) / 3
	expectedIlength := 4*roomCount + 3
	if len(message.I) != expectedIlength {
		return errorf("expected len(message.I)==%d, got %d", expectedIlength, len(message.I))
	}

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	for i := 0; i < roomCount; i++ {
		IOffset := 4*i + 3
		SOffset := 3 * i

		if err := bot.upsertTable(message.I[IOffset:IOffset+4], message.S[SOffset:SOffset+3]); err != nil {
			return err
		}
	}

	return nil
}

func (bot *Bot) handleUpdateTable(message Message) error {
	return bot.upsertTable(message.I[1:], message.S)
}

func (bot *Bot) handleJoinTable(message Message) error {
	bot.playok.currentTable.ID = message.I[1]
	return bot.upsertTable(message.I[1:], message.S)
}

func (bot *Bot) upsertTable(iSlice []int, sSlice []string) error {

	if len(iSlice) != 4 {
		return errorf("len(iSlice)==%d, expected 4", len(iSlice))
	}

	if len(sSlice) != 3 {
		return errorf("len(sSlice)==%d, expected 4", len(sSlice))
	}

	tableID := iSlice[0]
	_ = iSlice[1] // TODO
	_ = iSlice[2] // TODO
	_ = iSlice[3] // TODO

	tableSettings := sSlice[0]
	players := [2]string{
		sSlice[1],
		sSlice[2],
	}

	table := table{
		players: players,
		rated:   true,
	}

	if tableSettings != "" {
		for _, setting := range strings.Split(tableSettings, ", ") {

			settingError := errorf("unexpected table setting \"%s\"", setting)

			switch setting {
			case "x":
				table.rated = false
			case "xot":
				table.xot = true
			case "":
				return settingError
			default:
				timeLimit, err := strconv.Atoi(setting[:len(setting)-1])
				if err != nil {
					return settingError
				}
				table.timeLimit = timeLimit
			}
		}
	}

	bot.playok.tables[tableID] = table

	return nil
}

func (bot *Bot) handlePlayerLogout(message Message) error {
	username := message.S[0]

	if _, ok := bot.playok.players[username]; !ok {
		errors.New("received logout for player that's already logged out")
	}

	delete(bot.playok.players, username)
	return nil
}

func (bot *Bot) handleTableClose(message Message) error {
	tableID := message.I[1]

	if _, ok := bot.playok.tables[tableID]; !ok {
		return errors.New("closing table that's already closed")
	}

	delete(bot.playok.tables, tableID)
	return nil
}

func (bot *Bot) handleTableSettingsUpdate(message Message) error {

	expectedSlength := len(message.I) - 2
	if len(message.S) != expectedSlength {
		return errorf("expected len(message.S)==%d, got %d", expectedSlength, len(message.S))
	}

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	currentTable := bot.playok.currentTable

	for i, setting := range message.S {

		value := message.I[i+2]

		switch setting {
		case "ttype":
			// TODO minimum rating
		case "gtype":
			currentTable.rated = (value == 1)
		case "tm":
		case "tg":
			currentTable.timeLimit = value
		case "ud":
			currentTable.allowUndo = (value == 1)
		case "xot":
			currentTable.xot = (value == 1)
		case "tch":
			// TODO ???
		default:
			if !strings.HasPrefix(setting, "op:") {
				return errorf("unexpected table setting \"%s\"", setting)
			}
			currentTable.op = setting[3:]
		}
	}

	bot.playok.currentTable = currentTable
	return nil
}

func (bot *Bot) handleTableViewersUpdate(message Message) error {

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	bot.playok.currentTable.viewers = message.S
	return nil
}

func (bot *Bot) handleTableChatMessage(message Message) error {

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	log.Printf("Chat message: %s", message.S[0])
	return nil
}

func (bot *Bot) handleAlertMessage(message Message) error {

	log.Printf("Alert message: %s", message.S[0])
	return nil
}

func (bot *Bot) handleBootedFromTable(message Message) error {

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	// reset currentTable field
	bot.playok.currentTable = currentTable{}
	return nil
}

func (bot *Bot) handleTableBoardReset(message Message) error {

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	board := *othello.NewBoard()

	for _, value := range message.I[2:] {
		if value < -1 || value > 127 {
			return errorf("unexpected move value %d", value)
		}

		if value == -1 {
			continue
		}
		if value > 64 {
			value -= 64
		}

		flipped := board.DoMove(1 << value)
		if flipped == 0 {
			return errors.New("no discs flipped")
		}
	}

	bot.playok.currentTable.board = board
	return nil
}

func (bot *Bot) handleTableBoardUpdate(message Message) error {

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	value := message.I[2]

	if value < 0 || value > 127 {
		return errorf("unexpected move value %d", value)
	}

	if value > 64 {
		value -= 64
	}

	flipped := bot.playok.currentTable.board.DoMove(1 << value)
	if flipped == 0 {
		return errors.New("no discs flipped")
	}

	return nil
}

// readMessage blocks until a websocket message is received
// it updates the bot.playok state
func (bot *Bot) readMessage(messageBytes []byte) error {

	var message Message
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		return err
	}

	ignore := func(Message) error {
		return nil
	}

	todo := func(message Message) error {
		log.Printf("TODO [%d] %s", message.I[0], messageBytes)
		return nil
	}

	messageHandlers := map[int]func(Message) error{
		1:  ignore, // ping?
		18: bot.handleConnect,
		20: ignore, // frontend stuff
		22: ignore, // role? guest/player?
		23: ignore, // unused options
		24: bot.handlePlayerLogout,
		25: bot.handleUpdateUser,
		27: bot.handleInitUsers,
		28: ignore, // friends
		30: ignore, // frontend stuff
		31: ignore, // frontend stuff
		32: ignore, // frontend stuff: room list
		33: bot.handleRating,
		51: ignore, // frontend translations
		52: bot.handleAlertMessage,
		70: bot.handleUpdateTable,
		71: bot.handleInitTables,
		72: bot.handleTableClose,
		73: bot.handleJoinTable,
		74: bot.handleBootedFromTable,
		81: bot.handleTableChatMessage,
		84: ignore, // table join messsage
		86: bot.handleTableViewersUpdate,
		88: todo,
		89: bot.handleTableSettingsUpdate,
		90: todo,
		91: bot.handleTableBoardReset,
		92: bot.handleTableBoardUpdate,
	}

	handler, ok := messageHandlers[message.I[0]]
	if !ok {
		log.Printf("WARN [%d] unknown %s", message.I[0], messageBytes)
		return nil
	}

	return handler(message)
}

// Run is the entrypoint of the Bot
func (bot *Bot) Run() error {

	if err := bot.login(); err != nil {
		return errors.Wrap(err, "login failed")
	}

	if err := bot.visitReversiPage(); err != nil {
		return errors.Wrap(err, "visit reversi page failed")
	}

	if err := bot.connectWebSocket(); err != nil {
		return errors.Wrap(err, "connecting to websocket failed")
	}

	if err := bot.loop(); err != nil {
		return errors.Wrap(err, "websocket loop failed")
	}

	return nil
}

// printState prints the bot state in ascii art
func (bot *Bot) printState() {

	var tableIDs []int
	for ID := range bot.playok.tables {
		tableIDs = append(tableIDs, ID)
	}

	sort.Ints(tableIDs)

	var buff bytes.Buffer

	for _, ID := range tableIDs {
		table := bot.playok.tables[ID]

		var xot string
		if table.xot {
			xot = "xot"
		}

		var unrated string
		if !table.rated {
			unrated = "x"
		}

		buff.WriteString(fmt.Sprintf("%3s %1s %2dm %5d% 15s %15s\n",
			xot,
			unrated,
			table.timeLimit,
			ID,
			table.players[0],
			table.players[1],
		))
	}

	buff.WriteString("---\n")

	buff.WriteString(fmt.Sprintf("%8s:%10s\n", "username", bot.playok.userName))
	buff.WriteString(fmt.Sprintf("%8s:%10d\n", "rating", bot.playok.rating))
	buff.WriteString("---\n")

	currentTable := bot.playok.currentTable
	buff.WriteString(fmt.Sprintf("%8s:%10d\n", "table", currentTable.ID))
	if currentTable.ID != 0 {
		buff.WriteString(currentTable.board.String())
	}
	buff.WriteString("---\n")
	buff.WriteString(fmt.Sprintf("%8s:%10d\n", "players", len(bot.playok.players)))
	buff.WriteString(fmt.Sprintf("%8s:%10d\n", "tables", len(bot.playok.tables)))

	buff.WriteString("---\n")

	fmt.Print(buff.String())
}
