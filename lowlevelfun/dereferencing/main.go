package main

func main() {
	deref()
}

//go:noinline
func deref() {
	var a *int
	print(*a)
}
