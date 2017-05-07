package main

import (
	"math/rand"
	"os"
	"time"

	"dots/frontend"
	"dots/heuristic"
	"dots/minimax"
	"dots/players"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	//human := players.NewHuman(os.Stdin)
	smart_bot := players.NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{}, 7, 12, os.Stdout)

	controller := frontend.NewController(nil, smart_bot, os.Stdout,
		frontend.NewCommandLine())
	controller.Run()
}
