package main

import "fmt"

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
	if 0 < m && 0 < n {
		I = MakeMatrix(m, n, func(a, b int) (v float64) {
			if a == b {
				v = 1
			}
			return v
		})
	} else {
		panic("dimensions m and n must be positive, non-zero integers")
	}
	return I
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

// MMultiply returns C = A*B.
func MMultiply(A, B Matrix) (C Matrix) {
	if len(A[0]) == len(B) {
		C = MakeMatrix(len(A), len(B[0]), func(i, j int) float64 { return 0 })
		for i := range C {
			for j := range C[0] {
				for k := range A[0] {
					C[i][j] += A[i][k] * B[k][j]
				}
			}
		}
	} else {
		panic("A and B are of incompatible dimensions")
	}
	return C
}

// MPrint prints a formatted string representation of A.
func MPrint(A Matrix) {
	for i := range A {
		fmt.Print("[")
		for j := range A[0] {
			fmt.Print(A[i][j])
		}
		fmt.Println("]")
	}
	fmt.Println()
}

// MVectorToRowMatrix returns v converted to a 1-by-n matrix A.
func MVectorToRowMatrix(v Vector) (A Matrix) {
	A = MakeMatrix(1, len(v), func(i, j int) float64 { return v[j] })
	return A
}

// MVectorToColumnMatrix returns v converted to an n-by-1 matrix A.
func MVectorToColumnMatrix(v Vector) (A Matrix) {
	A = MakeMatrix(len(v), 1, func(i, j int) float64 { return v[i] })
	return A
}

// MGaussianElim returns x = inv(A)*b.
func MGaussianElim(A Matrix) (B Matrix) {
	B = MakeMatrix(len(A), len(A[0]), func(i, j int) float64 { return A[i][j] })
	return B
}
