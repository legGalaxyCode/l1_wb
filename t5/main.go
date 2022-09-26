package main

import (
	"fmt"
	"math/rand"
	"time"
)

func write(data chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Shutting down from write...")
			return
		default:
			time.Sleep(300 * time.Millisecond)
			data <- rand.Int()
		}
	}
}

func read(data chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Shutting down from main")
			return
		case <-data:
			fmt.Printf("Value: %d\n", <-data)
		}
	}
}

func main() {
	data := make(chan int)
	done := make(chan struct{})

	var workTime string
	fmt.Printf("Input time in seconds: ")
	fmt.Scan(&workTime)
	timeout, _ := time.ParseDuration(workTime)

	go func(done chan<- struct{}, dur time.Duration) {
		for {
			select {
			case <-time.After(dur):
				fmt.Println("Time has expired")
				for {
					done <- struct{}{}
				}
			}
		}
	}(done, timeout)

	go write(data, done)
	read(data, done)
}
