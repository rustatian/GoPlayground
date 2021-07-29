package main

import "crypto/rand"

func main() {

}

type Payload struct {
	Context []byte
	Body    []byte
}

type PayloadStr struct {
	Context string
	Body    string
}

//go:noinline
func One() int {
	a := new(int)
	for i := 0; i < 1000; i++ {
		a2 := new(int)
		*a2 = *a + 1
		a = a2
	}

	return *a
}

//go:noinline
func Two() {
	for i := 0; i < 1000; i++ {
		a := 100
		_ = a
	}
}

//go:noinline
func fooPtr(p *Payload) *Payload {
	a := len(p.Body)
	b := len(p.Context)

	_ = a
	_ = b

	newPtr := &Payload{}
	b1, b2 := generate()

	newPtr.Body = b1
	newPtr.Context = b2

	return newPtr
}

//go:noinline
func fooVal(p Payload) Payload {
	a := len(p.Body)
	b := len(p.Context)

	_ = a
	_ = b

	newVal := Payload{}
	b1, b2 := generate()

	newVal.Body = b1
	newVal.Context = b2

	return newVal
}

func fooPtrVal(p *Payload) Payload {
	a := len(p.Body)
	b := len(p.Context)

	_ = a
	_ = b

	newPtr := Payload{}
	b1, b2 := generate()

	newPtr.Body = b1
	newPtr.Context = b2

	return newPtr
}

func fooValPtr(p Payload) *Payload {
	a := len(p.Body)
	b := len(p.Context)

	_ = a
	_ = b

	newPtr := &Payload{}
	b1, b2 := generate()

	newPtr.Body = b1
	newPtr.Context = b2

	return newPtr
}

//go:noinline
func generate() ([]byte, []byte) {
	b1 := make([]byte, 0, 1000)
	b2 := make([]byte, 0, 1000)

	_, err := rand.Read(b1)
	if err != nil {
		panic(err)
	}

	_, err = rand.Read(b2)
	if err != nil {
		panic(err)
	}

	return b1, b2
}

//go:noinline
func fooPtr1(p *Payload) int {
	a := len(p.Body)
	b := len(p.Context)

	return a + b
}

//go:noinline
func fooVal1(p Payload) int {
	a := len(p.Body)
	b := len(p.Context)

	return a + b
}
