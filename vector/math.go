package vector

// maxInt returns the maximum value in a set.
func maxInt(x ...int) int {
	var index int // Index of maximum value in x
	for i := range x {
		if x[index] < x[i] {
			index = i
		}
	}
	return x[index]
}

// maxFloat64 returns the maximum value in a set.
func maxFloat64(x ...float64) float64 {
	var index int // Index of maximum value in x
	for i := range x {
		if x[index] < x[i] {
			index = i
		}
	}
	return x[index]
}

// minInt returns the minimum value in a set.
func minInt(x ...int) int {
	var index int // Index of minimum value in x
	for i := range x {
		if x[i] < x[index] {
			index = i
		}
	}
	return x[index]
}

// minFloat64 returns the minimum value in a set.
func minFloat64(x ...float64) float64 {
	var index int // Index of minumum value in x
	for i := range x {
		if x[i] < x[index] {
			index = i
		}
	}
	return x[index]
}
