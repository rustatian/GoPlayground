package main

func main() {
	y := 6
	other(y)
	print(y)
}

//go:noinline
func other(x int) {
	go func(i int) {
		print(i)
	}(x)
}
