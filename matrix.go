package main

type matrix []vector

func makeMatrix(m, n int, f func(a, b int) float64) (A matrix) {
	A = make(matrix, m)
	for i := 0; i < m; i++ {
		A[i] = make(vector, n)
		for j := 0; j < n; j++ {
			A[i][j] = f(i, j)
		}
	}
	return A
}

func identity(m, n int) (I matrix) {
	return makeMatrix(m, n, func(a, b int) float64 {
		if a == b {
			return 1
		}
		return 0
	})
}

func add(A, B matrix) (C matrix) {
	if len(A) == len(B) {
		for i := range A {
			if len(A[i]) != len(A[0]) || len(A[0]) != len(B[i]) {
				panic("matrices A and B must have the same number of columns")
			}
		}
		C = makeMatrix(len(A), len(A[0]), func(i, j int) float64 { return A[i][j] + B[i][j] })
	} else {
		panic("matrices A and B must have the same number of rows")
	}
	return C
}
