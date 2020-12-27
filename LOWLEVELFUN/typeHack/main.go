package main

import (
	"fmt"
	"unsafe"
)

// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
type tflag uint8

type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype

type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32   // hash of type; avoids computation in hash tables
	tflag      tflag    // extra type information flags
	align      uint8    // alignment of variable with this type
	fieldAlign uint8    // alignment of struct field with this type
	kind       uint8    // enumeration for C
	alg        *typeAlg // algorithm table
	gcdata     *byte    // garbage collection data
	str        nameOff  // string form
	ptrToThis  typeOff  // type for pointer to this type, may be zero
}

type eface struct {
	typ  *rtype
	word unsafe.Pointer
}

func (e eface) String() string {
	return fmt.Sprintf("type: %#v\n\ndataptr: %v", *e.typ, e.word)
}

func getEface(i interface{}) eface {
	return *(*eface)(unsafe.Pointer(&i))
}

func main() {
	var a int
	var b int
	var c string
	var d float32
	var e float64
	var f rtype
	var g eface

	fmt.Printf("a int:\n%s\n\n", getEface(a))
	fmt.Printf("b int:\n%s\n\n", getEface(b))
	fmt.Printf("c string:\n%s\n\n", getEface(c))
	fmt.Printf("d float32:\n%s\n\n", getEface(d))
	fmt.Printf("e float64:\n%s\n\n", getEface(e))
	fmt.Printf("f rtype:\n%s\n\n", getEface(f))
	fmt.Printf("g eface:\n%s\n\n", getEface(g))
}
