package control

import (
	hook "github.com/robotn/gohook"
	"testing"
)

func TestKeyStat(t *testing.T) {
	EvChan := hook.Start()
	defer hook.End()

	keyStat := NewKeyStat()

	for ev := range EvChan {
		switch ev.Kind {
		case hook.KeyUp:
			keyStat.Record(ev.Keycode)
			println(ev.Keycode, keyStat.keyMap[ev.Keycode])
		}
	}

}
