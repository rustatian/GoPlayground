package main

import (
	"time"
)

func main() {
	time.Sleep(time.Microsecond * 1)
}

// called from assembly
//func badmcall(fn func(*g)) {
//	throw("runtime: mcall called on m->g0 stack")
//}

//func badmcall2(fn func(*g)) {
//	throw("runtime: mcall function returned")
//}

// func mcall(fn func(*g))
// Switch to m->g0's stack, call fn(g).
// Fn must never return. It should gogo(&g->sched)
// to keep running g.
//TEXT runtime·mcall(SB), NOSPLIT, $0-8
//MOVQ	fn+0(FP), DI
//
//get_tls(CX)
//MOVQ	g(CX), AX	// save state in g->sched
//MOVQ	0(SP), BX	// caller's PC
//MOVQ	BX, (g_sched+gobuf_pc)(AX)
//LEAQ	fn+0(FP), BX	// caller's SP
//MOVQ	BX, (g_sched+gobuf_sp)(AX)
//MOVQ	AX, (g_sched+gobuf_g)(AX)
//MOVQ	BP, (g_sched+gobuf_bp)(AX)
//
//// switch to m->g0 & its stack, call fn
//MOVQ	g(CX), BX
//MOVQ	g_m(BX), BX
//MOVQ	m_g0(BX), SI
//CMPQ	SI, AX	// if g == m->g0 call badmcall
//JNE	3(PC)
//MOVQ	$runtime·badmcall(SB), AX
//JMP	AX
//MOVQ	SI, g(CX)	// g = m->g0
//MOVQ	(g_sched+gobuf_sp)(SI), SP	// sp = m->g0->sched.sp
//PUSHQ	AX
//MOVQ	DI, DX
//MOVQ	0(DI), DI
//CALL	DI
//POPQ	AX
//MOVQ	$runtime·badmcall2(SB), AX
//JMP	AX
//RET

//g0 is a special system goroutine.
//Each M (thread) has own g0.
//g0 executes scheduler code. I.e. when a goroutine finishes running,
//execution is switched to g0, g0 chooses next goroutine to run, and
//then execution switches to the new goroutine.
//g0 has relatively large stack (32K) and is intended to never grow it.
//g0 is also used for split stack management and GC.
//The other special goroutine is gsignal, it's used for signal handling.
