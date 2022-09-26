package main

import "fmt"

type Human struct {
}

func (h *Human) SayHello() {
	fmt.Println("Hello")
}

type Action struct {
	Human
}

func main() {
	var a Action
	a.SayHello()
}
