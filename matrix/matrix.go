package matrix

import (
	"strings"

	"github.com/nathangreene3/GoLinAlg/vector"
)

// A Matrix is an array of Vectors (rows).
type Matrix []vector.Vector

// MakeMatrix generates an m-by-n matrix with entries defined by a function f. The
// parameters a and b correspond to the (a,b) entry of the returned matrix. Panics
// if either m or n are less than one.
func MakeMatrix(m, n int, f func(a, b int) float64) Matrix {
	if m < 1 || n < 1 {
		panic("Identity: dimensions m and n must be positive integers")
	}

	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vector.Vector, 0, n))
		for j := 0; j < n; j++ {
			A[i] = append(A[i], f(i, j))
		}
	}
	return A
}

// EmptyMatrix returns an m-by-n matrix with zeroes for all entries.
func EmptyMatrix(m, n int) Matrix {
	return MakeMatrix(m, n, func(i, j int) float64 { return 0 })
}

// Identity returns the m-by-n identity matrix.
func Identity(m, n int) Matrix {
	return MakeMatrix(m, n, func(i, j int) float64 { return float64((i + j) % 2) })
}

// Add returns the sum of two matrices. Panics if the matrices are not equal in
// dimension.
func Add(A, B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	return MakeMatrix(ma, na, func(i, j int) float64 { return A[i][j] + B[i][j] })
}

// Subtract returns the subtraction of two matrices. Panics if the matrices are
// not equal in dimension.
func Subtract(A, B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	return MakeMatrix(ma, na, func(i, j int) float64 { return A[i][j] - B[i][j] })
}

// Transpose returns the transpose of a matrix.
func Transpose(A Matrix) Matrix {
	return MakeMatrix(len(A[0]), len(A), func(i, j int) float64 { return A[j][i] })
}

// Multiply returns C = A*B. To multiply by a vector, convert the vector to a
// column matrix.
func Multiply(A, B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("A and B are of incompatible dimensions") // Columns of A don't match rows of B
	}

	C := EmptyMatrix(ma, nb)
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
	sb := strings.Builder{}
	sb.WriteByte(byte('['))
	sb.WriteString(A[0].String())
	for i := 1; i < len(A); i++ {
		sb.WriteByte(',')
		sb.WriteString(A[i].String())
	}
	sb.WriteByte(']')
	return sb.String()
}

// Equals returns true if two matrices are equal in dimension and for each entry.
// Otherwise, it returns false.
func Equals(A, B Matrix) bool {
	// Compare dimensions
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		return false
	}

	// Compare rows
	for i := range A {
		if !vector.Equal(A[i], B[i]) {
			return false
		}
	}
	return true
}

// RowMatrix converts a vector v to a 1-by-n matrix.
func RowMatrix(v vector.Vector) Matrix {
	return MakeMatrix(1, len(v), func(i, j int) float64 { return v[j] })
}

// ColumnMatrix converts a vector v to an n-by-1 matrix.
func ColumnMatrix(v vector.Vector) Matrix {
	return MakeMatrix(len(v), 1, func(i, j int) float64 { return v[i] })
}

// Determinant returns the Determinant of a square matrix. Panics if matrix is
// empty or not a square.
func (A Matrix) Determinant() float64 {
	// TODO //
	m, n := A.Dimensions()
	if m*n < 1 {
		panic("cannot take determinant of empty matrix")
	}
	if m != n {
		panic("cannot take determinant of a non-square matrix")
	}

	switch m {
	case 1:
		return A[0][0]
	case 2:
		return A[0][0]*A[1][1] - A[0][1]*A[1][0]
	case 3:
		return A[0][0]*(A[1][2]*A[2][2]-A[1][2]*A[2][1]) - A[0][1]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) + A[0][2]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])
	default:
		// TODO //
		return 0
	}
}

// Dimensions returns the Dimensions (number of rows, number of columns) of a
// matrix. Panics if number of columns is not constant for each row.
func (A Matrix) Dimensions() (int, int) {
	m := len(A)
	n := len(A[0])
	for i := range A {
		if n != len(A[i]) {
			panic("inconsistent matrix dimensions")
		}
	}
	return m, n
}

// Copy returns a deep copied matrix.
func Copy(A Matrix) Matrix {
	m, n := A.Dimensions()
	return MakeMatrix(m, n, func(i, j int) float64 { return A[i][j] })
}

// Join returns a matrix that is the joining of two given matrices. Panics if number of rows are not equal.
func Join(A, B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb {
		panic("matrices must have equal number of rows")
	}

	return MakeMatrix(
		ma,
		na+nb,
		func(i, j int) float64 {
			if j < na {
				return A[i][j]
			}
			return B[i][j-na]
		},
	)
}

// AppendColumn returns a matrix that is the joining of a given matrix with a column Vector. Panics if vector dimensions are not equal to the number of matrix rows.
func AppendColumn(A Matrix, x vector.Vector) Matrix {
	return Join(A, ColumnMatrix(x))
}

// AppendRow returns a matrix that is the joining of a given matrix with a row vector. Panics if the vector dimensions are not equal to the number of matrix columns.
func AppendRow(A Matrix, x vector.Vector) Matrix {
	m, n := A.Dimensions()
	if n != len(x) {
		panic("matrix columns must be equal to vector dimensions")
	}

	return MakeMatrix(
		m+1,
		n,
		func(i, j int) float64 {
			if i < m {
				return A[i][j]
			}
			return x[j]
		},
	)
}

// SetDims returns a matrix with new dimensions. If the dimensions are larger than the dimensions of the given matrix, zeroes are entered in the larger rows and columns. That is, the given matrix is embedded within the new matrix.
func SetDims(A Matrix, m, n int) Matrix {
	// TODO: Consider renaming this function.
	ma, na := A.Dimensions()
	return MakeMatrix(
		m,
		n,
		func(a, b int) float64 {
			if a < ma && b < na {
				return A[a][b]
			}
			return 0
		},
	)
}

// Len returns the number of vectors in a given matrix.
func (A Matrix) Len() int {
	return len(A)
}

// Swap swaps two vectors in a matrix.
func (A Matrix) Swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}

// Less returns the comparison of two vectors in a matrix.
func (A Matrix) Less(i, j int) bool {
	return vector.Less(A[i], A[j])
}

// Bracket returns the bracket product [A,B] = AB - BA of two matrices. This may
// also be known as a Lie product.
func Bracket(A, B Matrix) Matrix {
	return Subtract(Multiply(A, B), Multiply(B, A))
}

// Solve TODO
func (A Matrix) Solve(y vector.Vector) vector.Vector {
	// 	// x := make(vector.Vector, 0, len(y))
	// 	// B := Join(A, ColumnMatrix(x))
	// 	// m, n := B.Dimensions()

	// 	// Sort on kth entry for each row
	// 	// sort.SliceStable(
	// 	// 	B,
	// 	// 	func(i, j int) bool {
	// 	// 		for k := range B[i] {
	// 	// 			if B[i][k] < B[j][k] {
	// 	// 				return true
	// 	// 			}
	// 	// 		}
	// 	// 		return false
	// 	// 	},
	// 	// )

	// 	// Iterate through each row
	// 	// for i := range B {
	// 	// 	// Divide each row by its diagonal to get 1 on the diagonal
	// 	// 	if B[i][i] == 0 {
	// 	// 		continue
	// 	// 	}
	// 	// 	for j := range B[i] {
	// 	// 		B[i][j] /= B[i][i]
	// 	// 	}

	// 	// 	// Divide each row by -B[i][i] except for ith row
	// 	// 	for j := range B {
	// 	// 		if j == i {
	// 	// 			continue
	// 	// 		}
	// 	// 		for k := range B[j] {
	// 	// 			B[j][k] = B[j][k]/-B[i][i] + B[i][k]
	// 	// 		}
	// 	// 	}
	// 	// }
	return y
}

// SwapRows swaps two rows.
func SwapRows(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	return MakeMatrix(
		m,
		n,
		func(a, b int) float64 {
			switch a {
			case i:
				return A[j][b]
			case j:
				return A[i][b]
			default:
				return A[a][b]
			}
		},
	)
}

// SwapCols swaps two columns.
func SwapCols(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	return MakeMatrix(
		m,
		n,
		func(a, b int) float64 {
			switch b {
			case i:
				return A[a][j]
			case j:
				return A[a][i]
			default:
				return A[a][b]
			}
		},
	)
}
