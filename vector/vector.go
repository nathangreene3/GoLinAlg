package vector

import (
	"fmt"
	"strings"
)

// A Vector is an array of numbers
type Vector []float64

// String returns a formatted string representation of a vector.
func (v Vector) String() string {
	n := len(v) // Length of v
	s := make([]string, 0, n+1)

	s = append(s, fmt.Sprintf("[%0.3f", v[0]))
	for i := 1; i < n; i++ {
		s = append(s, fmt.Sprintf(",%0.3f", v[i]))
	}
	s = append(s, "]")

	return strings.Join(s, "")
}

// Add returns w = u+v. Panics if |u| != |v|.
func Add(u, v Vector) Vector {
	n := len(u)
	if n != len(v) {
		panic("Vectors u and v must be of the same length")
	}

	w := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		w = append(w, u[i]+v[i])
	}
	return w
}

// Subtract returns w = u-v. Panics if |u| != |v|.
func Subtract(u, v Vector) Vector {
	return Add(u, ScalarMultiply(-1, v))
}

// ScalarMultiply returns v = a*u.
func ScalarMultiply(a float64, u Vector) Vector {
	v := make(Vector, 0, len(u))
	for i := range u {
		v = append(v, a*u[i])
	}
	return v
}

// Multiply returns w = u*v. Panics if |u| != |v|.
func Multiply(u, v Vector) Vector {
	n := len(u)
	if n != len(v) {
		panic("vectors u and v must be of the same length")
	}

	w := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		w = append(w, u[i]*v[i])
	}
	return w
}

// Sum returns the sum of the components of v.
func Sum(v Vector) float64 {
	var s float64
	for i := range v {
		s += v[i]
	}
	return s
}

// Dot returns p = sum u*v.
func Dot(u, v Vector) float64 {
	return Sum(Multiply(u, v))
}

// Mean returns the mean of the components of v.
func Mean(v Vector) float64 {
	return Sum(v) / float64(len(v))
}

// Less returns u < v. If |u| = |v| = 0, then false is returned. Panics
// if |u| != |v|.
func Less(u, v Vector) bool {
	n := len(u) // Length of vectors
	if n != len(v) {
		panic("")
	}

	if n == 0 {
		return false // Zero vectors are equal
	}
	for i := 0; i < n; i++ {
		if v[i] <= u[i] {
			return false
		}
	}
	return true
}

// Equal returns u = v. If |u| = |v| = 0, then true is returned. Panics
// if |u| != |v|.
func Equal(u, v Vector) bool {
	n := len(u)
	if n != len(v) {
		panic("vectors u and v must be of the same length")
	}

	if n == 0 {
		return true
	}
	for i := 0; i < n; i++ {
		if u[i] != v[i] {
			return false
		}
	}
	return true
}
