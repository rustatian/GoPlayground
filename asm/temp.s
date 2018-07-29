//int64
"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:3)	TEXT	"".main(SB), NOSPLIT, $0-0
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<unknown line number>)	RET
	0x0000 c3                                               .
"".add STEXT nosplit size=24 args=0x20 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:7)	TEXT	"".add(SB), NOSPLIT, $0-32
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (direct_topfunc_call.go:7)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (direct_topfunc_call.go:8)	ADDQ	CX, AX
	0x000d 00013 (direct_topfunc_call.go:8)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (direct_topfunc_call.go:8)	MOVB	$1, "".~r3+32(SP)
	0x0017 00023 (direct_topfunc_call.go:8)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c6 44 24 20 01 c3                          $..D$ ..



//int
"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:3)	TEXT	"".main(SB), NOSPLIT, $0-0
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<unknown line number>)	RET
	0x0000 c3                                               .
"".add STEXT nosplit size=24 args=0x20 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:7)	TEXT	"".add(SB), NOSPLIT, $0-32
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$0, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (direct_topfunc_call.go:7)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (direct_topfunc_call.go:8)	ADDQ	CX, AX
	0x000d 00013 (direct_topfunc_call.go:8)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (direct_topfunc_call.go:8)	MOVB	$1, "".~r3+32(SP)
	0x0017 00023 (direct_topfunc_call.go:8)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c6 44 24 20 01 c3                          $..D$ ..

//int32
"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:3)	TEXT	"".main(SB), NOSPLIT, $0-0
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<unknown line number>)	RET
	0x0000 c3                                               .
"".add STEXT nosplit size=20 args=0x10 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:7)	TEXT	"".add(SB), NOSPLIT, $0-16
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:7)	MOVL	"".b+12(SP), AX
	0x0004 00004 (direct_topfunc_call.go:7)	MOVL	"".a+8(SP), CX
	0x0008 00008 (direct_topfunc_call.go:8)	ADDL	CX, AX
	0x000a 00010 (direct_topfunc_call.go:8)	MOVL	AX, "".~r2+16(SP)
	0x000e 00014 (direct_topfunc_call.go:8)	MOVB	$1, "".~r3+20(SP)
	0x0013 00019 (direct_topfunc_call.go:8)	RET
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c6 44  .D$..L$....D$..D
	0x0010 24 14 01 c3                                      $...