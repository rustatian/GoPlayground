package main

var aaa interface{}

type foo interface {
	FooErr()
	Foo()
}

type foos struct {
	a  string
	b  string
	c  string
	f  string
	g  string
	h  string
	hh string
	aa string
	bb string
	cc string
	d  struct {
		a string
		b string
		c string
	}

	foo
}

func main() {

}

type aa struct {
}

func (aa) FooErr() {
	panic("implement me")
}

func (aa) Foo() {
	panic("implement me")
}

//go:noinline
func FooV(s foos) {
	aaa = s.a + s.b + s.c + s.f + s.g + s.h + s.hh + s.aa + s.bb + s.cc + s.d.a + s.d.b + s.d.c

}

//go:noinline
func FooP(s *foos) {
	aaa = s.a + s.b + s.c + s.f + s.g + s.h + s.hh + s.aa + s.bb + s.cc + s.d.a + s.d.b + s.d.c
}
