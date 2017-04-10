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

	random_bot := players.NewBotRandom()

	smart_bot := players.NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{}, 5, 10)

	cli_game := cli_game.NewCliGame(random_bot, smart_bot)
	cli_game.Run()
}
