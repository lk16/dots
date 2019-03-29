package main

import (
	"flag"
	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/players"
	"github.com/lk16/dots/internal/web"
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
			for i := 0; i < 1; i++ {
				board := othello.NewXotBoard()
				board.ASCIIArt(os.Stdout, false)
				bot := players.NewBotHeuristic(os.Stdout, 12, 0)
				bot.DoMove(board)
			}
		})
		return
	}

	web.Main()

}
