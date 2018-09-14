package main

import (
	"fmt"
	"math/rand"
)

func main() {
	A := makeMatrix(2, 3, func(i, j int) float64 { return rand.Float64() })
	B := makeMatrix(2, 3, func(i, j int) float64 { return rand.Float64() })
	C := matrix.add(A, B)
	fmt.Println(C)
}
