package playok

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/websocket"
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
			return fmt.Errorf("missing cookie with name %s", cookieName)
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

func (bot *Bot) getInitMessage() *Message {

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

	message := &Message{
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

	return message
}

func (bot *Bot) websocketLoop() error {

	initMessage := bot.getInitMessage()
	initMessageBytes, _ := json.Marshal(initMessage)
	err := bot.websocket.WriteMessage(websocket.TextMessage, initMessageBytes)
	if err != nil {
		return errors.Wrap(err, "write json error")
	}

	for {
		messageType, bytes, err := bot.websocket.ReadMessage()
		if err != nil {
			return errors.Wrap(err, "read message error")
		}

		if messageType != websocket.TextMessage {
			log.Printf("RECV ignoring message with type %d: %s", messageType, string(bytes))
		}

		if err = bot.handleMessage(bytes); err != nil {
			return errors.Wrap(err, "message handling error")
		}

		fmt.Printf("----\n%s", bot.playok.String())
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

func (bot *Bot) handleRating(message Message) error {
	bot.playok.rating = message.I[1]
	return nil
}

func (bot *Bot) handleInitUsers(message Message) error {
	expectedIlength := len(message.S)*3 + 3
	if len(message.I) != expectedIlength {
		return fmt.Errorf("expected len(message.I)==%d, got %d", expectedIlength, len(message.I))
	}

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	for i, playerName := range message.S {
		offset := 3 + 3*i

		_ = message.I[offset+0] // TODO
		_ = message.I[offset+1] // TODO
		rating := message.I[offset+2]

		bot.playok.players[playerName] = player{
			rating: rating,
		}
	}

	return nil
}

func (bot *Bot) handleUpdateUser(message Message) error {

	_ = message.I[0] // TODO
	_ = message.I[1] // TODO
	rating := message.I[2]

	playerName := message.S[0]

	bot.playok.players[playerName] = player{
		rating: rating,
	}

	return nil
}

func (bot *Bot) handleInitTables(message Message) error {
	roomCount := len(message.S) / 3
	expectedIlength := 4*roomCount + 3
	if len(message.I) != expectedIlength {
		return fmt.Errorf("expected len(message.I)==%d, got %d", expectedIlength, len(message.I))
	}

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	for i := 0; i < roomCount; i++ {
		IOffset := 4*i + 3
		SOffset := 3 * i

		tableID := message.I[IOffset]
		_ = message.I[IOffset+1] // TODO
		_ = message.I[IOffset+2] // TODO
		_ = message.I[IOffset+3] // TODO

		rules := message.S[SOffset]
		players := [2]string{
			message.S[SOffset+1],
			message.S[SOffset+2],
		}

		bot.playok.tables[tableID] = table{
			rules:   rules,
			players: players,
		}
	}

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

func (bot *Bot) handleUpdateTable(message Message) error {

	tableID := message.I[1]
	_ = message.I[2] // TODO
	_ = message.I[3] // TODO
	_ = message.I[4] // TODO

	rules := message.S[0]
	players := [2]string{
		message.S[1],
		message.S[2],
	}

	bot.playok.tables[tableID] = table{
		rules:   rules,
		players: players,
	}

	return nil
}

func (bot *Bot) handleMessage(messageBytes []byte) error {

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
		1:  todo,
		18: bot.handleConnect,
		20: ignore, // frontend stuff
		22: todo,
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
		70: bot.handleUpdateTable,
		71: bot.handleInitTables,
		72: todo,
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

	if err := bot.websocketLoop(); err != nil {
		return errors.Wrap(err, "websocket loop failed")
	}

	return nil
}
