package main

import (
	"flag"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
	"sync"
)

func main() {
	fs := flag.NewFlagSet("dockerMegaUtility", flag.ExitOnError)

	nrun := fs.Bool("run", true, "run all containers")
	nstop := fs.Bool("stop", true, "stop all containers")
	ndelete := fs.Bool("delete", true, "delete all containers")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	if *nstop {
		stop(ctx, cli)
		//return
	}

	if *ndelete {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
		if err != nil {
			panic(err)
		}

		for _, cont := range containers {
			remove(ctx, cli, cont.ID)
		}

		//return
	}

	wg := &sync.WaitGroup{}
	if *nrun {
		wg.Add(6)
		go rabbit(ctx, cli, wg)
		go zipkin(ctx, cli, wg)
		go consul(ctx, cli, wg)
		go postgresql(ctx, cli, wg)
		go redis(ctx, cli, wg)
		go elastic(ctx, cli, wg)
	}

	wg.Wait()

}
