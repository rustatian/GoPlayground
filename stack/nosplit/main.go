package main

type T [1024000000]byte

////go:nosplit
//func A(t T) {
//	B(t)
//}
//
////go:nosplit
//func B(t T) {
//	C(t)
//}
//
////go:nosplit
//func C(t T) {
//	D(t)
//}

//go:noinline
//go:nosplit
func D(t T) {}

func main() {
	var t T
	D(t)
}
