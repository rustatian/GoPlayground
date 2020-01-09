package main


func main() {
	a:= make([]int, 10)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	for i := range a {
		aa := &a[i]
		println(*aa)
	}
}

