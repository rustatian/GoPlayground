package main

func main() {
	data := make([]int, 5)
	data[0] = 7
	data[1] = 8
	data[2] = 9
	data[3] = 10
	data[4] = 11
	bts := data
	data = changeData(data)
	println(bts)
	println(data)
}

func changeData(data []int) []int {
	data = changeData2(data)
	return data
}

func changeData2(data []int) []int {
	data = data[:1]
	return data
}
