package main

import (
	"fmt"
	"time"
)

// завершение горутины с использованием второго параметра при чтении из канала
// если канал закрыт, то выходим из горутины
func main() {
	data := make(chan int)
	go func(data <-chan int) {
		for {
			if val, ok := <-data; !ok {
				fmt.Println("Channel was closed, g is shutting down...")
				return // или runtime.Goexit() которая завершает горутину
			} else {
				fmt.Printf("Value: %d\n", val)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}(data)
	for i := 0; i < 10; i++ {
		data <- i
	}
	fmt.Println("Closing channel")
	close(data)
	time.Sleep(time.Second)
}
