package main

func main() {

}

type Foo interface {
	State() int
}

type FooImpl struct{}

//go:noinline
func (f *FooImpl) State() int {
	return 1
}

//go:noinline
func and(a, b int) int {
	if (a & b) == 0 {
		return 0
	}

	return 1
}

//go:noinline
func eq(a, b int) int {
	if a == b {
		return 0
	}

	return 1
}
