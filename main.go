package main

import (
	"math/rand"
	"time"

	"dots/cli_game"
	"dots/heuristic"
	"dots/minimax"
	"dots/players"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	human := &players.Human{}
	smart_bot := players.NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{}, 5, 10)

	cli_game := cli_game.NewCliGame(human, smart_bot)
	cli_game.Run()
}
