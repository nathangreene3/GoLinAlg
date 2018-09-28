package main

func main() {
	A := MakeMatrix(2, 3, func(i, j int) float64 { return float64(i + j + 1) })
	B := MakeMatrix(3, 1, func(i, j int) float64 { return float64(i + j + 1) })
	MPrint(A)
	MPrint(B)
	MPrint(MMultiply(A, B))
}
