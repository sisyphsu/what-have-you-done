package main

import (
	"github.com/go-eden/goxui"
	"github.com/sisyphsu/what-have-you-done/control"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	goxui.Init()
	root := control.NewRoot()
	goxui.BindObject(root)
	goxui.StartRelative("view/Start.qml")

	root.Close()
}
