package main

func main() {
	//var ss []string

	do_primitives(nil)
	//do_primitives();
}

//go:noinline
func do_primitives(a ...*string) {
	for k, v := range a {
		print(k)
		print(v)
	}
}
