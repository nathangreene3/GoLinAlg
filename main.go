package main

import (
	"fmt"
	"math/rand"
)

func main() {
	A := MakeMatrix(2, 3, func(i, j int) float64 { return rand.Float64() })
	// B := MakeMatrix(2, 3, func(i, j int) float64 { return rand.Float64() })
	// C := MAdd(A, B)
	fmt.Println(A)
	// fmt.Println(B)
	// fmt.Println(C)
	fmt.Println(Transpose(A))
}
