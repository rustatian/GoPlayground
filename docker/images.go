package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
	"sync"
)

//docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672  --restart always --name rabbit rabbitmq:3-management
func rabbit(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/library/rabbitmq:3.7.8-management", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)
	fmt.Println("Image rabbitmq pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["5672/tcp"] = []nat.PortBinding{{HostPort: "5672", HostIP: ":"}}
	bindings["15672/tcp"] = []nat.PortBinding{{HostPort: "15672", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "my-rabbit",
		Image: "rabbitmq:3.7.8-management",
		Tty:   true,
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "rabbit")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	fmt.Println("Started successfully RABBITMQ")
	wg.Done()
}

//docker run -d -p 9411:9411 --restart always openzipkin/zipkin
func zipkin(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/openzipkin/zipkin:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)
	fmt.Println("Image zipkin pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["9411/tcp"] = []nat.PortBinding{{HostPort: "9411", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "zipkin",
		Image: "openzipkin/zipkin:latest",
		Tty:   true,
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "zipkin")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
	fmt.Println("Started successfully ZIPKIN")
	wg.Done()
}

//docker run --restart always -d -p 8500:8500 consul agent -dev -client=0.0.0.0
func consul(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/library/consul:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)
	fmt.Println("Image consul pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["8500/tcp"] = []nat.PortBinding{{HostPort: "8500", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "consul",
		Image: "consul:latest",
		Cmd:   []string{"agent", "-dev", "-client=0.0.0.0",},
		Tty:   true,
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "consul")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	fmt.Println("Started successfully CONSUL")

	wg.Done()
}

//docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=1 --restart always -d postgres
func postgresql(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/library/postgres:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)
	fmt.Println("Image postgresql pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["5432/tcp"] = []nat.PortBinding{{HostPort: "5432", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "postgres",
		Image: "postgres:latest",
		Tty:   true,
		Env:   []string{"POSTGRES_PASSWORD=1"},
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "postgres")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	fmt.Println("Started successfully POSTGRESQL")
	wg.Done()
}

//docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:6.5.3
func elastic(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/library/elasticsearch:6.5.4", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)

	fmt.Println("Image elasticsearch pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["9200/tcp"] = []nat.PortBinding{{HostPort: "9200", HostIP: ":"}}
	bindings["9300/tcp"] = []nat.PortBinding{{HostPort: "9300", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "elastic",
		Image: "elasticsearch:6.5.4",
		Tty:   true,
		Env:   []string{"discovery.type=single-node"},
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "es")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	fmt.Println("Started successfully ELASTIC")
	wg.Done()
}

//docker run --name redis -p 6379:6379 --restart always -d redis
func redis(ctx context.Context, cli *client.Client, wg *sync.WaitGroup) {
	r, err := cli.ImagePull(ctx, "docker.io/library/redis:latest", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(devNull(0), r)
	fmt.Println("Image redis pulling done")

	bindings := map[nat.Port][]nat.PortBinding{}
	bindings["6379/tcp"] = []nat.PortBinding{{HostPort: "6379", HostIP: ":"}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{Hostname: "redis",
		Image: "redis:latest",
		Tty:   true,
	}, &container.HostConfig{RestartPolicy: container.RestartPolicy{Name: "always",}, PortBindings: bindings,}, &network.NetworkingConfig{}, "redis")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	fmt.Println("Started successfully REDIS")
	wg.Done()
}

type devNull int

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}
