package main

import (
	"github.com/sisyphsu/goxui"
	"github.com/sisyphsu/what-have-you-done/control"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	root := control.NewRoot()

	goxui.Init()
	goxui.BindObject(root)
	goxui.StartRelative("view/hook.qml")

	root.Close()
}
