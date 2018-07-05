package main

func main() {
	print(fooo())

}

func fooo() (i int) {
	defer func() {
		i = 3
	}()

	print("whoo")
	return 2
}
