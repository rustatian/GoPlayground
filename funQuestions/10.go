package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(os.TempDir())
	err := eFoo()
	// if err != nil

	err = eFoo2()

	fmt.Print(err)
}

func eFoo() error {
	return fmt.Errorf("error1")
}

func eFoo2() error {
	return fmt.Errorf("error2")
}
