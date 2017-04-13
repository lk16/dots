package main

import (
	"math/rand"
	"os"
	"time"

	"dots/cli_game"
	"dots/heuristic"
	"dots/minimax"
	"dots/players"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	human := players.NewHuman(os.Stdin)
	smart_bot := players.NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{}, 7, 12)

	cli_game := cli_game.NewCliGame(human, smart_bot, os.Stdout)
	cli_game.Run()
}
