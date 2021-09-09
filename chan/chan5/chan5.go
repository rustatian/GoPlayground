package main

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	for i := 0; i < 10; i++ {
		j := <-ch

		println(j)

		ch <- j
	}
}
