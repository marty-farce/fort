package main

// #include "wrapper.h"
// #cgo LDFLAGS: -L/opt/homebrew/lib/gcc/13 -v
import "C"

func main() {
	// run the C handler
	C.CString("hello world")
}
