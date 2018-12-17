package main

func main() {
	print(xored())
}

//go:noinline
func xored() float64 {
	return ^float64(0)
}
