package main

func main() {
	x := mmap()
	_ = x
	x1 := nnew()
	_ = x1
}

//go:noinline
func mmap() []int {
	x := make([]int, 50, 100)
	return x
}

//go:noinline
func nnew() []int {
	x := new([100]int)[0:50]
	return x
}

//0x0000 00000 (main.go:11)	TEXT	"".mmap(SB), $56-24
//0x0000 00000 (main.go:11)	MOVQ	(TLS), CX
//0x0009 00009 (main.go:11)	CMPQ	SP, 16(CX)
//0x000d 00013 (main.go:11)	JLS	103
//0x000f 00015 (main.go:11)	SUBQ	$56, SP
//0x0013 00019 (main.go:11)	MOVQ	BP, 48(SP)
//0x0018 00024 (main.go:11)	LEAQ	48(SP), BP
//0x001d 00029 (main.go:12)	LEAQ	type.int(SB), AX
//0x0024 00036 (main.go:12)	MOVQ	AX, (SP)
//0x0028 00040 (main.go:12)	MOVQ	$50, 8(SP)
//0x0031 00049 (main.go:12)	MOVQ	$100, 16(SP)
//0x003a 00058 (main.go:12)	CALL	runtime.makeslice(SB)
//0x003f 00063 (main.go:12)	MOVQ	24(SP), AX
//0x0044 00068 (main.go:12)	MOVQ	32(SP), CX
//0x0049 00073 (main.go:12)	MOVQ	40(SP), DX
//0x004e 00078 (main.go:13)	MOVQ	AX, "".~r0+64(SP)
//0x0053 00083 (main.go:13)	MOVQ	CX, "".~r0+72(SP)
//0x0058 00088 (main.go:13)	MOVQ	DX, "".~r0+80(SP)
//0x005d 00093 (main.go:13)	MOVQ	48(SP), BP
//0x0062 00098 (main.go:13)	ADDQ	$56, SP
//0x0066 00102 (main.go:13)	RET
//0x0067 00103 (main.go:13)	NOP
//0x0067 00103 (main.go:11)	CALL	runtime.morestack_noctxt(SB)
//0x006c 00108 (main.go:11)	JMP	0


//0x0000 00000 (main.go:17)	TEXT	"".nnew(SB), $24-24
//0x0000 00000 (main.go:17)	MOVQ	(TLS), CX
//0x0009 00009 (main.go:17)	CMPQ	SP, 16(CX)
//0x000d 00013 (main.go:17)	JLS	83
//0x000f 00015 (main.go:17)	SUBQ	$24, SP
//0x0013 00019 (main.go:17)	MOVQ	BP, 16(SP)
//0x0018 00024 (main.go:17)	LEAQ	16(SP), BP
//0x001d 00029 (main.go:18)	LEAQ	type.[100]int(SB), AX
//0x0024 00036 (main.go:18)	MOVQ	AX, (SP)
//0x0028 00040 (main.go:18)	CALL	runtime.newobject(SB)
//0x002d 00045 (main.go:18)	MOVQ	8(SP), AX
//0x0032 00050 (main.go:19)	MOVQ	AX, "".~r0+32(SP)
//0x0037 00055 (main.go:19)	MOVQ	$50, "".~r0+40(SP)
//0x0040 00064 (main.go:19)	MOVQ	$100, "".~r0+48(SP)
//0x0049 00073 (main.go:19)	MOVQ	16(SP), BP
//0x004e 00078 (main.go:19)	ADDQ	$24, SP
//0x0052 00082 (main.go:19)	RET
//0x0053 00083 (main.go:19)	NOP
//0x0053 00083 (main.go:17)	CALL	runtime.morestack_noctxt(SB)
//0x0058 00088 (main.go:17)	JMP	0