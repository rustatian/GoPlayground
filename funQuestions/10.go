package main

import (
	"fmt"
)

func main() {
	err := eFoo()

	err = eFoo2()

	fmt.Print(err)
}

func eFoo() error {
	return fmt.Errorf("error1")
}

func eFoo2() error {
	return fmt.Errorf("error2")
}
