package main

type vector []float64

func add(u, v vector) (w vector) {
	if len(u) == len(v) {
		w = make(vector, len(u))
		for i := range u {
			w[i] = u[i] + v[i]
		}
	} else {
		panic("vectors u and v must be of the same length")
	}
	return w
}

func subtract(u, v vector) (w vector) {
	w = add(u, multiply(-1, v))
	return w
}

func multiply(a float64, u vector) (v vector) {
	if a != 0 {
		v = make(vector, len(u))
		for i := range u {
			v[i] = a * u[i]
		}
	} else {
		panic("scalar a must not be zero")
	}
	return v
}

func sum(v vector) (s float64) {
	for i := range v {
		s += v[i]
	}
	return s
}

func dotProduct(u, v vector) (p float64) {
	for i := range u {
		p += u[i] * v[i]
	}
	return p
}

func mean(v vector) (a float64) {
	a = sum(v) / float64(len(v))
	return a
}
