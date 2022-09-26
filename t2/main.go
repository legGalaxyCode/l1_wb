package main

import "fmt"

func main() {
	values := [...]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, v := range values {
		go func(c chan int, val int) {
			c <- val * val
		}(ch, v)
	}
	for range values {
		fmt.Println(<-ch)
	}
}
