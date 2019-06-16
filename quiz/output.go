package main

func main() {
	var a,b *int
	var c = b

	switch c {
	case a:
		print("a")
	case b:
		print("b")
	default:
		print("c")

	}
}
