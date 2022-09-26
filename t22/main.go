package main

import (
	"fmt"
	"math/big"
)

// we can use math.BigInt

func main() {
	a := big.NewInt(1 << 29)
	b := big.NewInt(1 << 30)

	op := big.NewInt(0)

	fmt.Printf("A= %d\nB= %d\n", a, b)
	op.Add(a, b)
	fmt.Printf("A+B = %d\n", op)
	op.Sub(a, b)
	fmt.Printf("A-B = %d\n", op)
	op.Mul(a, b)
	fmt.Printf("A*B = %d\n", op)
	op.Quo(b, a)
	fmt.Printf("B/A = %d\n", op)
}
