package main

import (
	"flag"
	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/treesearch"
	"github.com/lk16/dots/internal/web"
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
		log.Fatalf("Error creating profiling file: %s", err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatalf("Profiling error: %s", err)
	}
	defer pprof.StopCPUProfile()

	profiled()
}

func devMain() {
	log.Printf("devMain() running")
}

func main() {

	defaultSeed := time.Now().UTC().UnixNano()
	seed := flag.Int64("seed", defaultSeed, "Custom seed")

	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")

	dev := flag.Bool("dev", false, "run devMain()")

	flag.Parse()

	rand.Seed(*seed)

	if *dev {
		devMain()
		return
	}

	if *cpuprofile != "" {
		profile(*cpuprofile, func() {
			rand.Seed(0)
			exactDepth := 18
			bot := treesearch.NewBot(ioutil.Discard, 2, exactDepth)
			for i := 0; i < 1; i++ {
				board := othello.NewXotBoard()
				for board.CountEmpties() >= exactDepth {
					var err error
					board, err = bot.DoMove(*board)
					if err != nil {
						log.Printf("error: %s", err)
						return
					}
				}
				log.Printf("bot lifetime stats: %s\n", bot.LifetimeStats.String())
			}
		})
		return
	}

	web.Main()

}
