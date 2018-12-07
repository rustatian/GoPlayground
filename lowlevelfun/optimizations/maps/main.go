package main

import "runtime"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()

	for !done {
		runtime.Gosched()
	}
	println(a) // expect to print: hello, world
}


//package main

//func main() {
//	bt := []byte{'a', 'b', 'c'}
//	m := make(map[string]int, 10)
//
//	m[string(bt[0])] = 1
//	fmt.Println(m)
//}

//go:noinline
//func fun() {
//	bt := []byte{'a', 'b', 'c'}
//	m := make(map[string]int, 10)
//
//	m[string(bt[0])] = 1
//}

//go:noinline
//func notfun() {
//	bt := []byte{'a', 'b', 'c'}
//	m := make(map[string]int, 10)
//	tmp := string(bt[0])
//	m[tmp] = 1
//}

//pcdata  $2, $1
//leaq    type.map[string]int(SB), AX
//pcdata  $2, $0
//movq    AX, (SP)
//pcdata  $2, $2
//movq    "".m+56(SP), CX
//pcdata  $2, $0
//movq    CX, 8(SP)
//call    runtime.mapassign_faststr(SB)
//pcdata  $2, $1
//movq    32(SP), AX
//pcdata  $2, $0
//movq    $1, (AX)
//
//
//
//movblzx ""..autotmp_5+53(SP), CX
//movq    CX, 8(SP)
//movq    $0, (SP)
//call    runtime.intstring(SB)
//pcdata  $2, $1
//leaq    type.map[string]int(SB), AX
//pcdata  $2, $0
//movq    AX, (SP)
//pcdata  $2, $2
//movq    "".m+56(SP), CX
//pcdata  $2, $0
//movq    CX, 8(SP)
//call    runtime.mapassign_faststr(SB)
//pcdata  $2, $1
//movq    32(SP), AX
//pcdata  $2, $0
//movq    $1, (AX)