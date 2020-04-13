package playok

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

// takeAction gives the bot to take initiative, such as joining tables or doing moves
func (bot *Bot) takeAction() error {

	log.Printf("DEBG: takeAction is not implemented")
	return nil

	// wait until we have the tables list
	retries := 5
	for {
		bot.playok.RLock()
		success := len(bot.playok.tables) != 0
		bot.playok.RUnlock()

		if !success {
			retries--
			if retries == 0 {
				return errors.New("get tables list: max retries reached")
			}
			time.Sleep(time.Second)
		}
		break
	}

	// join some table with one player
	for tableID, table := range bot.playok.tables {
		// only care about public tables with one player
		if table.minRatingID != 0 || (table.players[0] == "" && table.players[1] == "") || (table.players[0] != "" && table.players[1] != "") {
			continue
		}

		if err := bot.sendJoinTableRequest(tableID); err != nil {
			return errors.Wrap(err, "joining table request failed")
		}
		break
	}

	// wait until we joined the table
	for {
		bot.playok.RLock()
		success := bot.playok.currentTable.ID != 0
		bot.playok.RUnlock()

		if !success {
			retries--
			if retries == 0 {
				return errors.New("joining table: max retries reached")
			}
			time.Sleep(time.Second)
		}
		break
	}

	time.Sleep(time.Second)

	bot.playok.RLock()

	seatID := 1
	if bot.playok.currentTable.players[0] != "" {
		seatID = 0
	}

	bot.playok.RUnlock()

	if err := bot.sendTakeSeatRequest(bot.playok.currentTable.ID, seatID); err != nil {
		return errors.Wrap(err, "taking seat request failed")
	}

	for {
		bot.playok.RLock()
		success := bot.playok.currentTable.players[0] == bot.playok.userName || bot.playok.currentTable.players[1] == bot.playok.userName
		bot.playok.RUnlock()

		if !success {
			retries--
			if retries == 0 {
				return errors.New("taking seat: max retries reached")
			}
			time.Sleep(time.Second)
		}
		break
	}

	// start game

	// play full game

	return nil
}
