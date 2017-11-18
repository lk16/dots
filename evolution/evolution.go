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
	"runtime"
	"sort"
	"syscall"
	"time"
)

const (
	maxParticipants = 20
	startRating     = 1500
	paramMax        = 10000
)

type Participant struct {
	Params players.Parameters
	Rating int
	Wins   int
	Losses int
	Draws  int
}

func NewParticipantRandom(min, max int) (particpant Participant) {
	return Participant{
		Params: players.RandomParameters(min, max),
		Rating: startRating}
}

func NewParticipantMutate(parent Participant) (offspring Participant) {
	offspring = parent

	for i := range offspring.Params.PositionValue {
		if rand.Intn(2) == 0 {
			offspring.Params.PositionValue[i] += rand.Intn(500) - rand.Intn(500)
		}
	}
	offspring.Rating = startRating
	return
}

type result struct {
	blackID    int
	whiteID    int
	blackScore float64
}

// Evolution runs an evolutionary algorithm
type Evolution struct {
	killChan          chan bool
	resultChan        chan result
	participants      map[int]Participant
	filename          string
	searchDepth       int
	exactDepth        int
	nextParticipantID int
}

// NewEvolution creates a new Evolution
func NewEvolution(filename string) (evolution *Evolution) {
	evolution = &Evolution{
		killChan:     make(chan bool),
		resultChan:   make(chan result),
		filename:     filename,
		searchDepth:  7,
		exactDepth:   10,
		participants: make(map[int]Participant, 0)}

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		evolution.killChan <- true
	}()

	evolution.Load()

	return
}

func playGame(
	blackID int, blackParams players.Parameters,
	whiteID int, whiteParams players.Parameters,
	searchDepth, exactDepth int, b board.Board, ch chan result) {
	players := []players.Player{
		players.NewBotEvolve(searchDepth, exactDepth, blackParams),
		players.NewBotEvolve(searchDepth, exactDepth, whiteParams)}

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

	if turn == 1 {
		b.SwitchTurn()
	}

	blackCount := bits.OnesCount64(b.Me())
	whiteCount := bits.OnesCount64(b.Opp())

	r := result{
		blackID: blackID,
		whiteID: whiteID}

	if blackCount < whiteCount {
		r.blackScore = 0
	} else if blackCount > whiteCount {
		r.blackScore = 1
	} else {
		r.blackScore = 0.5
	}

	ch <- r
}

func (evolution *Evolution) launchGame() {

	a := rand.Intn(len(evolution.participants))
	var b int

	for a == b {
		b = rand.Intn(len(evolution.participants))
	}

	var blackID, whiteID int

	for ID := range evolution.participants {
		if a == 0 {
			blackID = ID
		}
		if b == 0 {
			whiteID = ID
		}
		a--
		b--
	}

	xot := board.Xot()
	blackParams := evolution.participants[blackID].Params
	whiteParams := evolution.participants[whiteID].Params

	go playGame(blackID, blackParams, whiteID, whiteParams,
		evolution.searchDepth, evolution.exactDepth, xot,
		evolution.resultChan)
}

func (evolution *Evolution) printStats() {
	participantsCopy := []Participant{}
	for _, v := range evolution.participants {
		participantsCopy = append(participantsCopy, v)
	}
	sort.Slice(participantsCopy, func(i, j int) bool {
		return participantsCopy[i].Rating > participantsCopy[j].Rating
	})

	fmt.Printf("Rating\tGames\tWins\tLosses\tDraws\tWin Rate\tParams\n")
	total := Participant{}
	for _, p := range participantsCopy {
		games := p.Draws + p.Wins + p.Losses
		winRate := 100.0 * float64(p.Wins) / float64(games)
		fmt.Printf("%d\t%d\t%d\t%d\t%d\t%2.2f%%\t%v\n", p.Rating, games, p.Wins, p.Losses, p.Draws, winRate, p.Params)
		total.Draws += p.Draws
		total.Losses += p.Losses
		total.Rating += p.Rating
		total.Wins += p.Wins
	}
	totalGames := total.Draws + total.Wins + total.Losses
	totalWinRate := 100 * float64(total.Wins) / float64(totalGames)
	fmt.Printf("---\n")
	fmt.Printf("%d\t%d\t%d\t%d\t%d\t%2.2f%%\t%v\n", total.Rating, totalGames, total.Wins, total.Losses, total.Draws, totalWinRate, total.Params)
	fmt.Printf("\n\n")
}

// Run runs evolutionary algorithm
func (evolution *Evolution) Run() {
	timeoutChan := make(chan bool)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			timeoutChan <- true
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		evolution.launchGame()
	}

	for {
		select {
		case <-evolution.killChan:
			fmt.Printf("\nCleaning up\n")
			evolution.Save()
			return
		case r := <-evolution.resultChan:
			evolution.updateRating(r)
			evolution.launchGame()
			break
		case <-timeoutChan:
			evolution.replaceParticipants()
			evolution.printStats()
		}
	}
}

func (evolution *Evolution) replaceParticipants() {

	deadParticipants := []int{}
	inflationTotal := 0

	for ID, p := range evolution.participants {
		if p.Rating > startRating {
			continue
		}

		games := p.Draws + p.Losses + p.Draws
		if p.Rating < (startRating-300)+(20*games) {
			deadParticipants = append(deadParticipants, ID)
			inflationTotal += (startRating - p.Rating)
		}
	}

	for _, ID := range deadParticipants {
		delete(evolution.participants, ID)
	}

	for _ = range deadParticipants {
		var newParticipant Participant
		if rand.Intn(3) == 0 {
			newParticipant = NewParticipantRandom(-paramMax, paramMax)
		} else {
			n := rand.Intn(len(evolution.participants))
			var parent Participant
			for _, parent = range evolution.participants {
				if n == 0 {
					break
				}
			}
			newParticipant = NewParticipantMutate(parent)
		}
		newParticipant.Rating = startRating
		evolution.participants[evolution.nextParticipantID] = newParticipant
		evolution.nextParticipantID++
	}

	i := 0
	for ID, p := range evolution.participants {
		p.Rating -= inflationTotal / len(evolution.participants)
		if i < inflationTotal%len(evolution.participants) {
			p.Rating--
		}
		evolution.participants[ID] = p

		i++
	}
}

func (evolution *Evolution) updateRating(r result) {
	RatingDiff := evolution.participants[r.whiteID].Rating - evolution.participants[r.blackID].Rating
	expected := 1.0 / (1.0 + (math.Pow(10.0, (float64(RatingDiff) / 400.0))))
	blackChange := int(32.0 * (r.blackScore - expected))

	var whiteParticipant, blackParticipant Participant
	var ok bool

	if whiteParticipant, ok = evolution.participants[r.whiteID]; !ok {
		return
	}
	if blackParticipant, ok = evolution.participants[r.blackID]; !ok {
		return
	}

	blackParticipant.Rating += blackChange
	whiteParticipant.Rating -= blackChange
	if r.blackScore == 1.0 {
		blackParticipant.Wins++
		whiteParticipant.Losses++
	} else if r.blackScore == 0.5 {
		blackParticipant.Draws++
		whiteParticipant.Draws++
	} else {
		blackParticipant.Losses++
		whiteParticipant.Wins++
	}

	evolution.participants[r.blackID] = blackParticipant
	evolution.participants[r.whiteID] = whiteParticipant
}

// Save saves the state of Evolution
func (evolution *Evolution) Save() {
	file, err := os.OpenFile(evolution.filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	var ps []Participant
	for _, p := range evolution.participants {
		ps = append(ps, p)
	}
	json.NewEncoder(file).Encode(ps)
}

// Load loads the state of Evolution
func (evolution *Evolution) Load() {
	file, err := os.Open(evolution.filename)
	if err == nil {
		var ps []Participant
		json.NewDecoder(file).Decode(&ps)

		evolution.participants = make(map[int]Participant, len(ps))
		for i, p := range ps {
			p.Rating = startRating
			evolution.participants[i] = p
		}
	} else {
		fmt.Printf("Could not load state: %s", err)
	}

	evolution.nextParticipantID = len(evolution.participants)
	if len(evolution.participants) < maxParticipants {
		fmt.Printf("Adding random initial participants\n")
		for len(evolution.participants) < maxParticipants {
			p := NewParticipantRandom(-paramMax, paramMax)
			p.Rating = startRating
			evolution.participants[evolution.nextParticipantID] = p
			evolution.nextParticipantID++
		}
	}

}
