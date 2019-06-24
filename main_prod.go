package main

import (
	"github.com/go-eden/goxui"
	slog "github.com/go-eden/slf4go"
	"github.com/sisyphsu/what-have-you-done/control"
	"github.com/sisyphsu/what-have-you-done/view"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	goxui.Init()

	slog.SetLevel(slog.InfoLevel)

	root := control.NewRoot()
	goxui.BindObject(root)
	goxui.AddResourceData("/", view.DATA)
	goxui.MapResource("img", ":/images/")
	goxui.Start("qrc:/Start.qml")

	root.Close()
}
