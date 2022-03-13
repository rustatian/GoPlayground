package main

import (
	"errors"
	"reflect"
	"unsafe"
)

type S1 struct {
}

func (s *S1) Init() error {
	println("hello")
	return nil
}

func funcAddr(fn interface{}) uintptr {
	// emptyInterface is the header for an interface{} value.
	type emptyInterface struct {
		typ   uintptr
		value *uintptr
	}
	e := (*emptyInterface)(unsafe.Pointer(&fn))
	return *e.value
}

func CreateFunc(fType reflect.Type, f func(args []reflect.Value) (results []reflect.Value)) (reflect.Value, error) {
	if fType.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("invalid input")
	}

	var ins, outs *[]reflect.Type

	ins = new([]reflect.Type)
	outs = new([]reflect.Type)

	for i := 0; i < fType.NumIn(); i++ {
		*ins = append(*ins, fType.In(i))
	}

	for i := 0; i < fType.NumOut(); i++ {
		*outs = append(*outs, fType.Out(i))
	}
	var variadic bool
	variadic = fType.IsVariadic()
	return AllocateStackFrame(*ins, *outs, variadic, f), nil
}
func AllocateStackFrame(ins []reflect.Type, outs []reflect.Type, variadic bool, f func(args []reflect.Value) (results []reflect.Value)) reflect.Value {
	var funcType reflect.Type
	funcType = reflect.FuncOf(ins, outs, variadic)
	return reflect.MakeFunc(funcType, f)
}
