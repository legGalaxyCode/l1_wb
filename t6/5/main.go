package main

import (
	"fmt"
	"sync"
	"time"
)

// мы можем ждать завершения всех горутин при помощи вейт группы
// но wg.Done() не завершает текущую горутину
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Hello from g with wait group")
		time.Sleep(time.Second)
	}(&wg)

	wg.Wait()
	fmt.Println("All g are done")
}
