package playok

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/lk16/dots/internal/othello"
	"github.com/pkg/errors"
)

// handleMessage processes any incoming Message
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
		1:  ignore, // ping?
		18: bot.handleConnectMessage,
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
		87: todo,
		88: todo,
		89: bot.handleTableSettingsUpdate,
		90: bot.handleCurrentTableUpdate,
		91: bot.handleTableBoardReset,
		92: bot.handleTableBoardUpdate,
	}

	log.Printf("RECV [%d] %s", message.I[0], messageBytes)

	handler, ok := messageHandlers[message.I[0]]
	if !ok {
		log.Printf("WARN [%d] unknown %s", message.I[0], messageBytes)
		return nil
	}

	return handler(message)
}

func (bot *Bot) handleRating(message Message) error {
	bot.playok.Lock()
	defer bot.playok.Unlock()

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

	_ = iSlice[0]        // TODO
	tableID := iSlice[1] // TODO viewing table
	rating := iSlice[2]

	playerName := s

	bot.playok.Lock()
	defer bot.playok.Unlock()

	// this informs us we have left a table and are in the lobby
	if playerName == bot.playok.userName && tableID == 0 {
		bot.playok.currentTable = currentTable{}
	}

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
	bot.playok.Lock()
	defer bot.playok.Unlock()

	return bot.upsertTable(message.I[1:], message.S)
}

func (bot *Bot) handleJoinTable(message Message) error {
	bot.playok.Lock()
	defer bot.playok.Unlock()

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
	minRatingID := iSlice[1] // TODO rating requirement
	_ = iSlice[2]            // TODO 1 if first seat is taken, 0 if not, 2 if abandonned?
	_ = iSlice[3]            // TODO 1 if second seat is taken, 0 if not, 2 if abandonned?

	tableSettings := sSlice[0]
	players := [2]string{
		sSlice[1],
		sSlice[2],
	}

	table := table{
		players:     players,
		rated:       true,
		minRatingID: minRatingID,
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

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if _, ok := bot.playok.players[username]; !ok {
		errors.New("received logout for player that's already logged out")
	}

	delete(bot.playok.players, username)
	return nil
}

func (bot *Bot) handleTableClose(message Message) error {
	tableID := message.I[1]

	bot.playok.Lock()
	defer bot.playok.Unlock()

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

	bot.playok.Lock()
	defer bot.playok.Unlock()

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

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	bot.playok.currentTable.viewers = message.S
	return nil
}

func (bot *Bot) handleTableChatMessage(message Message) error {

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	log.Printf("INFO Chat message: %s", message.S[0])
	return nil
}

func (bot *Bot) handleAlertMessage(message Message) error {

	log.Printf("INFO Alert message: %s", message.S[0])
	return nil
}

func (bot *Bot) handleBootedFromTable(message Message) error {

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	// reset currentTable field
	bot.playok.currentTable = currentTable{}
	return nil
}

func (bot *Bot) handleTableBoardReset(message Message) error {

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	bot.playok.currentTable.board = *othello.NewBoardWithTurn()

	for _, value := range message.I[2:] {
		if err := bot.handleBoardUpdate(value); err != nil {
			return err
		}
	}

	return nil
}

func (bot *Bot) handleTableBoardUpdate(message Message) error {

	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	value := message.I[2]
	return bot.handleBoardUpdate(value)
}

func (bot *Bot) handleBoardUpdate(value int) error {
	if value < -1 || value > 127 {
		return errorf("unexpected move value %d", value)
	}

	if value == -1 {
		bot.playok.currentTable.board = *othello.NewBoardWithTurn()
		return nil
	}
	if value >= 64 {
		value -= 64
	}

	board := bot.playok.currentTable.board

	flipped := board.DoMove(othello.BitSet(1) << uint(value))
	if flipped == 0 {
		log.Printf("DEBG board=%s\nfailed at value=%d", board.String(), value)
		return errors.New("no discs flipped")
	}

	bot.playok.currentTable.board = board
	return nil
}

func (bot *Bot) handleCurrentTableUpdate(message Message) error {
	bot.playok.Lock()
	defer bot.playok.Unlock()

	if message.I[1] != bot.playok.currentTable.ID {
		return errors.New("received table update for unexpected table")
	}

	bot.playok.currentTable.playerToMove = message.I[3]
	return nil
}

func (bot *Bot) handleConnectMessage(message Message) error {

	bot.playok.Lock()
	defer bot.playok.Unlock()

	bot.playok.userName = message.S[0]

	_ = message.I[1] // TODO
	_ = message.I[2] // TODO

	return nil
}
