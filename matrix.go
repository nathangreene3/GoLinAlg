package main

// A Matrix is an array of Vectors.
type Matrix []Vector

// MakeMatrix generates an m-by-n matrix with entries defined by a function f.
func MakeMatrix(m, n int, f func(a, b int) float64) (A Matrix) {
	A = make(Matrix, m)
	for i := 0; i < m; i++ {
		A[i] = make(Vector, n)
		for j := 0; j < n; j++ {
			A[i][j] = f(i, j)
		}
	}
	return A
}

// Identity returns the m-by-n identity matrix.
func Identity(m, n int) (I Matrix) {
	return MakeMatrix(m, n, func(a, b int) float64 {
		if a == b {
			return 1
		}
		return 0
	})
}

// MAdd returns the sum of two matrices.
func MAdd(A, B Matrix) (C Matrix) {
	if len(A) == len(B) {
		for i := range A {
			if len(A[i]) != len(A[0]) || len(A[0]) != len(B[i]) {
				panic("matrices A and B must have the same number of columns")
			}
		}
		C = MakeMatrix(len(A), len(A[0]), func(i, j int) float64 { return A[i][j] + B[i][j] })
	} else {
		panic("matrices A and B must have the same number of rows")
	}
	return C
}

// Transpose returns the transpose of A.
func Transpose(A Matrix) (B Matrix) {
	B = MakeMatrix(len(A[0]), len(A), func(i, j int) float64 { return A[j][i] })
	return B
}
