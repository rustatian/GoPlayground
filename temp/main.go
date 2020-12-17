package main

type err string

func (e err) Error() string { return string(e) }

const FooErr = "foo"

type err2 string

func (e err2) Error() string {
	return string(e)
}

var FooErr2 = err2("foo")

func main() {
	println(FooErr)
	println(FooErr2)
}
