#include "textflag.h"
#include "funcdata.h"

TEXT ·neg(SB), NOSPLIT, $0
	MOVQ 	x+0(FP), AX
	NEGQ 	AX
	MOVQ 	AX, ret+8(FP)
	RET