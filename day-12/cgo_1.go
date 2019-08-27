package main
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"


func main() {
	println("hello cgo")
	C.puts(C.CString("Hello, World\n"))
}
