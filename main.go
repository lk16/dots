package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"dots/frontend"
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

	fe := frontend.Get(*frontend_name)

	white_player := players.Get(*white_name, *white_lvl)
	black_player := players.Get(*black_name, *black_lvl)

	controller := frontend.NewController(black_player, white_player, os.Stdout, fe)
	controller.Run()
}
