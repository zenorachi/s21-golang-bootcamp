package main

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "application.h"
// #include "window.h"
import "C"
import (
	"github.com/go-vgo/robotgo"
	"unsafe"
)

func main() {
	C.InitApplication()

	title := C.CString("School 21")

	screenWidth, screenHeight := robotgo.GetScreenSize()
	cWidth := C.int(screenWidth)
	cHeight := C.int(screenHeight)

	wPtr := C.Window_Create(cWidth/2, cHeight/2, 300, 200, title)
	C.Window_MakeKeyAndOrderFront(wPtr)

	defer C.free(unsafe.Pointer(title))
	C.RunApplication()
}
