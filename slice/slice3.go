package main

func main() {
	a := make([]int, 10)
	a[0] = 10
	b := a
	b[0] = 11

	println(a[0])
	println(b[0])

}
