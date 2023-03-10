package main

import (
	"sync"
)

func main() {
	a := make([]byte, 10)
	foo(a)
}

func foo(b []byte) {
	b[1] = 10
}

//go:noinline
func xy(a []string) {
	x := a[0]
	y := a[1]
	_ = x + y
}

//go:noinline
func yx_faster(a []string) {
	if len(a) > 0 {
		x := a[1]
		y := a[0]
		_ = x + y
	}
}

//go:noinline
func sliceUniqueStd(ss []string) []string {
	seen := make(map[string]bool, len(ss))
	i := 0
	for _, v := range ss {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = true
		ss[i] = v

		i++
	}
	return ss[:i]
}

var sourcePool = sync.Pool{
	New: func() interface{} {
		return make(map[string]bool)
	},
}

//go:noinline
func getSource() map[string]bool {
	r := sourcePool.Get().(map[string]bool)
	return r
}

//go:noinline
func putSource(r map[string]bool) {
	sourcePool.Put(r)
}

//go:noinline
func sliceUniqueUpdated(ss []string) []string {
	seen := make(map[string]bool, len(ss))
	//seen := getSource()
	ii := 0

	for i := 0; i < len(ss); i++ {
		if _, ok := seen[ss[i]]; ok {
			continue
		}

		seen[ss[i]] = true
		if ii > 0 && len(ss) > ii {
			ss[ii] = ss[i]
		}

		ii++
	}

	//putSource(seen)

	if len(ss) > ii && ii > 0 {
		return ss[:ii]
	}

	return nil
}

type SomeStruct struct {
	header []byte
}

func (sh *SomeStruct) WritePayloadLenFun(header []byte, v uint32) {
	_ = header[3]
	header[0] = byte(v)
	header[1] = byte(v >> 8)
	header[2] = byte(v >> 16)
	header[3] = byte(v >> 24)
}

func (sh *SomeStruct) WritePayloadLenNotFun(len uint32) {
	sh.header[0] = byte(len)
	sh.header[1] = byte(len >> 8)
	sh.header[2] = byte(len >> 16)
	sh.header[3] = byte(len >> 24)
}

/*
       MOVL    "".len+40(SP), SI
        MOVB    SIB, (BX)


   MOVQ    "".sh+32(SP), DX
   MOVQ    8(DX), CX
   MOVQ    (DX), BX
   TESTQ   CX, CX
   JLS     _SomeStruct_WritePayloadLenNotFun_pc143
   MOVL    "".len+40(SP), SI
   MOVB    SIB, (BX)












        MOVQ    8(DX), CX
        MOVQ    (DX), BX
        CMPQ    CX, $1
        JLS     _SomeStruct_WritePayloadLenFun_pc143
        MOVL    SI, DI
        SHRL    $8, SI
        MOVB    SIB, 1(BX)



        MOVQ    8(DX), CX
        MOVQ    (DX), BX
        CMPQ    CX, $1
        JLS     _SomeStruct_WritePayloadLenNotFun_pc133
        MOVL    SI, DI
        SHRL    $8, SI
        MOVB    SIB, 1(BX)

*/
