package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	set := make(map[int][]float64)
	for i := range temps {
		group := 10 * int(temps[i]/10)
		fmt.Println(group, temps[i], int(temps[i]/10)) // для понимания, почему группы именно так собираются
		if set[group] == nil {
			set[group] = make([]float64, 0)
			set[group] = append(set[group], temps[i])
		} else {
			set[group] = append(set[group], temps[i])
		}
	}
	fmt.Println(set)
}
