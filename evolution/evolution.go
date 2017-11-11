package evolution

import (
	"dots/board"
	"dots/players"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/bits"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"
)

const (
	maxParticipants = 15
	WHITE           = 0
	DRAW            = 1
	BLACK           = 2
)

type Participant struct {
	Params players.Parameters
	rating int
}

func NewParticipantRandom(min, max int) (particpant Participant) {
	return Participant{
		Params: players.RandomParameters(min, max),
		rating: 1500}
}

func NewParticipantMutate(parent Participant) (offspring Participant) {
	offspring = parent

	for i := range offspring.Params.PositionValue {
		if rand.Intn(8) == 0 {
			offspring.Params.PositionValue[i] += rand.Intn(100) - rand.Intn(100)
		}
	}
	offspring.rating = 1500
	return
}

type result struct {
	blackId int
	whiteId int
	winner  int
}

// Evolution runs an evolutionary algorithm
type Evolution struct {
	killChan     chan bool
	resultChan   chan result
	participants []Participant
	filename     string
	searchDepth  int
	exactDepth   int
	runningGames int
	launchesLeft int
}

// NewEvolution creates a new Evolution
func NewEvolution(filename string) (evolution *Evolution) {
	evolution = &Evolution{
		killChan:    make(chan bool),
		resultChan:  make(chan result),
		filename:    filename,
		searchDepth: 5,
		exactDepth:  10}

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		evolution.killChan <- true
	}()

	evolution.Load()
	for len(evolution.participants) < maxParticipants {
		evolution.participants = append(evolution.participants, NewParticipantRandom(-10000, 10000))
	}
	return
}

func (evolution *Evolution) playGame(blackId, whiteId int, b board.Board) {
	players := []players.Player{
		players.NewBotEvolve(evolution.searchDepth, evolution.exactDepth, evolution.participants[blackId].Params),
		players.NewBotEvolve(evolution.searchDepth, evolution.exactDepth, evolution.participants[whiteId].Params)}

	turn := 0

	for {
		if b.Moves() == 0 {
			b.SwitchTurn()
			turn = 1 - turn

			if b.Moves() == 0 {
				turn = 1 - turn
				break
			}
		}

		b = players[turn].DoMove(b)
	}

	r := result{
		blackId: blackId,
		whiteId: whiteId,
	}

	if turn == 1 {
		b.SwitchTurn()
	}

	blackCount := bits.OnesCount64(b.Me())
	whiteCount := bits.OnesCount64(b.Opp())

	if blackCount < whiteCount {
		r.winner = WHITE
	} else if blackCount > whiteCount {
		r.winner = BLACK
	} else {
		r.winner = DRAW
	}

	evolution.resultChan <- r
}

func (evolution *Evolution) launchGames() {
	for i := range evolution.participants {
		j := rand.Intn(len(evolution.participants))
		if i != j {
			evolution.participants[i], evolution.participants[j] = evolution.participants[j], evolution.participants[i]
		}
	}

	for g := 0; g < len(evolution.participants)/2; g++ {
		xot := board.Xot()
		go evolution.playGame(2*g, (2*g)+1, xot)
	}
	evolution.runningGames = len(evolution.participants) / 2
	evolution.launchesLeft--
}

// Run runs evolutionary algorithm
func (evolution *Evolution) Run() {
	for {
		timeoutChan := make(chan bool)

		go func() {
			for {
				time.Sleep(1 * time.Second)
				timeoutChan <- true
			}
		}()

		evolution.launchesLeft = 10
		evolution.launchGames()

		for {
			select {
			case <-evolution.killChan:
				fmt.Printf("\nCleaning up\n")
				evolution.Save()
				return
			case r := <-evolution.resultChan:
				score := float64(r.winner) / 2
				ratingDiff := evolution.participants[r.whiteId].rating - evolution.participants[r.blackId].rating
				expected := 1.0 / (1.0 + (math.Pow(10.0, (float64(ratingDiff) / 400.0))))
				change := int(32.0 * (score - expected))
				evolution.participants[r.blackId].rating += change
				evolution.participants[r.whiteId].rating -= change
				evolution.runningGames--
				fmt.Printf("%d games left running.\n", evolution.runningGames)
				if evolution.runningGames == 0 {

					fmt.Printf("%d launches left.\n", evolution.launchesLeft)

					if evolution.launchesLeft == 0 {
						fmt.Printf("Generating new participants.\n")
						sort.Slice(evolution.participants, func(i, j int) bool {
							return evolution.participants[i].rating > evolution.participants[j].rating
						})

						removedCount := len(evolution.participants) / 4

						for i := len(evolution.participants) - removedCount; i < len(evolution.participants); i++ {
							evolution.participants[i] = NewParticipantMutate(evolution.participants[rand.Intn(len(evolution.participants)-removedCount)])
						}

						evolution.launchesLeft = 10
					}

					fmt.Printf("Launching new games.\n")
					evolution.launchGames()
				}
				fmt.Printf("\n")
				break
			case <-timeoutChan:
				participantsCopy := make([]Participant, len(evolution.participants))
				copy(participantsCopy, evolution.participants)
				sort.Slice(participantsCopy, func(i, j int) bool {
					return participantsCopy[i].rating > participantsCopy[j].rating
				})
				fmt.Printf("Rating\tParams\n")
				for _, p := range participantsCopy {
					fmt.Printf("%d\t%v\n", p.rating, p.Params)
				}
				fmt.Printf("\n\n")
			}
		}
	}
}

// Save saves the state of Evolution
func (evolution *Evolution) Save() {
	file, err := os.OpenFile(evolution.filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(file).Encode(evolution.participants)
}

// Load loads the state of Evolution
func (evolution *Evolution) Load() {
	file, err := os.Open(evolution.filename)
	if err != nil {
		fmt.Printf("Could not load state: %s", err)
		return
	}
	json.NewDecoder(file).Decode(&evolution.participants)
	for i := range evolution.participants {
		evolution.participants[i].rating = 1500
	}
	fmt.Printf("Loaded state\n\n")
}
