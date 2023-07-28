package main

// #include "wrapper.h"
// #cgo LDFLAGS: -L/opt/homebrew/lib/gcc/13
import "C"

func main() {
	// run the C handler
	inputCStr := C.CString("hello world")
	C.RunModel(inputCStr)
}
