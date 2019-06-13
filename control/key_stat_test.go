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
			key := KeyCodeMap[ev.Keycode]
			keyStat.Record(key)
			println(key)
		}
	}

}
