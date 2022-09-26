package main

import (
	"context"
	"fmt"
	"time"
)

func testGor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down from gor")
			return
		default:
			fmt.Println("Hello from gor")
			time.Sleep(time.Second)
		}
	}
}

// завершение горутины при помощи контекста: с таймаутом или с отменой
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel2()
	go testGor(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)
	go testGor(ctx2)
	time.Sleep(time.Second * 2)
}
