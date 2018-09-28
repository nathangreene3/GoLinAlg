package main

// A Vector is an array of numbers
type Vector []float64

// VAdd returns w = u+v.
func VAdd(u, v Vector) (w Vector) {
	if len(u) == len(v) {
		w = make(Vector, len(u))
		for i := range u {
			w[i] = u[i] + v[i]
		}
	} else {
		panic("Vectors u and v must be of the same length")
	}
	return w
}

// VSubtract returns w = u-v.
func VSubtract(u, v Vector) (w Vector) {
	w = VAdd(u, VScalarMultiply(-1, v))
	return w
}

// VScalarMultiply returns v = a*u.
func VScalarMultiply(a float64, u Vector) (v Vector) {
	if a != 0 {
		v = make(Vector, len(u))
		for i := range u {
			v[i] = a * u[i]
		}
	} else {
		panic("scalar a must not be zero")
	}
	return v
}

// VMultiply returns w = u*v.
func VMultiply(u, v Vector) (w Vector) {
	if len(u) == len(v) {
		w = make(Vector, len(u))
		for i := range u {
			w[i] = u[i] * v[i]
		}
	} else {
		panic("vectors u and v must be of the same length")
	}
	return w
}

// VSum returns the sum of the components of v.
func VSum(v Vector) (s float64) {
	for i := range v {
		s += v[i]
	}
	return s
}

// DotProduct returns p = sum u*v.
func DotProduct(u, v Vector) (p float64) {
	p = VSum(VMultiply(u, v))
	return p
}

// VMean returns the mean of the components of v.
func VMean(v Vector) (a float64) {
	a = VSum(v) / float64(len(v))
	return a
}

// VMatrixMultiply returns b = A*x.
func VMatrixMultiply(A Matrix, x Vector) (b Vector) {
	if len(A[0]) == len(x) {
		b = make(Vector, len(A))
		for i := range A {
			for j := range A[i] {
				b[i] += A[i][j] * x[j]
			}
		}
	} else {
		panic("matrix A[0] (width or n) and vector x must have the same length")
	}
	return b
}
