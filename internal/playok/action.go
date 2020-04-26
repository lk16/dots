package playok

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/lk16/dots/internal/treesearch"
	"github.com/pkg/errors"
)

var (
	errWrongTable        = errors.New("bot is not at the expected table")
	errNotPlaying        = errors.New("bot is only viewing the table, not playing")
	errNotEnoughPlayers  = errors.New("not enough players")
	errGameEnded         = errors.New("the game ended")
	errMaxRetriesReached = errors.New("max retries reached")
)

type actionState int

const (
	stateAwaitTablesList actionState = iota
	stateAwaitFindTable
	stateAwaitFindSeat
	stateAwaitStartGame
	stateAwaitLeaveTable
	stateAwaitTurn
	stateStop
)

func info(format string, args ...interface{}) {
	log.Printf("INFO BOT: "+format, args...)
}

// takeAction gives the bot to take initiative, such as joining tables or doing moves
func (bot *Bot) takeAction() error {

	var tableID int
	var seatID int

	othelloBot := treesearch.NewBot(log.Writer(), 12, 18,
		treesearch.NewMtdf(
			treesearch.NewMemoryCache(),
			treesearch.FastHeuristic))

	state := stateAwaitTablesList

	for {
		time.Sleep(200 * time.Millisecond)

		switch state {

		case stateAwaitTablesList:

			if bot.checkTablesList() {

				bot.playok.RLock()
				tableID = bot.playok.currentTable.ID
				bot.playok.RUnlock()

				if tableID != 0 {
					state = stateAwaitTurn
					break
				}
				state = stateAwaitFindTable
			}

		case stateAwaitLeaveTable:
			bot.sendLeaveTableRequest(tableID)
			if err := bot.awaitJoinLobby(tableID); err != nil {
				info("joining lobby failed: %s", tableID, err.Error())
				break
			}
			state = stateAwaitFindTable

		case stateAwaitFindTable:
			if bot.shutdownSoon {
				state = stateStop
				break
			}

			tableID = bot.awaitFindOnePlayerTable()
			info("table %d has one player", tableID)
			bot.sendJoinTableRequest(tableID)

			if err := bot.awaitJoinTable(tableID); err != nil {
				info("joining table %d failed: %s", tableID, err.Error())
				break
			}
			state = stateAwaitFindSeat

		case stateAwaitFindSeat:
			var ok bool
			if seatID, ok = bot.findEmptySeat(tableID); !ok {
				info("could not find empty seat")
				state = stateAwaitLeaveTable
				break
			}
			bot.sendTakeSeatRequest(tableID, seatID)

			if err := bot.awaitTakeSeat(tableID); err != nil {
				info("taking seat failed: %s", err.Error())
				state = stateAwaitLeaveTable
				break
			}
			state = stateAwaitStartGame

		case stateAwaitStartGame:
			if bot.shutdownSoon {
				state = stateStop
				break
			}
			bot.sendStartGameRequest(tableID)
			if err := bot.awaitStartGame(tableID); err != nil {
				state = stateAwaitLeaveTable
				break
			}
			state = stateAwaitTurn

		case stateAwaitTurn:
			if err := bot.awaitTurn(tableID, seatID); err != nil {
				if err == errGameEnded {
					state = stateAwaitStartGame
					break
				}
				return errors.Wrap(err, "waiting for turn failed")
			}

			bot.playok.RLock()
			info("\n%s", bot.playok.currentTable.board.String())
			bot.playok.RUnlock()

			info("bot is to move")

			discCount, err := bot.computeAndSendMove(othelloBot)
			if err != nil {
				return errors.Wrap(err, "compute and send move failed")
			}
			info("bot sent move")
			err = bot.awaitMoveConfirmation(discCount)
			if err != nil {
				info("waiting for move confirmation failed: %s", err.Error())
				break
				// TODO do not recompute move on retry
			}

			bot.playok.RLock()
			info("\n%s", bot.playok.currentTable.board.String())
			bot.playok.RUnlock()

		case stateStop:
			info("shutting down")
			// TODO this is ugly
			os.Exit(0)

		default:
			return errorf("Unhandled state value: %d", state)
		}
	}
}

// awaitMoveConfirmation waits until the server tells us the move is received
// we get an update of our own move for some reason
func (bot *Bot) awaitMoveConfirmation(discCount int) error {
	retries := 20
	for {
		bot.playok.RLock()

		table := bot.playok.currentTable
		success := table.board.CountDiscs() == discCount

		bot.playok.RUnlock()

		if success {
			return nil
		}

		retries--
		if retries == 0 {
			return errors.New("max retries reached")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (bot *Bot) computeAndSendMove(othelloBot *treesearch.Bot) (int, error) {

	info("computing move")

	bot.playok.RLock()
	board := bot.playok.currentTable.board
	bot.playok.RUnlock()

	move, err := othelloBot.DoMove(board.Board)
	if err != nil {
		return 0, errors.Wrap(err, "bot failed to compute move")
	}

	moveID, ok := board.GetMoveField(*move)
	if !ok {
		return 0, errors.New("GetMoveField failed")
	}

	randomDelay := (time.Duration(300+rand.Intn(300)+rand.Intn(300)) * time.Millisecond)
	info("delaying sending move %dms", randomDelay.Milliseconds())
	time.Sleep(randomDelay)

	info("sending move")
	if err := bot.sendMove(moveID); err != nil {
		return 0, errors.Wrap(err, "failed to send move")
	}

	stats := othelloBot.LifetimeStats
	info("total bot stats: %s", stats)

	return move.CountDiscs(), nil
}

func (bot *Bot) awaitTurn(tableID, seatID int) error {

	for {
		bot.playok.RLock()

		table := bot.playok.currentTable
		stillOnSameTable := table.ID == tableID
		turn := table.playerToMove

		bot.playok.RUnlock()

		if !stillOnSameTable {
			return errWrongTable
		}

		if turn == -1 {
			return errGameEnded
		}

		if turn == seatID {
			return nil
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (bot *Bot) awaitStartGame(tableID int) error {

	retries := 10
	for {
		info("await start game: %d retries left", retries)

		bot.playok.RLock()

		table := bot.playok.tables[tableID]
		stillOnSameTable := bot.playok.currentTable.ID == tableID

		botHasSeat := false
		for _, player := range bot.playok.tables[tableID].players {
			if player == bot.playok.userName {
				botHasSeat = true
			}
		}

		playerCount := table.countPlayers()
		gameHasStarted := bot.playok.currentTable.playerToMove != -1

		bot.playok.RUnlock()

		if !stillOnSameTable {
			return errWrongTable
		}

		if !botHasSeat {
			return errNotPlaying
		}

		if playerCount != 2 {
			return errNotEnoughPlayers
		}

		if gameHasStarted {
			return nil
		}

		retries--
		if retries == 0 {
			return errMaxRetriesReached
		}
		time.Sleep(time.Second)
	}
}

func (bot *Bot) findEmptySeat(tableID int) (int, bool) {
	info("finding empty seat at table %d", tableID)

	bot.playok.RLock()
	defer bot.playok.RUnlock()

	table := bot.playok.tables[tableID]
	for seatID, player := range table.players {
		if player == "" {
			return seatID, true
		}
	}
	return 0, false
}

func (bot *Bot) checkTablesList() bool {

	bot.playok.RLock()
	success := len(bot.playok.tables) != 0
	bot.playok.RUnlock()

	return success
}

// awaitFindOnePlayerTable blocks until a table with one player is found
func (bot *Bot) awaitFindOnePlayerTable() int {
	for {
		info("await find one player table")

		foundTableID := 0

		bot.playok.RLock()
		tableIDs := bot.playok.getShuffledTableIDs()

		for _, ID := range tableIDs {
			table := bot.playok.tables[ID]
			if table.rated && table.timeLimit > 1 && table.countPlayers() == 1 {
				foundTableID = ID
				break
			}
		}
		bot.playok.RUnlock()

		if foundTableID != 0 {
			return foundTableID
		}

		time.Sleep(time.Second)
	}
}

func (bot *Bot) awaitJoinLobby(tableID int) error {
	return bot.awaitJoinTable(0)
}

func (bot *Bot) awaitJoinTable(tableID int) error {

	retries := 5
	for {
		info("await join table: %d retries left", retries)

		bot.playok.RLock()
		success := bot.playok.currentTable.ID == tableID
		bot.playok.RUnlock()

		if success {
			return nil
		}

		retries--
		if retries == 0 {
			return errMaxRetriesReached
		}
		time.Sleep(time.Second)
	}
}

func (bot *Bot) awaitTakeSeat(tableID int) error {

	retries := 5
	for {
		info("await take empty seat: %d retries left", retries)

		bot.playok.RLock()
		stillOnSameTable := bot.playok.currentTable.ID == tableID
		tookEmptySeat := false

		table, ok := bot.playok.tables[tableID]
		if ok {
			for _, player := range table.players {
				if player == bot.playok.userName {
					tookEmptySeat = true
				}
			}
		} else {
			// table is gone somehow
			stillOnSameTable = false
		}
		bot.playok.RUnlock()

		if !stillOnSameTable {
			return errWrongTable
		}

		if tookEmptySeat {
			return nil
		}

		retries--
		if retries == 0 {
			return errMaxRetriesReached
		}
		time.Sleep(time.Second)
	}
}
