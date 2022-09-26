package main

import (
	"fmt"
	"time"
)

// прерываение горутины через канал
func main() {
	done := make(chan struct{})
	go func(done <-chan struct{}) {
		for {
			select {
			case <-done:
				fmt.Println("Shutting down from g")
				return
			default:
				fmt.Println("Hello from g")
				time.Sleep(time.Second)
			}
		}
	}(done)
	time.Sleep(time.Second * 3)
	done <- struct{}{}
	time.Sleep(time.Second)
}
