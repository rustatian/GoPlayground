package main

type foo struct {
	pSlice []*int
	vSlice []int
}

func main() {
	f := &foo{
		pSlice: make([]*int, 1e8),
		vSlice: make([]int, 1e8),
	}

}
