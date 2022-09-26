package main

import (
	"fmt"
)

func main() {
	values := [...]int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for _, v := range values {
		go func(ch chan int, v int) {
			ch <- v * v
		}(ch, v)
	}
	result := 0
	for range values {
		result += <-ch
	}
	fmt.Println(result)
}
