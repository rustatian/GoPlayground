package main

func main() {
	var a, b *int
	var c = b

	switch c {
	case a:
		print("a")
		fallthrough
	case b:
		print("b")
		fallthrough
	default:
		print("c")

	}
}
