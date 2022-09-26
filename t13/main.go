package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
