package main

import (
	"fmt"
	"sync"
	"time"
)

type IntStringMap struct {
	data map[int]string
	mutx sync.RWMutex
}

func (m *IntStringMap) Load(key int) (string, bool) {
	m.mutx.RLock()
	defer m.mutx.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *IntStringMap) Store(key int, value string) {
	m.mutx.Lock()
	defer m.mutx.Unlock()
	m.data[key] = value
}

func MakeMap() IntStringMap {
	return IntStringMap{
		data: make(map[int]string),
		mutx: sync.RWMutex{},
	}
}

func main() {
	myMap := MakeMap()
	wg := sync.WaitGroup{}

	// write first
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *IntStringMap, key int, value string) {
			m.Store(key, value)
			wg.Done()
		}(&wg, &myMap, i, fmt.Sprintf("value%d", i))
	}
	wg.Wait() // we need to wait all writing

	// read unique now
	for i := 0; i < 4; i++ {
		go func(m *IntStringMap, key int) {
			if data, ok := m.Load(key); ok {
				fmt.Printf("Read value: %s with key %d from map\n", data, key)
			} else {
				fmt.Println("Map doesn't contain value with this key")
			}
		}(&myMap, i)
	}
	// read same
	for i := 0; i < 3; i++ {
		go func(m *IntStringMap, key int, id int) {
			if data, ok := m.Load(key); ok {
				fmt.Printf("Read value: %s with key %d from g %d\n", data, key, id)
			} else {
				fmt.Println("Map doesn't contain value with this key")
			}
		}(&myMap, 1, i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Shutting down from main...")
}
