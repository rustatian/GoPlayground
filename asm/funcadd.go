package main

func main() {

}

func FuncAdd(x, y, z int) int {
	return x + y - z
}

func DoCallAdd() int {
	return FuncAdd(1, 2, 3)
}
