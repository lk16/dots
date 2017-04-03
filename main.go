package main

import (
    "dots/players"
    "dots/cli_game"
)

func main() {
    random_bot := players.NewBotRandom()

    cli_game := cli_game.NewCliGame(random_bot,random_bot)
    cli_game.RunGame()
}

