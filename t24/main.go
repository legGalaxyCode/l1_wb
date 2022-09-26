package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func makePoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func EuclideanDistance(x, y Point) float64 {
	return math.Sqrt(math.Pow(y.x-x.x, 2) + math.Pow(y.y-x.y, 2))
}

func main() {
	a := makePoint(1.0, 2.0)
	b := makePoint(4.0, 6.0)
	fmt.Println(EuclideanDistance(a, b))
}
