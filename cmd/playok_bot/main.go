package main

import (
	"log"
	"os"

	"github.com/lk16/dots/internal/playok"
	"github.com/pkg/errors"
)

var (
	userName = os.Getenv("PLAYOK_USERNAME")
	password = os.Getenv("PLAYOK_PASSWORD")
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	bot := playok.NewBot(userName, password)
	err := bot.Run()

	if err != nil {
		log.Print("bot crashed:")
		log.Print(err.Error())

		if stackTraceErr, ok := err.(stackTracer); ok {
			log.Printf("%+v", stackTraceErr.StackTrace())
		}

		os.Exit(1)
	}
	log.Print("bot exited normally")
}
