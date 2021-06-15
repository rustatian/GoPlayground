package main

func main() {
	v := FuncAdd(1, 2, 3)
        _ = v
}

//go:noinline
func FuncAdd(x, y, z int) int { 
        res := x + y - z 
        return res
}

/*
"".main STEXT size=77 args=0x0 locals=0x28 funcid=0x0
        0x0000 00000 (main.go:3)        TEXT    "".main(SB), ABIInternal, $40-0
        0x0000 00000 (main.go:3)        MOVQ    (TLS), CX
        0x0009 00009 (main.go:3)        CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:3)        PCDATA  $0, $-2
        0x000d 00013 (main.go:3)        JLS     70
        0x000f 00015 (main.go:3)        PCDATA  $0, $-1
        0x000f 00015 (main.go:3)        SUBQ    $40, SP
        0x0013 00019 (main.go:3)        MOVQ    BP, 32(SP)
        0x0018 00024 (main.go:3)        LEAQ    32(SP), BP
        0x001d 00029 (main.go:4)        MOVQ    $1, (SP)
        0x0025 00037 (main.go:4)        MOVQ    $2, 8(SP)
        0x002e 00046 (main.go:4)        MOVQ    $3, 16(SP)
        0x0037 00055 (main.go:4)        PCDATA  $1, $0
        0x0037 00055 (main.go:4)        CALL    "".FuncAdd(SB)
        0x003c 00060 (main.go:5)        MOVQ    32(SP), BP
        0x0041 00065 (main.go:5)        ADDQ    $40, SP
        0x0045 00069 (main.go:5)        RET
        0x0046 00070 (main.go:5)        NOP
        0x0046 00070 (main.go:3)        PCDATA  $1, $-1
        0x0046 00070 (main.go:3)        PCDATA  $0, $-2
        0x0046 00070 (main.go:3)        CALL    runtime.morestack_noctxt(SB)
        0x004b 00075 (main.go:3)        PCDATA  $0, $-1
        0x004b 00075 (main.go:3)        JMP     0
"".FuncAdd STEXT nosplit size=27 args=0x20 locals=0x0 funcid=0x0
        0x0000 00000 (main.go:5)        TEXT    "".FuncAdd(SB), NOSPLIT|ABIInternal, $0-32
        0x0000 00000 (main.go:5)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:5)        FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:5)        MOVQ    "".y+16(SP), AX
        0x0005 00005 (main.go:5)        MOVQ    "".x+8(SP), CX
        0x000a 00010 (main.go:5)        ADDQ    CX, AX
        0x000d 00013 (main.go:5)        MOVQ    "".z+24(SP), CX
        0x0012 00018 (main.go:5)        SUBQ    CX, AX
        нu
        0x001a 00026 (main.go:5)        RET




"".main STEXT size=57 args=0x0 locals=0x20 funcid=0x0
        0x0000 00000 (main.go:3)        TEXT    "".main(SB), ABIInternal, $32-0
        0x0000 00000 (main.go:3)        CMPQ    SP, 16(R14)
        0x0004 00004 (main.go:3)        PCDATA  $0, $-2
        0x0004 00004 (main.go:3)        JLS     50
        0x0006 00006 (main.go:3)        PCDATA  $0, $-1
        0x0006 00006 (main.go:3)        SUBQ    $32, SP
        0x000a 00010 (main.go:3)        MOVQ    BP, 24(SP)
        0x000f 00015 (main.go:3)        LEAQ    24(SP), BP
        0x0014 00020 (main.go:3)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (main.go:3)        FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (main.go:4)        MOVL    $1, AX
        0x0019 00025 (main.go:4)        MOVL    $2, BX
        0x001e 00030 (main.go:4)        MOVL    $3, CX
        0x0023 00035 (main.go:4)        PCDATA  $1, $0
        0x0023 00035 (main.go:4)        CALL    "".FuncAdd(SB)
        0x0028 00040 (main.go:5)        MOVQ    24(SP), BP
        0x002d 00045 (main.go:5)        ADDQ    $32, SP
        0x0031 00049 (main.go:5)        RET
        0x0032 00050 (main.go:5)        NOP
        0x0032 00050 (main.go:3)        PCDATA  $1, $-1
        0x0032 00050 (main.go:3)        PCDATA  $0, $-2
        0x0032 00050 (main.go:3)        CALL    runtime.morestack_noctxt(SB)
        0x0037 00055 (main.go:3)        PCDATA  $0, $-1
        0x0037 00055 (main.go:3)        JMP     0
"".FuncAdd STEXT nosplit size=7 args=0x18 locals=0x0 funcid=0x0
        0x0000 00000 (main.go:5)        TEXT    "".FuncAdd(SB), NOSPLIT|ABIInternal, $0-24
        0x0000 00000 (main.go:5)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:5)        FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:5)        FUNCDATA        $5, "".FuncAdd.arginfo1(SB)
        0x0000 00000 (main.go:5)        ADDQ    BX, AX
        0x0003 00003 (main.go:5)        SUBQ    CX, AX
        0x0006 00006 (main.go:5)        RET



*/
