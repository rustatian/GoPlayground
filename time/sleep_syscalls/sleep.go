package main

import "fmt"

type A struct {
	val int
}

func main() {
	a := A{val: 1}
	surprise(&a)
	fmt.Println(a)
}

//go:noinline
func surprise(a ...*A) {
	a[0].val = 5
}

//"".main STEXT size=203 args=0x0 locals=0x60
//0x0000 00000 (sleep.go:9)       TEXT    "".main(SB), ABIInternal, $96-0
//0x0000 00000 (sleep.go:9)       MOVQ    (TLS), CX
//0x0009 00009 (sleep.go:9)       CMPQ    SP, 16(CX)
//0x000d 00013 (sleep.go:9)       JLS     193
//0x0013 00019 (sleep.go:9)       SUBQ    $96, SP
//0x0017 00023 (sleep.go:9)       MOVQ    BP, 88(SP)
//0x001c 00028 (sleep.go:9)       LEAQ    88(SP), BP
//0x0021 00033 (sleep.go:11)      MOVQ    $0, ""..autotmp_5+64(SP)
//0x002a 00042 (sleep.go:11)      MOVQ    $1, ""..autotmp_5+64(SP)
//0x0033 00051 (sleep.go:11)      LEAQ    ""..autotmp_5+64(SP), AX
//0x0038 00056 (sleep.go:11)      MOVQ    AX, (SP)
//0x003c 00060 (sleep.go:11)      MOVQ    $1, 8(SP)
//0x0045 00069 (sleep.go:11)      MOVQ    $1, 16(SP)
//0x004e 00078 (sleep.go:11)      CALL    "".surprise(SB)
//0x0053 00083 (sleep.go:12)      MOVQ    $1, (SP)
//0x005b 00091 (sleep.go:12)      CALL    runtime.convT64(SB)
//0x0060 00096 (sleep.go:12)      MOVQ    8(SP), AX
//0x0065 00101 (sleep.go:12)      XORPS   X0, X0
//0x0068 00104 (sleep.go:12)      MOVUPS  X0, ""..autotmp_14+72(SP)
//0x006d 00109 (sleep.go:12)      LEAQ    type."".A(SB), CX
//0x0074 00116 (sleep.go:12)      MOVQ    CX, ""..autotmp_14+72(SP)
//0x0079 00121 (sleep.go:12)      MOVQ    AX, ""..autotmp_14+80(SP)
//0x007e 00126 (sleep.go:12)      XCHGL   AX, AX

//"".surprise STEXT nosplit size=50 args=0x18 locals=0x8
//0x0000 00000 (sleep.go:16)      TEXT    "".surprise(SB), NOSPLIT|ABIInternal, $8-24
//0x0000 00000 (sleep.go:16)      SUBQ    $8, SP
//0x0004 00004 (sleep.go:16)      MOVQ    BP, (SP)
//0x0008 00008 (sleep.go:16)      LEAQ    (SP), BP
//0x000c 00012 (sleep.go:17)      MOVQ    "".a+24(SP), AX
//0x0011 00017 (sleep.go:17)      TESTQ   AX, AX
//0x0014 00020 (sleep.go:17)      JLS     43
//0x0016 00022 (sleep.go:17)      MOVQ    "".a+16(SP), AX
//0x001b 00027 (sleep.go:17)      MOVQ    $5, (AX)
//0x0022 00034 (sleep.go:18)      MOVQ    (SP), BP
//0x0026 00038 (sleep.go:18)      ADDQ    $8, SP
//0x002a 00042 (sleep.go:18)      RET
//0x002b 00043 (sleep.go:17)      CALL    runtime.panicindex(SB)
//0x0030 00048 (sleep.go:17)      UNDEF
















//"".main STEXT size=223 args=0x0 locals=0x68
//0x0000 00000 (sleep.go:9)       TEXT    "".main(SB), ABIInternal, $104-0
//0x0000 00000 (sleep.go:9)       MOVQ    (TLS), CX
//0x0009 00009 (sleep.go:9)       CMPQ    SP, 16(CX)
//0x000d 00013 (sleep.go:9)       JLS     213
//0x0013 00019 (sleep.go:9)       SUBQ    $104, SP
//0x0017 00023 (sleep.go:9)       MOVQ    BP, 96(SP)
//0x001c 00028 (sleep.go:9)       LEAQ    96(SP), BP
//0x0021 00033 (sleep.go:10)      MOVQ    $0, "".a+64(SP)
//0x002a 00042 (sleep.go:10)      MOVQ    $1, "".a+64(SP)
//0x0033 00051 (sleep.go:11)      MOVQ    $0, ""..autotmp_5+72(SP)
//0x003c 00060 (sleep.go:11)      LEAQ    "".a+64(SP), AX
//0x0041 00065 (sleep.go:11)      MOVQ    AX, ""..autotmp_5+72(SP)
//0x0046 00070 (sleep.go:11)      LEAQ    ""..autotmp_5+72(SP), AX
//0x004b 00075 (sleep.go:11)      MOVQ    AX, (SP)
//0x004f 00079 (sleep.go:11)      MOVQ    $1, 8(SP)
//0x0058 00088 (sleep.go:11)      MOVQ    $1, 16(SP)
//0x0061 00097 (sleep.go:11)      CALL    "".surprise(SB)
//0x0066 00102 (sleep.go:12)      MOVQ    "".a+64(SP), AX
//0x006b 00107 (sleep.go:12)      MOVQ    AX, (SP)
//0x006f 00111 (sleep.go:12)      CALL    runtime.convT64(SB)
//0x0074 00116 (sleep.go:12)      MOVQ    8(SP), AX
//0x0079 00121 (sleep.go:12)      XORPS   X0, X0
//0x007c 00124 (sleep.go:12)      MOVUPS  X0, ""..autotmp_14+80(SP)
//0x0081 00129 (sleep.go:12)      LEAQ    type."".A(SB), CX
//0x0088 00136 (sleep.go:12)      MOVQ    CX, ""..autotmp_14+80(SP)
//0x008d 00141 (sleep.go:12)      MOVQ    AX, ""..autotmp_14+88(SP)
//0x0092 00146 (sleep.go:12)      XCHGL   AX, AX


//"".surprise STEXT nosplit size=53 args=0x18 locals=0x8
//        0x0000 00000 (sleep.go:16)      TEXT    "".surprise(SB), NOSPLIT|ABIInternal, $8-24
//        0x0000 00000 (sleep.go:16)      SUBQ    $8, SP
//        0x0004 00004 (sleep.go:16)      MOVQ    BP, (SP)
//        0x0008 00008 (sleep.go:16)      LEAQ    (SP), BP
//        0x000c 00012 (sleep.go:17)      MOVQ    "".a+24(SP), AX
//        0x0011 00017 (sleep.go:17)      TESTQ   AX, AX
//        0x0014 00020 (sleep.go:17)      JLS     46
//        0x0016 00022 (sleep.go:17)      MOVQ    "".a+16(SP), AX
//        0x001b 00027 (sleep.go:17)      MOVQ    (AX), AX
//        0x001e 00030 (sleep.go:17)      MOVQ    $5, (AX)
//        0x0025 00037 (sleep.go:18)      MOVQ    (SP), BP
//        0x0029 00041 (sleep.go:18)      ADDQ    $8, SP
//        0x002d 00045 (sleep.go:18)      RET
//        0x002e 00046 (sleep.go:17)      CALL    runtime.panicindex(SB)
//        0x0033 00051 (sleep.go:17)      UNDEF

















