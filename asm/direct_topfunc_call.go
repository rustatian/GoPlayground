package main

func main() {
	add(10, 32)
}

func add(a, b int32) (int32, bool) {
	return a + b, true
}
