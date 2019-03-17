package inter

func main() {

}

//with &
//"".main STEXT size=74 args=0x0 locals=0x18
//0x0000 00000 (ifacetask.go:3)	TEXT	"".main(SB), $24-0
//0x0000 00000 (ifacetask.go:3)	MOVQ	(TLS), CX
//0x0009 00009 (ifacetask.go:3)	CMPQ	SP, 16(CX)
//0x000d 00013 (ifacetask.go:3)	JLS	67
//0x000f 00015 (ifacetask.go:3)	SUBQ	$24, SP
//0x0013 00019 (ifacetask.go:3)	MOVQ	BP, 16(SP)
//0x0018 00024 (ifacetask.go:3)	LEAQ	16(SP), BP
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
//0x001d 00029 (ifacetask.go:4)	PCDATA	$2, $1
//0x001d 00029 (ifacetask.go:4)	PCDATA	$0, $0
//0x001d 00029 (ifacetask.go:4)	MOVQ	"".someVar(SB), AX
//0x0024 00036 (ifacetask.go:4)	PCDATA	$2, $2
//0x0024 00036 (ifacetask.go:4)	LEAQ	type.*"".foo(SB), CX
//0x002b 00043 (ifacetask.go:4)	PCDATA	$2, $1
//0x002b 00043 (ifacetask.go:4)	MOVQ	CX, (SP)
//0x002f 00047 (ifacetask.go:4)	PCDATA	$2, $0
//0x002f 00047 (ifacetask.go:4)	MOVQ	AX, 8(SP)
//0x0034 00052 (ifacetask.go:4)	CALL	"".InterfaceFunc(SB)
//0x0039 00057 (ifacetask.go:5)	MOVQ	16(SP), BP
//0x003e 00062 (ifacetask.go:5)	ADDQ	$24, SP
//0x0042 00066 (ifacetask.go:5)	RET
//0x0043 00067 (ifacetask.go:5)	NOP
//0x0043 00067 (ifacetask.go:3)	PCDATA	$0, $-1
//0x0043 00067 (ifacetask.go:3)	PCDATA	$2, $-1
//0x0043 00067 (ifacetask.go:3)	CALL	runtime.morestack_noctxt(SB)
//0x0048 00072 (ifacetask.go:3)	JMP	0

//w/o &
//"".main STEXT size=96 args=0x0 locals=0x28
//0x0000 00000 (ifacetask.go:3)	TEXT	"".main(SB), $40-0
//0x0000 00000 (ifacetask.go:3)	MOVQ	(TLS), CX
//0x0009 00009 (ifacetask.go:3)	CMPQ	SP, 16(CX)
//0x000d 00013 (ifacetask.go:3)	JLS	89
//0x000f 00015 (ifacetask.go:3)	SUBQ	$40, SP
//0x0013 00019 (ifacetask.go:3)	MOVQ	BP, 32(SP)
//0x0018 00024 (ifacetask.go:3)	LEAQ	32(SP), BP
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$1, gclocals·db9a6544d085c0622e79e6568b99b095(SB)
//0x001d 00029 (ifacetask.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//0x001d 00029 (ifacetask.go:4)	PCDATA	$2, $1
//0x001d 00029 (ifacetask.go:4)	PCDATA	$0, $0
//0x001d 00029 (ifacetask.go:4)	MOVQ	"".someVar(SB), AX
//0x0024 00036 (ifacetask.go:4)	MOVQ	"".someVar+8(SB), CX
//0x002b 00043 (ifacetask.go:4)	PCDATA	$2, $0
//0x002b 00043 (ifacetask.go:4)	PCDATA	$0, $1
//0x002b 00043 (ifacetask.go:4)	MOVQ	AX, ""..autotmp_0+16(SP)
//0x0030 00048 (ifacetask.go:4)	MOVQ	CX, ""..autotmp_0+24(SP)
//0x0035 00053 (ifacetask.go:4)	PCDATA	$2, $1
//0x0035 00053 (ifacetask.go:4)	LEAQ	type."".foo(SB), AX
//0x003c 00060 (ifacetask.go:4)	PCDATA	$2, $0
//0x003c 00060 (ifacetask.go:4)	MOVQ	AX, (SP)
//0x0040 00064 (ifacetask.go:4)	PCDATA	$2, $1
//0x0040 00064 (ifacetask.go:4)	LEAQ	""..autotmp_0+16(SP), AX
//0x0045 00069 (ifacetask.go:4)	PCDATA	$2, $0
//0x0045 00069 (ifacetask.go:4)	MOVQ	AX, 8(SP)
//0x004a 00074 (ifacetask.go:4)	CALL	"".InterfaceFunc(SB)
//0x004f 00079 (ifacetask.go:5)	MOVQ	32(SP), BP
//0x0054 00084 (ifacetask.go:5)	ADDQ	$40, SP
//0x0058 00088 (ifacetask.go:5)	RET
//0x0059 00089 (ifacetask.go:5)	NOP
//0x0059 00089 (ifacetask.go:3)	PCDATA	$0, $-1
//0x0059 00089 (ifacetask.go:3)	PCDATA	$2, $-1
//0x0059 00089 (ifacetask.go:3)	CALL	runtime.morestack_noctxt(SB)
//0x005e 00094 (ifacetask.go:3)	JMP	0
