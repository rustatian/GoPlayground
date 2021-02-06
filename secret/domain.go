package secret

type AppSpace interface {
	DoWork()
}

type Domain struct {
	// application layer
	app AppSpace
}

func NewDomainLayer(app AppSpace) *Domain {
	return &Domain{app: app}
}

func (d *Domain) DoWork() {
	// TODO domain checks
	// pass work to the application
	d.app.DoWork()
}
