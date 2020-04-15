package playok

import (
	"log"
	"math/rand"
	"time"

	"github.com/lk16/dots/internal/treesearch"
	"github.com/pkg/errors"
)

var (
	errWrongTable       = errors.New("bot is not at the expected table")
	errNotPlaying       = errors.New("bot is only viewing the table, not playing")
	errNotEnoughPlayers = errors.New("not enough players")
	errGameEnded        = errors.New("the game ended")
)

func info(format string, args ...interface{}) {
	log.Printf("INFO BOT: "+format, args...)
}

// takeAction gives the bot to take initiative, such as joining tables or doing moves
func (bot *Bot) takeAction() error {

	if err := bot.awaitTablesList(); err != nil {
		return errors.Wrap(err, "await tables list failed")
	}

	tableID := bot.awaitFindOnePlayerTable()
	info("table %d has one player", tableID)

	if err := bot.sendJoinTableRequest(tableID); err != nil {
		return errors.Wrap(err, "failed sending join table request")
	}

	if err := bot.awaitJoinTable(tableID); err != nil {
		return errors.Wrap(err, "joining table failed")
	}
	info("joined table %d", tableID)

	// server doesn't seem to keep up?
	time.Sleep(time.Second)

	seatID, ok := bot.findEmptySeat(tableID)
	if !ok {
		// TODO leave table and try again
		return errorf("could not find empty seat")
	}

	if err := bot.sendTakeSeatRequest(tableID, seatID); err != nil {
		return errors.Wrap(err, "failed sending take seat request")
	}

	if err := bot.awaitTakeSeat(tableID); err != nil {
		// TODO leave table and try again
		return errors.Wrap(err, "taking seat failed")
	}

	for {

		if err := bot.sendStartGameRequest(tableID); err != nil {
			return errors.Wrap(err, "failed sending start game request")
		}

		if err := bot.awaitStartGame(tableID); err != nil {
			// TODO leave table and try again
			return errors.Wrap(err, "starting game failed")
		}

		for {
			if err := bot.awaitTurn(tableID, seatID); err != nil {
				if err == errGameEnded {
					break
				}
				return errors.Wrap(err, "waiting for turn failed")
			}

			bot.playok.RLock()
			info("\n%s", bot.playok.currentTable.board.String())
			bot.playok.RUnlock()

			info("bot is to move")

			discCount, err := bot.computeAndSendMove()
			if err != nil {
				return errors.Wrap(err, "compute and send move failed")
			}

			info("bot sent move")

			if err := bot.awaitMoveConfirmation(discCount); err != nil {
				return errors.Wrap(err, "waiting for move confirmation failed")
			}

			info("bot received move confirmation")

			bot.playok.RLock()
			info("\n%s", bot.playok.currentTable.board.String())
			bot.playok.RUnlock()

			time.Sleep(time.Second)
		}
	}
}

// awaitMoveConfirmation waits until the server tells us the move is received
// we get an update of our own move for some reason
func (bot *Bot) awaitMoveConfirmation(discCount int) error {
	retries := 30
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
		time.Sleep(200 * time.Millisecond)
	}
}

func (bot *Bot) computeAndSendMove() (int, error) {

	info("computing move")

	bot.playok.RLock()
	board := bot.playok.currentTable.board
	bot.playok.RUnlock()

	othelloBot := treesearch.NewBot(log.Writer(), 10, 16)

	move, err := othelloBot.DoMove(board.Board)
	if err != nil {
		return 0, errors.Wrap(err, "bot failed to compute move")
	}

	moveBit := (board.Me() | board.Opp()) ^ (move.Me() | move.Opp())
	moveID := moveBit.Lowest()

	randomDelay := time.Duration(500+rand.Intn(500)+rand.Intn(500)) * time.Millisecond
	info("delaying sending move %dms", randomDelay.Milliseconds())
	time.Sleep(randomDelay)

	info("sending move")
	if err := bot.sendMove(moveID); err != nil {
		return 0, errors.Wrap(err, "failed to send move")
	}

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
			return errors.New("max retries reached")
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

func (bot *Bot) awaitTablesList() error {

	retries := 5
	for {
		info("await tables list: %d retries left", retries)

		bot.playok.RLock()
		success := len(bot.playok.tables) != 0
		bot.playok.RUnlock()

		if success {
			return nil
		}

		retries--
		if retries == 0 {
			return errors.New("max retries reached")
		}
		time.Sleep(time.Second)
	}
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
			if table.countPlayers() == 1 && table.minRatingID == 0 {
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
			return errors.New("max retries reached")
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
			return errors.New("max retries reached")
		}
		time.Sleep(time.Second)
	}
}
