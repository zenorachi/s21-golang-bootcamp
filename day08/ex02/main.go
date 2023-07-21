package main

import (
	"day08/ex02/window"
	"github.com/go-vgo/robotgo"
)

// #include <stdlib.h>
import "C"

func main() {
	window.InitApplication()

	title := "School 21"

	screenWidth, screenHeight := robotgo.GetScreenSize()

	wPtr := window.Create(screenWidth/2, screenHeight/2, 300, 200, title)
	defer C.free(wPtr)
	window.MakeKeyAndOrderFront(wPtr)

	window.RunApplication()
}
