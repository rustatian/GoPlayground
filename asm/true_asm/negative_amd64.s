#include "textflag.h"
#include "funcdata.h"

TEXT ·Negative(SB), NOSPLIT, $0
    MOVQ 	x+0(FP), AX
    CMPQ    0x10(CX), SP
    NEGQ 	AX
    MOVQ 	AX, ret+8(FP)
    RET
