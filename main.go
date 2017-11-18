package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"dots/evolution"
	"dots/frontend"
	"dots/players"

	"runtime/pprof"
)

func main() {

	defaultSeed := time.Now().UTC().UnixNano()
	seed := flag.Int64("seed", defaultSeed, "Custom seed")

	blackName := flag.String("bp", "human", "Black player: Bot name or \"human\"")
	blackLevel := flag.Int("bl", 5, "Black player search level (ignored for human)")

	whiteName := flag.String("wp", "human", "White player: Bot name or \"human\"")
	whiteLevel := flag.Int("wl", 5, "White player search level (ignored for human)")

	frontendName := flag.String("frontend", "gtk", "Frontend: \"gtk\" or \"cli\"")

	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")

	parallelSearch := flag.Bool("ps", true, "enable parallel search")

	evolutionFileName := flag.String("evolution", "", "run evolutionary algorithm with savefile name as parameter")

	enableXot := flag.Bool("xot", false, "enable xot")

	flag.Parse()

	rand.Seed(*seed)

	if *evolutionFileName != "" {
		evolution.NewEvolution(*evolutionFileName).Run()
		return
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fe := frontend.Get(*frontendName)

	whitePlayer := players.Get(*whiteName, *whiteLevel, *parallelSearch)
	blackPlayer := players.Get(*blackName, *blackLevel, *parallelSearch)

	controller := frontend.NewController(blackPlayer, whitePlayer, os.Stdout, fe, *enableXot)
	controller.Run()
}
