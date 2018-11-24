package main

import "fmt"

func main() {
	var a, b S
	a.addr()
	b.addr()
}

type S struct {
}

func (s *S) addr() {
	fmt.Printf("%p\n", s)
}

//func Foo(i interface{}) {
//	print(i)
//}

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.
type contextKey struct {
	//name string
}

//func (k *contextKey) String() string { return "net/http context value " + k.name }

var (
	// ServerContextKey is a context key. It can be used in HTTP
	// handlers with context.WithValue to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	ServerContextKey = &contextKey{}

	// LocalAddrContextKey is a context key. It can be used in
	// HTTP handlers with context.WithValue to access the local
	// address the connection arrived on.
	// The associated value will be of type net.Addr.
	LocalAddrContextKey = &contextKey{}
)

//
//0x0000 00000 (main.go:3)	TEXT	"".main(SB), $32-0
//0x0000 00000 (main.go:3)	MOVQ	(TLS), CX
//0x0009 00009 (main.go:3)	CMPQ	SP, 16(CX)
//0x000d 00013 (main.go:3)	JLS	87
//0x000f 00015 (main.go:3)	SUBQ	$32, SP
//0x0013 00019 (main.go:3)	MOVQ	BP, 24(SP)
//0x0018 00024 (main.go:3)	LEAQ	24(SP), BP
//0x001d 00029 (main.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
//0x001d 00029 (main.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//0x001d 00029 (main.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//0x001d 00029 (main.go:4)	PCDATA	$2, $1
//0x001d 00029 (main.go:4)	PCDATA	$0, $0
//0x001d 00029 (main.go:4)	MOVQ	"".LocalAddrContextKey(SB), AX
//0x0024 00036 (main.go:4)	PCDATA	$2, $0
//0x0024 00036 (main.go:4)	PCDATA	$0, $1
//0x0024 00036 (main.go:4)	MOVQ	AX, ""..autotmp_4+16(SP)
//0x0029 00041 (main.go:4)	CALL	runtime.printlock(SB)
//0x002e 00046 (main.go:4)	PCDATA	$2, $1
//0x002e 00046 (main.go:4)	LEAQ	type.*"".contextKey(SB), AX
//0x0035 00053 (main.go:4)	PCDATA	$2, $0
//0x0035 00053 (main.go:4)	MOVQ	AX, (SP)
//0x0039 00057 (main.go:4)	PCDATA	$2, $1
//0x0039 00057 (main.go:4)	PCDATA	$0, $0
//0x0039 00057 (main.go:4)	MOVQ	""..autotmp_4+16(SP), AX
//0x003e 00062 (main.go:4)	PCDATA	$2, $0
//0x003e 00062 (main.go:4)	MOVQ	AX, 8(SP)
//0x0043 00067 (main.go:4)	CALL	runtime.printeface(SB)
//0x0048 00072 (main.go:4)	CALL	runtime.printunlock(SB)
//0x004d 00077 (<unknown line number>)	MOVQ	24(SP), BP
//0x0052 00082 (<unknown line number>)	ADDQ	$32, SP
//0x0056 00086 (<unknown line number>)	RET
//0x0057 00087 (<unknown line number>)	NOP
//0x0057 00087 (main.go:3)	PCDATA	$0, $-1
//0x0057 00087 (main.go:3)	PCDATA	$2, $-1
//0x0057 00087 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
//0x005c 00092 (main.go:3)	JMP	0

//0x0000 00000 (main.go:3)	TEXT	"".main(SB), $24-0
//0x0000 00000 (main.go:3)	MOVQ	(TLS), CX
//0x0009 00009 (main.go:3)	CMPQ	SP, 16(CX)
//0x000d 00013 (main.go:3)	JLS	77
//0x000f 00015 (main.go:3)	SUBQ	$24, SP
//0x0013 00019 (main.go:3)	MOVQ	BP, 16(SP)
//0x0018 00024 (main.go:3)	LEAQ	16(SP), BP
//0x001d 00029 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x001d 00029 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x001d 00029 (main.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//0x001d 00029 (main.go:4)	PCDATA	$2, $0
//0x001d 00029 (main.go:4)	PCDATA	$0, $0
//0x001d 00029 (main.go:4)	CALL	runtime.printlock(SB)
//0x0022 00034 (main.go:4)	PCDATA	$2, $1
//0x0022 00034 (main.go:4)	LEAQ	type."".contextKey(SB), AX
//0x0029 00041 (main.go:4)	PCDATA	$2, $0
//0x0029 00041 (main.go:4)	MOVQ	AX, (SP)
//0x002d 00045 (main.go:4)	PCDATA	$2, $1
//0x002d 00045 (main.go:4)	LEAQ	runtime.zerobase(SB), AX
//0x0034 00052 (main.go:4)	PCDATA	$2, $0
//0x0034 00052 (main.go:4)	MOVQ	AX, 8(SP)
//0x0039 00057 (main.go:4)	CALL	runtime.printeface(SB)
//0x003e 00062 (main.go:4)	CALL	runtime.printunlock(SB)
//0x0043 00067 (<unknown line number>)	MOVQ	16(SP), BP
//0x0048 00072 (<unknown line number>)	ADDQ	$24, SP
//0x004c 00076 (<unknown line number>)	RET
//0x004d 00077 (<unknown line number>)	NOP
//0x004d 00077 (main.go:3)	PCDATA	$0, $-1
//0x004d 00077 (main.go:3)	PCDATA	$2, $-1
//0x004d 00077 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
//0x0052 00082 (main.go:3)	JMP	0
