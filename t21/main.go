package main

import (
	"container/list"
	"fmt"
)

// adapter pattern
/*
	для меня классическим примером паттера адаптер являются очередь и стек
	т.к. они адаптируют контейнер (list, vector, deque) под семантику FIFO/LIFO
*/

// IntStack is an LIFO adapter for the list interface
type IntStack struct {
	buf list.List
}

func (st *IntStack) Top() int {
	return st.buf.Front().Value.(int)
}

func (st *IntStack) Size() int {
	return st.buf.Len()
}

func (st *IntStack) Empty() bool {
	return st.Size() == 0
}

func (st *IntStack) Push(value int) {
	st.buf.PushFront(value)
}

func (st *IntStack) Pop() {
	st.buf.Remove(st.buf.Front())
}

func makeIntStack() IntStack {
	return IntStack{}
}

func main() {
	stack := makeIntStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Printf("Stack size: %d\n", stack.Size())
	for !stack.Empty() {
		fmt.Printf("Stack top element: %d\n", stack.Top())
		stack.Pop()
	}
	fmt.Printf("Stack size: %d\n", stack.Size())
}
