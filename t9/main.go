package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := make(chan int)
	squareX := make(chan int)
	go func() {
		for _, val := range data {
			x <- val
		}
		close(x)
	}()
	go func() {
		for {
			val, ok := <-x
			if !ok {
				close(squareX)
				return
			}
			squareX <- 2 * val
		}
	}()
	for val := range squareX {
		fmt.Printf("Value from chan is: %d\n", val)
	}
}
