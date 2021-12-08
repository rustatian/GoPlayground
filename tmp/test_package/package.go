package test_package

import (
	"fmt"
)

var MAP = map[string]interface{}{
	"a": &ConfigA{},
}

type ConfigA struct {
	Foo string
}

type Plugin struct {
	mm string
}

//go:noinline
func (p *Plugin) Init() {
	a := MAP["a"]
	var bb = a.(*ConfigA)
	fmt.Println(bb.Foo)
	bb.Foo = "foo"
	MAP["a"] = bb
}
