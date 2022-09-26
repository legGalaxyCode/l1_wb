package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	counter uint32
}

// Increment по моему это хороший способ
func (c *Counter) Increment() {
	atomic.AddUint32(&c.counter, 1)
}

type CounterM struct {
	counter uint32
	m       sync.Mutex
}

// Increment этот хуже, но можно и так
func (c *CounterM) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.counter++
}

func main() {
	counter := Counter{counter: 0}
	counterM := CounterM{
		counter: 0,
		m:       sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				counter.Increment()
				counterM.Increment()
				fmt.Println(counter.counter, counterM.counter)
			}
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("Final counter: %d\n", counter.counter)
	fmt.Printf("Final counter1: %d", counterM.counter)
}
