package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"dots/frontend"
	"dots/heuristic"
	"dots/minimax"
	"dots/players"
)

func main() {

	default_seed := time.Now().UTC().UnixNano()
	seed := flag.Int64("seed", default_seed, "Custom seed")

	black_name := flag.String("bp", "human", "Black player: Bot name or \"human\"")
	black_lvl := flag.Uint("bl", 5, "Black player search level (ignored for human)")

	white_name := flag.String("wp", "human", "White player: Bot name or \"human\"")
	white_lvl := flag.Uint("wl", 5, "White player search level (ignored for human)")

	frontend_name := flag.String("frontend", "gtk", "Frontend: \"gtk\" or \"cli\"")

	flag.Parse()

	rand.Seed(*seed)

	getPlayer := func(name string, lvl uint) (player players.Player) {

		if name == "human" {
			player = nil
		} else if name == "random" {
			player = players.NewBotRandom()
		} else if name == "heur" {
			search_depth := lvl
			perfect_depth := 2 * lvl
			if perfect_depth > 6 {
				perfect_depth -= 2
			}
			player = players.NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{},
				search_depth, perfect_depth, os.Stdout)
		} else {
			fmt.Printf("Invalid player name %s\n", name)
			os.Exit(1)
		}
		return

	}

	var fe frontend.Frontend
	if *frontend_name == "gtk" {
		fe = frontend.NewGtk()
	} else if *frontend_name == "cli" {
		fe = frontend.NewCommandLine()
	} else {
		fmt.Printf("Invalid frontend name %s\n", *frontend_name)
		os.Exit(1)
	}

	white_player := getPlayer(*white_name, *white_lvl)
	black_player := getPlayer(*black_name, *black_lvl)

	controller := frontend.NewController(black_player, white_player, os.Stdout, fe)
	controller.Run()
}
