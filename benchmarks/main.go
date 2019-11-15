package main

func main() {
	
}

//go:noinline
func One() int {
	a := new(int);
	for i := 0; i < 1000; i++ {
		a2 := new(int)
		*a2 = *a + 1;
		a = a2
	}

	return *a
}

//go:noinline
func Two() {
	for i := 0; i < 1000; i++ {
		a := 100
		_ = a;
	}
}
