package main

func main() {
	pack("", 1, nil)
}

//go:noinline
func pack(m string, s uint64, buf []byte) {
	copy(buf[0:], m)
	copy(buf[len(m):], []byte{
		byte(s),
		byte(s >> 8),
		byte(s >> 16),
		byte(s >> 24),
		byte(s >> 32),
		byte(s >> 40),
		byte(s >> 48),
		byte(s >> 56),
	})
}
//placeholder.go:8	0x452610		64488b0c25f8ffffff	MOVQ FS:0xfffffff8, CX
//placeholder.go:8	0x452619		483b6110		CMPQ 0x10(CX), SP
//placeholder.go:8	0x45261d		0f86cb000000		JBE 0x4526ee
//placeholder.go:8	0x452623		4883ec28		SUBQ $0x28, SP
//placeholder.go:8	0x452627		48896c2420		MOVQ BP, 0x20(SP)
//placeholder.go:8	0x45262c		488d6c2420		LEAQ 0x20(SP), BP
//placeholder.go:9	0x452631		488b4c2450		MOVQ 0x50(SP), CX
//placeholder.go:9	0x452636		488b442438		MOVQ 0x38(SP), AX
//placeholder.go:9	0x45263b		4839c1			CMPQ AX, CX
//placeholder.go:9	0x45263e		4889ca			MOVQ CX, DX
//placeholder.go:9	0x452641		480f4fc8		CMOVG AX, CX
//placeholder.go:9	0x452645		488b5c2448		MOVQ 0x48(SP), BX
//placeholder.go:9	0x45264a		488b742430		MOVQ 0x30(SP), SI
//placeholder.go:9	0x45264f		4839f3			CMPQ SI, BX
//placeholder.go:9	0x452652		756a			JNE 0x4526be
//placeholder.go:10	0x452654		4839d0			CMPQ DX, AX
//placeholder.go:10	0x452657		0f8788000000		JA 0x4526e5
//placeholder.go:10	0x45265d		48c744241800000000	MOVQ $0x0, 0x18(SP)
//placeholder.go:18	0x452666		488b4c2440		MOVQ 0x40(SP), CX
//placeholder.go:18	0x45266b		48894c2418		MOVQ CX, 0x18(SP)
//placeholder.go:10	0x452670		4829c2			SUBQ AX, DX
//placeholder.go:10	0x452673		4883fa08		CMPQ $0x8, DX
//placeholder.go:10	0x452677		b908000000		MOVL $0x8, CX
//placeholder.go:10	0x45267c		480f4fd1		CMOVG CX, DX
//placeholder.go:10	0x452680		488b4c2458		MOVQ 0x58(SP), CX
//placeholder.go:10	0x452685		4829c1			SUBQ AX, CX
//placeholder.go:10	0x452688		48f7d9			NEGQ CX
//placeholder.go:10	0x45268b		48c1f93f		SARQ $0x3f, CX
//placeholder.go:10	0x45268f		4821c8			ANDQ CX, AX
//placeholder.go:10	0x452692		4801d8			ADDQ BX, AX
//placeholder.go:10	0x452695		488d4c2418		LEAQ 0x18(SP), CX
//placeholder.go:10	0x45269a		4839c8			CMPQ CX, AX
//placeholder.go:10	0x45269d		750a			JNE 0x4526a9
//placeholder.go:10	0x45269f		488b6c2420		MOVQ 0x20(SP), BP
//placeholder.go:10	0x4526a4		4883c428		ADDQ $0x28, SP
//placeholder.go:10	0x4526a8		c3			RET
//placeholder.go:10	0x4526a9		48890424		MOVQ AX, 0(SP)
//placeholder.go:10	0x4526ad		48894c2408		MOVQ CX, 0x8(SP)
//placeholder.go:10	0x4526b2		4889542410		MOVQ DX, 0x10(SP)
//placeholder.go:10	0x4526b7		e8b4a9ffff		CALL runtime.memmove(SB)
//placeholder.go:10	0x4526bc		ebe1			JMP 0x45269f
//placeholder.go:9	0x4526be		48891c24		MOVQ BX, 0(SP)
//placeholder.go:9	0x4526c2		4889742408		MOVQ SI, 0x8(SP)
//placeholder.go:9	0x4526c7		48894c2410		MOVQ CX, 0x10(SP)
//placeholder.go:9	0x4526cc		e89fa9ffff		CALL runtime.memmove(SB)
//placeholder.go:10	0x4526d1		488b442438		MOVQ 0x38(SP), AX
//placeholder.go:10	0x4526d6		488b542450		MOVQ 0x50(SP), DX
//placeholder.go:10	0x4526db		488b5c2448		MOVQ 0x48(SP), BX
//placeholder.go:9	0x4526e0		e96fffffff		JMP 0x452654
//placeholder.go:10	0x4526e5		4889d1			MOVQ DX, CX
//placeholder.go:10	0x4526e8		e8e3a1ffff		CALL runtime.panicSliceB(SB)
//placeholder.go:10	0x4526ed		90			NOPL
//placeholder.go:8	0x4526ee		e80d7affff		CALL runtime.morestack_noctxt(SB)
//placeholder.go:8	0x4526f3		e918ffffff		JMP main.pack(SB)
//
//placeholder.go:8	0x452610		64488b0c25f8ffffff	MOVQ FS:0xfffffff8, CX
//placeholder.go:8	0x452619		483b6110		CMPQ 0x10(CX), SP
//placeholder.go:8	0x45261d		0f86cb000000		JBE 0x4526ee
//placeholder.go:8	0x452623		4883ec28		SUBQ $0x28, SP
//placeholder.go:8	0x452627		48896c2420		MOVQ BP, 0x20(SP)
//placeholder.go:8	0x45262c		488d6c2420		LEAQ 0x20(SP), BP
//placeholder.go:9	0x452631		488b4c2450		MOVQ 0x50(SP), CX
//placeholder.go:9	0x452636		488b442438		MOVQ 0x38(SP), AX
//placeholder.go:9	0x45263b		4839c1			CMPQ AX, CX
//placeholder.go:9	0x45263e		4889ca			MOVQ CX, DX
//placeholder.go:9	0x452641		480f4fc8		CMOVG AX, CX
//placeholder.go:9	0x452645		488b5c2448		MOVQ 0x48(SP), BX
//placeholder.go:9	0x45264a		488b742430		MOVQ 0x30(SP), SI
//placeholder.go:9	0x45264f		4839f3			CMPQ SI, BX
//placeholder.go:9	0x452652		756a			JNE 0x4526be
//placeholder.go:10	0x452654		4839d0			CMPQ DX, AX
//placeholder.go:10	0x452657		0f8788000000		JA 0x4526e5
//placeholder.go:10	0x45265d		48c744241800000000	MOVQ $0x0, 0x18(SP)
//placeholder.go:18	0x452666		488b4c2440		MOVQ 0x40(SP), CX
//placeholder.go:18	0x45266b		48894c2418		MOVQ CX, 0x18(SP)
//placeholder.go:10	0x452670		4829c2			SUBQ AX, DX
//placeholder.go:10	0x452673		4883fa08		CMPQ $0x8, DX
//placeholder.go:10	0x452677		b908000000		MOVL $0x8, CX
//placeholder.go:10	0x45267c		480f4fd1		CMOVG CX, DX
//placeholder.go:10	0x452680		488b4c2458		MOVQ 0x58(SP), CX
//placeholder.go:10	0x452685		4829c1			SUBQ AX, CX
//placeholder.go:10	0x452688		48f7d9			NEGQ CX
//placeholder.go:10	0x45268b		48c1f93f		SARQ $0x3f, CX
//placeholder.go:10	0x45268f		4821c8			ANDQ CX, AX
//placeholder.go:10	0x452692		4801d8			ADDQ BX, AX
//placeholder.go:10	0x452695		488d4c2418		LEAQ 0x18(SP), CX
//placeholder.go:10	0x45269a		4839c8			CMPQ CX, AX
//placeholder.go:10	0x45269d		750a			JNE 0x4526a9
//placeholder.go:10	0x45269f		488b6c2420		MOVQ 0x20(SP), BP
//placeholder.go:10	0x4526a4		4883c428		ADDQ $0x28, SP
//placeholder.go:10	0x4526a8		c3			RET
//placeholder.go:10	0x4526a9		48890424		MOVQ AX, 0(SP)
//placeholder.go:10	0x4526ad		48894c2408		MOVQ CX, 0x8(SP)
//placeholder.go:10	0x4526b2		4889542410		MOVQ DX, 0x10(SP)
//placeholder.go:10	0x4526b7		e8b4a9ffff		CALL runtime.memmove(SB)
//placeholder.go:10	0x4526bc		ebe1			JMP 0x45269f
//placeholder.go:9	0x4526be		48891c24		MOVQ BX, 0(SP)
//placeholder.go:9	0x4526c2		4889742408		MOVQ SI, 0x8(SP)
//placeholder.go:9	0x4526c7		48894c2410		MOVQ CX, 0x10(SP)
//placeholder.go:9	0x4526cc		e89fa9ffff		CALL runtime.memmove(SB)
//placeholder.go:10	0x4526d1		488b442438		MOVQ 0x38(SP), AX
//placeholder.go:10	0x4526d6		488b542450		MOVQ 0x50(SP), DX
//placeholder.go:10	0x4526db		488b5c2448		MOVQ 0x48(SP), BX
//placeholder.go:9	0x4526e0		e96fffffff		JMP 0x452654
//placeholder.go:10	0x4526e5		4889d1			MOVQ DX, CX
//placeholder.go:10	0x4526e8		e8e3a1ffff		CALL runtime.panicSliceB(SB)
//placeholder.go:10	0x4526ed		90			NOPL
//placeholder.go:8	0x4526ee		e80d7affff		CALL runtime.morestack_noctxt(SB)
//placeholder.go:8	0x4526f3		e918ffffff		JMP main.pack(SB)
