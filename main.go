package main

import (
	"github.com/robotn/gohook"
)

func main() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		switch ev.Kind {
		case hook.KeyUp:
			println(ev.Keycode, ev.Rawcode)
		}
	}
}
