package main

import (
	"flag"
	"github.com/lk16/dots/othello"
	"github.com/lk16/dots/players"
	"github.com/lk16/dots/web"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func profile(cpuprofile string, profiled func()) {

	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Printf("Error creating profiling file: %s", err)
		return
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Printf("Profiling error: %s", err)
		return
	}
	defer pprof.StopCPUProfile()

	profiled()
}

func main() {

	defaultSeed := time.Now().UTC().UnixNano()
	seed := flag.Int64("seed", defaultSeed, "Custom seed")

	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()

	rand.Seed(*seed)

	if *cpuprofile != "" {
		profile(*cpuprofile, func() {
			rand.Seed(0)
			board := othello.NewXotBoard()
			bot := players.NewBotHeuristic(ioutil.Discard, 12, 0)
			bot.DoMove(board)
		})
		return
	}

	web.Main()

}
