package main

//go:noinline
func Add(a, b int32) int32 {
	return a + b
}

type Adder struct {
	id int32
}

//go:noinline
func (adder *Adder) AddPtr(a, b int32) int32 {
	return a + b
}

//go:noinline
func (adder Adder) AddVal(a, b int32) int32 {
	return a + b
}

func main() {
	Add(10, 32)

	adder := Adder{id: 6754}
	adder.AddPtr(10, 32)
	adder.AddVal(10, 32)

	(&adder).AddVal(10, 32)
}
