package main

func main() {
	aa := 1
	bb := 2
	cc := 3

	a := []*int{&aa, &bb, &cc}

	for _, v := range a {
		println(*v)
		*v = 16
	}

	for _, v := range a {
		println(*v)
	}
}
