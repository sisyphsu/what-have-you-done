package control

import hook "github.com/robotn/gohook"

type Root struct {
	KeyStat *KeyStat
}

func NewRoot() (r *Root) {
	r = &Root{
		KeyStat: NewKeyStat(),
	}
	go r.deamonHook()
	return
}

// Start global hook
func (r *Root) deamonHook() {
	EvChan := hook.Start()
	for ev := range EvChan {
		switch ev.Kind {
		case hook.KeyUp:
		}
	}
}

// Close release global hook
func (r *Root) Close() {
	hook.StopEvent()
}
