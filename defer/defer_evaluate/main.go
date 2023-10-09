package main

const X = 1

func main() {
	const (
		X = X + 1
		Y
	)

	println(X, Y)
}
