package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func worker(i int, data <-chan int, exit chan struct{}) {
	for {
		select {
		case d := <-data:
			fmt.Printf("Worker %d prints: %d\n", i, d)
		case <-exit:
			fmt.Printf("Worker %d is shutting down...\n", i)
			return
		}
	}
}

func main() {
	data := make(chan int)
	exit := make(chan struct{})
	signalCall := make(chan os.Signal)
	var workers int
	fmt.Printf("Input num of workers: ")
	fmt.Scan(&workers)
	for i := 0; i < workers; i++ {
		go worker(i+1, data, exit)
	}

	signal.Notify(signalCall, os.Interrupt)

	for {
		select {
		case data <- rand.Int():
			continue
		case <-signalCall:
			for i := 0; i < workers; i++ {
				exit <- struct{}{}
			}
			time.Sleep(time.Second)
			return
		}
	}
}
