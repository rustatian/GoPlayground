package main

import "fmt"

type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}
func (p *message) print() string {
	return p.content
}

func main() {
	m := &message{content: "Hello"}

	defer func() {
		fmt.Print(m.print())
	}()

	m.set("World")


	defer func(mm string) {
		fmt.Println("Some message")
		fmt.Println(mm)
	}("Hello")
}
