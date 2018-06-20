package main

func main() {
	println(example2())
}

func example2() (int) {
	var i int
	defer func(){
		i = 3
	}()
	return 2
}
