package vector

import (
	"fmt"
	"math"
	"strings"
)

// A Vector is an array of numbers. The dimension refers to the number of entries
// in the vector.
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

// Add returns the vector addition of two vectors. Panics if the given vectors are
// not of equal dimension.
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

// Subtract returns the vector subtraction of two vectors. Panics if the given
// vectors are not of equal dimension.
func Subtract(u, v Vector) Vector {
	return Add(u, ScalarMultiply(-1, v))
}

// ScalarMultiply returns a vector with each entry defined as the product of the
// given vector and a given scalar.
func ScalarMultiply(a float64, u Vector) Vector {
	v := make(Vector, 0, len(u))
	for i := range u {
		v = append(v, a*u[i])
	}
	return v
}

// Multiply returns a vector with each entry defined as the product of the given
// vectors' entries. Panics if the vectors are not of equal dimension.
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

// Sum returns the sum of the components of a vector.
func Sum(v Vector) float64 {
	var s float64
	for i := range v {
		s += v[i]
	}
	return s
}

// Dot returns the dot product of two vectors.
func Dot(u, v Vector) float64 {
	return Sum(Multiply(u, v))
}

// Cross returns the cross product of two three-dimensional vectors.
func Cross(u, v Vector) Vector {
	if len(u) != 3 || len(v) != 3 {
		panic("vectors must be of dimension three")
	}

	return Vector{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
}

// Mean returns the mean of the components of v.
func Mean(v Vector) float64 {
	return Sum(v) / float64(len(v))
}

// Length returns the geometric length of a vector. The result is the mathematical
// quantity and is NOT the number of dimensions of the vector. Use len(v) for the
// number of dimensions of a vector.
func Length(v Vector) float64 {
	return math.Sqrt(Dot(v, v))
}

// Unit returns the unit vector of a vector. The sum of the unit vector components
// add to one and the unit vector is parallel to the given vector.
func Unit(v Vector) Vector {
	length := Length(v)
	u := make(Vector, 0, len(v))
	for i := range v {
		u = append(u, v[i]/length)
	}
	return u
}

// AngleR returns the angle (in radians) between two vectors.
func AngleR(u, v Vector) float64 {
	return Dot(Unit(u), Unit(v))
}

// Proj returns the vector component of a vector u parallel to a non-zero vector v.
func Proj(u, v Vector) Vector {
	w := Unit(v)
	a := Dot(u, w)
	for i := range w {
		w[i] *= a
	}
	return w
}

// Less returns the less-than comparison of two vectors. If the number of
// dimensions is zero, then false is returned. Panics if the vectors hav
// different dimensions.
func Less(u, v Vector) bool {
	n := len(u)
	if n != len(v) {
		panic("Less: vectors must be of the same length")
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

// Equal returns the comparison of each entry in two given vectors. Panics if the
// vectors have different dimensions.
func Equal(u, v Vector) bool {
	n := len(u)
	if n != len(v) {
		panic("Equal: vectors must be of the same length")
	}

	for i := 0; i < n; i++ {
		if u[i] != v[i] {
			return false
		}
	}
	return true
}
