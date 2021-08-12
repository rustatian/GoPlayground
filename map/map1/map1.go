package main

func main() {
	m := make(map[int]int, 10)
	m[1] = 1
	delete(m, 1)
}
