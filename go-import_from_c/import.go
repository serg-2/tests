package main
/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lfoo
#include "foo.h"
*/
import "C"
import (
	"fmt"
//	"unsafe"
)

func main() {
//    textConst = "Super"
    fmt.Printf("forty-three == %d\n", C.fortythree())
}

//For go install
//#cgo CFLAGS: -g -Wall

//For go dynamic or static
//#cgo CFLAGS: -I .
//#cgo LDFLAGS: -L . -lfoo


