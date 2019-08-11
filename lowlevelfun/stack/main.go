package main

func main() {
	y := 6
	other(y)
	print(y)
}

//go:noinline
func other(x int) {
	x = 5
	print(x)
}
