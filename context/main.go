package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")

	ctx2, c := context.WithCancel(ctx)
	defer c()

	fmt.Println(ctx2.Value("key"))
}
