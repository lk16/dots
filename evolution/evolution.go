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

type output struct {
	Timestamp int64
}

func writeState(latest output) {
	file, err := os.OpenFile("evolution.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(file).Encode(latest)
}

func readState() (state output, err error) {
	file, err := os.Open("evolution.json")
	if err != nil {
		// TODO
	}
	json.NewDecoder(file).Decode(&state)
	return
}

// Run runs evolutionary algorithm
func Run() {
	signals := make(chan os.Signal)
	killChan := make(chan bool)
	resultChan := make(chan output)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		killChan <- true
	}()

	work := func() {
		time.Sleep(1 * time.Second)
		resultChan <- output{Timestamp: time.Now().Unix()}
	}

	state, err := readState()
	if err != nil {
		fmt.Printf("Could not load state: %s", err)
	}
	fmt.Printf("Loaded state: %d\n", state.Timestamp)

	go work()

	for {
		select {
		case <-killChan:
			fmt.Printf("\nCleaning up\n")
			writeState(state)
			return
		case state = <-resultChan:
			fmt.Printf("Got result %d\n", state.Timestamp)
			go work()
		}
	}
}
