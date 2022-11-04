package driver

type Driver struct {
	A string
	B string
}

func (d *Driver) InitFromConfig(a, b string) error {
	d.A = a
	return nil
}

func (d *Driver) InitFromPipeline(c, dd string) error {
	d.B = dd
	return nil
}

func (d *Driver) Foo1() error {
	return nil
}

func (d *Driver) Foo2() error {
	return nil
}
