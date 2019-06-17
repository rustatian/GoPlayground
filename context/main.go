package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()

	ctxc, cancel := context.WithCancel(ctx)

	fmt.Print(ctxc)
	cancel()
}


