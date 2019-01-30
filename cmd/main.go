package main

import (
	"fmt"

	"github.com/nathangreene3/GoLinAlg/matrix"
	"github.com/nathangreene3/GoLinAlg/vector"
)

func main() {
	A := matrix.MakeMatrix(2, 3, func(i, j int) float64 { return float64(i + j + 1) })
	B := matrix.MakeMatrix(3, 1, func(i, j int) float64 { return float64(i + j + 1) })
	x := vector.Vector{1, 2, 3}
	fmt.Println(
		A.String(),
		B.String(),
		x.String(),
		matrix.Multiply(A, B),
		matrix.Multiply(A, matrix.ColumnMatrix(x)),
	)
	fmt.Println(matrix.MakeMatrix(2, 2, func(i, j int) float64 { return float64(i + j) }).Determinant())
}
