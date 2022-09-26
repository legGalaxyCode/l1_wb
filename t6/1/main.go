package main

import (
	"fmt"
	"time"
)

// завершение всех горутин при завершении основной горутины
func main() {
	go func() {
		for {
			fmt.Println("Hello from g")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
}
