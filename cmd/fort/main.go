package main

// #include "wrapper.h"
// #cgo LDFLAGS: -v
import "C"

func main() {
	// run the C handler
	C.CString("hello world")
}
