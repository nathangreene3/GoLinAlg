package matrix

import (
	"strings"

	"github.com/nathangreene3/GoLinAlg/vector"
)

// A Matrix is an array of Vectors.
type Matrix []vector.Vector

// MakeMatrix generates an m-by-n matrix with entries defined by a function f.
func MakeMatrix(m, n int, f func(a, b int) float64) Matrix {
	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vector.Vector, 0, n))
		for j := 0; j < n; j++ {
			A[i] = append(A[i], f(i, j))
		}
	}
	return A
}

// Identity returns the m-by-n identity matrix.
func Identity(m, n int) Matrix {
	if m < 0 || n < 0 {
		panic("Identity: dimensions m and n must be positive integers")
	}

	I := MakeMatrix(
		m,
		n,
		func(a, b int) float64 {
			if a == b {
				return 1
			}
			return 0
		},
	)
	return I
}

// Add returns the sum of two matrices.
func Add(A, B Matrix) Matrix {
	m := len(A)
	if m != len(B) {
		panic("matrices A and B must have the same number of rows")
	}

	n := len(A[0])
	for i := range A {
		if len(A[i]) != n || len(B[i]) != n {
			panic("matrices A and B must have the same number of columns")
		}
	}

	return MakeMatrix(m, n, func(i, j int) float64 { return A[i][j] + B[i][j] })
}

// Transpose returns the transpose of a matrix.
func Transpose(A Matrix) Matrix {
	return MakeMatrix(len(A[0]), len(A), func(i, j int) float64 { return A[j][i] })
}

// Multiply returns C = A*B. To multiply by a vector, convert the vector to a column matrix.
func Multiply(A, B Matrix) Matrix {
	if len(A[0]) != len(B) {
		panic("A and B are of incompatible dimensions") // Columns of A don't match rows of B
	}

	C := MakeMatrix(len(A), len(B[0]), func(i, j int) float64 { return 0 })
	for i := range C {
		for j := range C[0] {
			for k := range A[0] {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// String returns a formatted string representation of a matrix.
func (A Matrix) String() string {
	s := make([]string, 0, len(A)+1)

	s = append(s, "["+A[0].String())
	for i := 1; i < len(A); i++ {
		s = append(s, ","+A[i].String())
	}
	s = append(s, "]")

	return strings.Join(s, "")
}

// RowMatrix converts a vector v to a 1-by-n matrix.
func RowMatrix(v vector.Vector) Matrix {
	return MakeMatrix(1, len(v), func(i, j int) float64 { return v[j] })
}

// ColumnMatrix converts a vector v to an n-by-1 matrix.
func ColumnMatrix(v vector.Vector) Matrix {
	return MakeMatrix(len(v), 1, func(i, j int) float64 { return v[i] })
}
