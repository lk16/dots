package main

import (
	"dots/cli_game"
	"dots/minimax"
	"dots/players"
)

func main() {
	random_bot := players.NewBotRandom()

	smart_bot := players.NewBotHeuristic(players.SquaredHeuristic, &minimax.Minimax{}, 6, 10)

	cli_game := cli_game.NewCliGame(random_bot, smart_bot)
	cli_game.Run()
}
