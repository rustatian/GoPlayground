package main

func main() {

}

//go:noinline
func A(a ...*int) int {
	var res int
	for i := 0; i < len(a); i++ {
		res += *a[i]
	}

	return res
}

//go:noinline
func B(b []*int) int {
	var res int
	for i := 0; i < len(b); i++ {
		res += *b[i]
	}

	return res
}
