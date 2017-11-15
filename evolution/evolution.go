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
	maxParticipants = 40
	WHITE           = 0
	DRAW            = 1
	BLACK           = 2
	startRating     = 1500
	paramMax        = 10000
)

type Participant struct {
	Params players.Parameters
	rating int
	Wins   int
	Losses int
	Draws  int
}

func NewParticipantRandom(min, max int) (particpant Participant) {
	return Participant{
		Params: players.RandomParameters(min, max),
		rating: startRating}
}

func NewParticipantMutate(parent Participant) (offspring Participant) {
	offspring = parent

	for i := range offspring.Params.PositionValue {
		if rand.Intn(8) == 0 {
			offspring.Params.PositionValue[i] += rand.Intn(500) - rand.Intn(500)
		}
	}
	offspring.rating = startRating
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
		searchDepth: 4,
		exactDepth:  8}

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		evolution.killChan <- true
	}()

	evolution.Load()
	for len(evolution.participants) < maxParticipants {
		evolution.participants = append(evolution.participants, NewParticipantRandom(-paramMax, paramMax))
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
				if r.winner == WHITE {
					evolution.participants[r.whiteId].Wins++
					evolution.participants[r.blackId].Losses++
				} else if r.winner == DRAW {
					evolution.participants[r.whiteId].Draws++
					evolution.participants[r.blackId].Draws++
				} else {
					evolution.participants[r.whiteId].Losses++
					evolution.participants[r.blackId].Wins++
				}

				evolution.runningGames--
				if evolution.runningGames == 0 {

					if evolution.launchesLeft == 0 {
						sort.Slice(evolution.participants, func(i, j int) bool {
							return evolution.participants[i].rating > evolution.participants[j].rating
						})

						removedParticipantsCount := len(evolution.participants) / 2
						removedParticipantsStart := len(evolution.participants) - removedParticipantsCount
						newRandomParticipantsStart := len(evolution.participants) - (removedParticipantsCount / 3)

						for i := removedParticipantsStart; i < len(evolution.participants); i++ {
							if i >= newRandomParticipantsStart {
								evolution.participants[i] = NewParticipantRandom(-paramMax, paramMax)
							} else {
								evolution.participants[i] = NewParticipantMutate(evolution.participants[rand.Intn(removedParticipantsStart)])
							}
							evolution.participants[i].Draws = 0
							evolution.participants[i].Wins = 0
							evolution.participants[i].Losses = 0
						}

						for i := range evolution.participants {
							evolution.participants[i].rating = startRating
						}

						evolution.launchesLeft = 10
					}
					evolution.launchGames()
				}
				break
			case <-timeoutChan:
				participantsCopy := make([]Participant, len(evolution.participants))
				copy(participantsCopy, evolution.participants)
				sort.Slice(participantsCopy, func(i, j int) bool {
					return participantsCopy[i].rating > participantsCopy[j].rating
				})

				fmt.Printf("%d launches left.\n", evolution.launchesLeft)
				fmt.Printf("%d games running.\n", evolution.runningGames)
				fmt.Printf("Rating\tGames\tWins\tLosses\tDraws\tWin Rate\tParams\n")
				for _, p := range participantsCopy {
					games := p.Draws + p.Wins + p.Losses
					winRate := 100.0 * float64(p.Wins) / float64(games)
					fmt.Printf("%d\t%d\t%d\t%d\t%d\t%2.2f%%\t%v\n", p.rating, games, p.Wins, p.Losses, p.Draws, winRate, p.Params)
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
		evolution.participants[i].rating = startRating
	}
	fmt.Printf("Loaded state\n\n")
}
