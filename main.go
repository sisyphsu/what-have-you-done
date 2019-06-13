package main

import (
	"github.com/robotn/gohook"
	"github.com/sisyphsu/what-have-you-done/control"
)

func main() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		switch ev.Kind {
		case hook.KeyUp:
			println(ev.Keycode, hook.RawcodetoKeychar(ev.Rawcode), control.KeyCodeMap[ev.Keycode])
		}
	}
}
