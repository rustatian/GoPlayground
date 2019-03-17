package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func remove(ctx context.Context, cli *client.Client, id string) {
	fmt.Print("Removing container ", id[:10], "... ")
	err := cli.ContainerRemove(ctx,id, types.ContainerRemoveOptions{Force:true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}
