#include "textflag.h"
#include "funcdata.h"

TEXT ·Negative(SB), NOSPLIT, $0
    MOVQ 	x+0(FP), AX
    NEGQ 	AX
    CALL runtime·memmove(SB)
    MOVQ 	AX, ret+8(FP)
    RET
