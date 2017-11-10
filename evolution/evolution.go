package evolution

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run runs evolutionary algorithm
func Run() {
	signals := make(chan os.Signal)
	killChan := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		killChan <- true
	}()

	for {
		select {
		case <-killChan:
			fmt.Printf("exiting\n")
			return
		default:
			fmt.Printf("%v\n", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
