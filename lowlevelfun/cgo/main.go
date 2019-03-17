package main

/*
#include "stdio.h"
*/
import "C"

func main() {
	C.printf(C.CString("Hello world\n"))
}
