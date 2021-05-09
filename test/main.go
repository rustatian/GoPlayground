package main

func main() {

}

//go:noinline
func test_btb(bp bool) {
	if bp {
		println("hello, I'm function")
	}
}
