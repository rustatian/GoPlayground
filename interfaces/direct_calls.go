package inter

//func main() {
//	Add(10, 32)
//
//	adder := Adder{id: 6754}
//	adder.AddPtr(10, 32)
//	adder.AddVal(10, 32)
//	(&adder).AddVal(10, 32)
//}

////go:noinline
func Add(a, b int32) int32 {
	return a + b
}

type Adder struct {
	id int32
}

////go:noinline
func (adder *Adder) AddPtr(a, b int32) int32 {
	return a + b
}

////go:noinline
func (adder Adder) AddVal(a, b int32) int32 {
	return a + b
}

//"".main STEXT size=171 args=0x0 locals=0x28
//0x0000 00000 (direct_calls.go:3)	TEXT	"".main(SB), $40-0
//0x0000 00000 (direct_calls.go:3)	MOVQ	(TLS), CX
//0x0009 00009 (direct_calls.go:3)	CMPQ	SP, 16(CX)
//0x000d 00013 (direct_calls.go:3)	JLS	161
//0x0013 00019 (direct_calls.go:3)	SUBQ	$40, SP
//0x0017 00023 (direct_calls.go:3)	MOVQ	BP, 32(SP)
//0x001c 00028 (direct_calls.go:3)	LEAQ	32(SP), BP
//0x0021 00033 (direct_calls.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x0021 00033 (direct_calls.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x0021 00033 (direct_calls.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//0x0021 00033 (direct_calls.go:4)	PCDATA	$2, $0
//0x0021 00033 (direct_calls.go:4)	PCDATA	$0, $0
//0x0021 00033 (direct_calls.go:4)	MOVQ	$137438953482, AX
//0x002b 00043 (direct_calls.go:4)	MOVQ	AX, (SP)
//0x002f 00047 (direct_calls.go:4)	CALL	"".Add(SB)
//0x0034 00052 (direct_calls.go:6)	MOVL	$0, "".adder+28(SP)
//0x003c 00060 (direct_calls.go:6)	MOVL	$6754, "".adder+28(SP)
//0x0044 00068 (direct_calls.go:7)	PCDATA	$2, $1
//0x0044 00068 (direct_calls.go:7)	LEAQ	"".adder+28(SP), AX
//0x0049 00073 (direct_calls.go:7)	PCDATA	$2, $0
//0x0049 00073 (direct_calls.go:7)	MOVQ	AX, (SP)
//0x004d 00077 (direct_calls.go:7)	MOVQ	$137438953482, AX
//0x0057 00087 (direct_calls.go:7)	MOVQ	AX, 8(SP)
//0x005c 00092 (direct_calls.go:7)	CALL	"".(*Adder).AddPtr(SB)
//0x0061 00097 (direct_calls.go:8)	MOVL	"".adder+28(SP), AX
//0x0065 00101 (direct_calls.go:8)	MOVL	AX, (SP)
//0x0068 00104 (direct_calls.go:8)	MOVQ	$137438953482, AX
//0x0072 00114 (direct_calls.go:8)	MOVQ	AX, 4(SP)
//0x0077 00119 (direct_calls.go:8)	CALL	"".Adder.AddVal(SB)
//0x007c 00124 (direct_calls.go:9)	MOVL	"".adder+28(SP), AX
//0x0080 00128 (direct_calls.go:9)	MOVL	AX, (SP)
//0x0083 00131 (direct_calls.go:9)	MOVQ	$137438953482, AX
//0x008d 00141 (direct_calls.go:9)	MOVQ	AX, 4(SP)
//0x0092 00146 (direct_calls.go:9)	CALL	"".Adder.AddVal(SB)
//0x0097 00151 (direct_calls.go:10)	MOVQ	32(SP), BP
//0x009c 00156 (direct_calls.go:10)	ADDQ	$40, SP
//0x00a0 00160 (direct_calls.go:10)	RET
//0x00a1 00161 (direct_calls.go:10)	NOP
//0x00a1 00161 (direct_calls.go:3)	PCDATA	$0, $-1
//0x00a1 00161 (direct_calls.go:3)	PCDATA	$2, $-1
//0x00a1 00161 (direct_calls.go:3)	CALL	runtime.morestack_noctxt(SB)
//0x00a6 00166 (direct_calls.go:3)	JMP	0

//0x0000 00000 (direct_calls.go:3)	TEXT	"".main(SB), NOSPLIT, $0-0
//0x0000 00000 (direct_calls.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x0000 00000 (direct_calls.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x0000 00000 (direct_calls.go:3)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x0000 00000 (<unknown line number>)	RET
//0x0000 c3
