package main

import (
	"dots/cli_game"
	"dots/players"
)

func main() {
	random_bot := players.NewBotRandom()

	cli_game := cli_game.NewCliGame(random_bot, random_bot)
	cli_game.RunGame()
}
