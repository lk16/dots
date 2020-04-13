package playok

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"sort"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
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
	connector *connector
	websocket *websocket.Conn
	playok    *state
	errChan   chan error
}

// NewBot initializes a new bot
func NewBot(userName, password string) *Bot {

	return &Bot{
		connector: newConnector(userName, password),
		playok:    newState(),
		errChan:   make(chan error),
	}
}

// Run is the entrypoint of the Bot
func (bot *Bot) Run() error {

	websocket, err := bot.connector.connect()
	if err != nil {
		return errors.Wrap(err, "failed to connect to websocket")
	}
	bot.websocket = websocket

	err = bot.sendInitMessage()
	if err != nil {
		return errors.Wrap(err, "sending init message failed")
	}

	go bot.loopKeepAlive()
	go bot.loopHandleMessage()
	go bot.loopPrintState()
	go bot.takeActionEntryPoint()

	// stop bot if any error shows up
	return <-bot.errChan
}

func (bot *Bot) loopKeepAlive() {
	for {
		time.Sleep(30 * time.Second)
		if err := bot.sendKeepAlive(); err != nil {
			bot.errChan <- errors.Wrap(err, "sending keep alive failed")
			return
		}
	}
}

func (bot *Bot) loopHandleMessage() {
	for {
		messageType, messageBytes, err := bot.websocket.ReadMessage()
		if err != nil {
			bot.errChan <- errors.Wrap(err, "read message error:")
			return
		}

		if messageType != websocket.TextMessage {
			log.Printf("RECV ignoring message with type %d: %s", messageType, string(messageBytes))
			continue
		}

		if err = bot.handleMessage(messageBytes); err != nil {
			bot.errChan <- errors.Wrap(err, "message handling error")
			return
		}
	}
}

func (bot *Bot) loopPrintState() {

	for {
		time.Sleep(500 * time.Millisecond)
		//bot.printState()
	}
}

func (bot *Bot) takeActionEntryPoint() {
	if err := bot.takeAction(); err != nil {
		bot.errChan <- errors.Wrap(err, "bot action failed")
		return
	}
	log.Printf("INFO bot.takeAction() returned")
}

// printState prints the bot state in ascii art
func (bot *Bot) printState() {

	bot.playok.RLock()
	defer bot.playok.RUnlock()

	var tableIDs []int
	for ID := range bot.playok.tables {
		tableIDs = append(tableIDs, ID)
	}

	sort.Ints(tableIDs)

	var buff bytes.Buffer

	/*
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
	*/

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
