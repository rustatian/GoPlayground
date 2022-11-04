package p1

import (
	"github.com/rustatian/GoPlayground/endure/p1/driver"
)

type Plugin struct {
}

func (p *Plugin) Init() error {
	return nil
}

func (p *Plugin) Name() string {
	return "p1"
}

// InitKVDriver ..
// *driver.Driver -> implements interface
func (p *Plugin) InitKVDriver() (*driver.Driver, error) {
	return &driver.Driver{}, nil
}
