package evolution

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// State is a dummy
type State struct {
	Timestamp int64
}

// Evolution runs an evolutionary algorithm
type Evolution struct {
	killChan   chan bool
	resultChan chan State
	state      State
	filename   string
}

// NewEvolution creates a new Evolution
func NewEvolution(filename string) (evolution *Evolution) {
	evolution = &Evolution{
		killChan:   make(chan bool),
		resultChan: make(chan State),
		filename:   filename}

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		evolution.killChan <- true
	}()

	_ = evolution.Load()

	return
}

// Run runs evolutionary algorithm
func (evolution *Evolution) Run() {

	work := func() {
		time.Sleep(1 * time.Second)
		evolution.resultChan <- State{Timestamp: time.Now().Unix()}
	}

	go work()

	for {
		select {
		case <-evolution.killChan:
			fmt.Printf("\nCleaning up\n")
			evolution.Save()
			return
		case evolution.state = <-evolution.resultChan:
			fmt.Printf("Got result %d\n", evolution.state.Timestamp)
			go work()
		}
	}
}

// Save saves the state of Evolution
func (evolution *Evolution) Save() {
	file, err := os.OpenFile(evolution.filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(file).Encode(evolution.state)
}

// Load loads the state of Evolution
func (evolution *Evolution) Load() (err error) {
	file, err := os.Open(evolution.filename)
	if err != nil {
		fmt.Printf("Could not load state: %s", err)
		return
	}
	json.NewDecoder(file).Decode(&evolution.state)
	fmt.Printf("Loaded state: %d\n", evolution.state.Timestamp)
	return
}
