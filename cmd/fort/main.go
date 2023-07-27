package main

// #include "wrapper.h"
// #cgo LDFLAGS: -L/opt/homebrew/opt/gfortran/lib/gcc -v
import "C"

func main() {
	// run the C handler
	C.CString("hello world")
}
