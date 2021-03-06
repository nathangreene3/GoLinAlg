package main

import (
	"fmt"

	"github.com/nathangreene3/GoLinAlg/matrix"

	"github.com/nathangreene3/GoLinAlg/vector"
)

type thing struct {
	value1 string
	value2 *string
}

func main() {
	// swapRowsTest()
	gravity()
}

func gravity() {
	// m := 1.0
	r := vector.Vector{0.50, 0.87}
	v := vector.Vector{0.25, 0.97}
	a := vector.ScalarMultiply(-1, vector.Unit(r))
	dt := 0.01
	fmt.Println(r, v, a)
	for t := 0.0; t <= 1.0; t += dt {
		r = vector.Add(r, vector.ScalarMultiply(dt, v))
		v = vector.Add(v, vector.ScalarMultiply(dt, a))
		fmt.Println(t, r)
	}
}

func swapRowsTest() {
	m, n := 3, 4
	A := matrix.EmptyMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A[i][j] = float64(i + j)
		}
	}
	fmt.Println(A, matrix.SwapRows(A, 1, 2))
}
