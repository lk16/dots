package playok

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func (bot *Bot) sendMessage(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	log.Printf("SEND %s", string(bytes))
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
			rand.Intn(100), // TODO random timestamp value?
		},
	}

	return bot.sendMessage(message)
}

func (bot *Bot) sendTakeSeatRequest(tableID, seatID int) error {
	message := Message{
		I: []int{
			83,
			tableID,
			seatID,
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

func (bot *Bot) sendInitMessage() error {

	var kSessionCookie string
	for _, cookie := range bot.connector.browser.Jar.Cookies(playokParsedURL) {
		if cookie.Name == "ksession" {
			kSessionCookie = cookie.Value
		}
	}

	if kSessionCookie == "" {
		log.Printf("warning: could not find ksession cookie")
	}

	splitKSession := strings.Split(kSessionCookie, ":")

	firstArg := fmt.Sprintf("%s+|%s|%s",
		splitKSession[0],       // part of a cookie
		bot.connector.windowAp, // scraped JS value window.ap
		bot.connector.windowGe, // scraped JS value window.ge
	)

	nowMilli := time.Now().UnixNano() / 1000000

	message := Message{
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

	return bot.sendMessage(message)
}
