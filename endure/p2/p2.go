package p2

import (
	endure "github.com/roadrunner-server/endure/pkg/container"
)

type Fooer interface {
	InitFromConfig(a, b string) error
	InitFromPipeline(c, dd string) error
	Foo1() error
	Foo2() error
}

type Plugin struct {
}

func (p *Plugin) Init() error {
	return nil
}

func (p *Plugin) Collects() []any {
	return []any{
		p.Drivers,
	}
}

func (p *Plugin) Drivers(en endure.Named, fooer Fooer) {
	println(en.Name())
}
