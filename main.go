package main

import (
	"dots/cli_game"
	"dots/players"
)

func main() {
	random_bot := players.NewBotRandom()
	human := &players.Human{}

	cli_game := cli_game.NewCliGame(random_bot, human)
	cli_game.Run()
}
