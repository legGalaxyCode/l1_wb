package main

import (
	"fmt"
	"time"
)

func mySleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	fmt.Println(time.Now())
	mySleep(time.Second * 3)
	fmt.Println(time.Now())
}
