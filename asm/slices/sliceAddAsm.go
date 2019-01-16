package main

import (
	"github.com/mmcloughlin/avo/attr"
	"github.com/mmcloughlin/avo/build"
)

func sliceAddAsm() {
	//	TEXT("Add", NOSPLIT, "func(x, y uint64) uint64")
	//	Doc("Add adds x and y.")
	//	x := Load(Param("x"), GP64())
	//	y := Load(Param("y"), GP64())
	//	ADDQ(x, y)
	//	Store(y, ReturnIndex(0))
	//	RET()
	//	Generate()

	build.TEXT("Add", 0, "fucn(a []int, b int) []int")
	a := build.Load(build.Param("a"), build.GP64())

}

//0x0000 00000 (sliceAdd.go:3)    TEXT    "".sliceAdd(SB), $72-56
//0x0000 00000 (sliceAdd.go:3)    MOVQ    (TLS), CX
//0x0009 00009 (sliceAdd.go:3)    CMPQ    SP, 16(CX)
//0x000d 00013 (sliceAdd.go:3)    JLS     161
//0x0013 00019 (sliceAdd.go:3)    SUBQ    $72, SP
//0x0017 00023 (sliceAdd.go:3)    MOVQ    BP, 64(SP)
//0x001c 00028 (sliceAdd.go:3)    LEAQ    64(SP), BP
//0x0021 00033 (sliceAdd.go:3)    FUNCDATA        $0, gclocals·729ba994ca0c03ec3d4026aa1f1b1005(SB)
//0x0021 00033 (sliceAdd.go:3)    FUNCDATA        $1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
//0x0021 00033 (sliceAdd.go:3)    FUNCDATA        $3, gclocals·d571c0f6cf0af59df28f76498f639cf2(SB)
//0x0021 00033 (sliceAdd.go:4)    PCDATA  $2, $0
//0x0021 00033 (sliceAdd.go:4)    PCDATA  $0, $0
//0x0021 00033 (sliceAdd.go:4)    MOVQ    "".a+88(SP), AX
//0x0026 00038 (sliceAdd.go:4)    LEAQ    1(AX), CX
//0x002a 00042 (sliceAdd.go:4)    MOVQ    "".a+96(SP), DX
//0x002f 00047 (sliceAdd.go:4)    CMPQ    CX, DX
//0x0032 00050 (sliceAdd.go:4)    JGT     94
//0x0034 00052 (sliceAdd.go:4)    PCDATA  $2, $1
//0x0034 00052 (sliceAdd.go:4)    PCDATA  $0, $1
//0x0034 00052 (sliceAdd.go:4)    MOVQ    "".a+80(SP), BX
//0x0039 00057 (sliceAdd.go:4)    MOVQ    "".b+104(SP), SI
//0x003e 00062 (sliceAdd.go:4)    MOVQ    SI, (BX)(AX*8)
//0x0042 00066 (sliceAdd.go:4)    PCDATA  $2, $0
//0x0042 00066 (sliceAdd.go:4)    PCDATA  $0, $2
//0x0042 00066 (sliceAdd.go:4)    MOVQ    BX, "".~r2+112(SP)
//0x0047 00071 (sliceAdd.go:4)    MOVQ    CX, "".~r2+120(SP)
//0x004c 00076 (sliceAdd.go:4)    MOVQ    DX, "".~r2+128(SP)
//0x0054 00084 (sliceAdd.go:4)    MOVQ    64(SP), BP
//0x0059 00089 (sliceAdd.go:4)    ADDQ    $72, SP
//0x005d 00093 (sliceAdd.go:4)    RET
//0x005e 00094 (sliceAdd.go:4)    PCDATA  $2, $1
//0x005e 00094 (sliceAdd.go:4)    PCDATA  $0, $0
//0x005e 00094 (sliceAdd.go:4)    LEAQ    type.int(SB), BX
//0x0065 00101 (sliceAdd.go:4)    PCDATA  $2, $0
//0x0065 00101 (sliceAdd.go:4)    MOVQ    BX, (SP)
//0x0069 00105 (sliceAdd.go:4)    PCDATA  $2, $1
//0x0069 00105 (sliceAdd.go:4)    MOVQ    "".a+80(SP), BX
//0x006e 00110 (sliceAdd.go:4)    PCDATA  $2, $0
//0x006e 00110 (sliceAdd.go:4)    MOVQ    BX, 8(SP)
//0x0073 00115 (sliceAdd.go:4)    MOVQ    AX, 16(SP)
//0x0078 00120 (sliceAdd.go:4)    MOVQ    DX, 24(SP)
//0x007d 00125 (sliceAdd.go:4)    MOVQ    CX, 32(SP)
//0x0082 00130 (sliceAdd.go:4)    CALL    runtime.growslice(SB)
//0x0087 00135 (sliceAdd.go:4)    PCDATA  $2, $1
//0x0087 00135 (sliceAdd.go:4)    MOVQ    40(SP), BX
//0x008c 00140 (sliceAdd.go:4)    MOVQ    48(SP), AX
//0x0091 00145 (sliceAdd.go:4)    MOVQ    56(SP), DX
//0x0096 00150 (sliceAdd.go:4)    LEAQ    1(AX), CX
//0x009a 00154 (sliceAdd.go:4)    PCDATA  $0, $1
//0x009a 00154 (sliceAdd.go:4)    MOVQ    "".a+88(SP), AX
//0x009f 00159 (sliceAdd.go:4)    JMP     57
//0x00a1 00161 (sliceAdd.go:4)    NOP
//0x00a1 00161 (sliceAdd.go:3)    PCDATA  $0, $-1
//0x00a1 00161 (sliceAdd.go:3)    PCDATA  $2, $-1
//0x00a1 00161 (sliceAdd.go:3)    CALL    runtime.morestack_noctxt(SB)
//0x00a6 00166 (sliceAdd.go:3)    JMP     0
