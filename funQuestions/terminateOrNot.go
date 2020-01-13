package main

func main() {
	for {
		defer func() {
			for{}
		}()
		panic("yolo")
	}

}