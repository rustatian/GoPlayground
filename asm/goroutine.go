package main

func main() {
	NonFunConvert("a")
}

func NonFunConvert(intr interface{}) string {
	return intr.(string)
}
