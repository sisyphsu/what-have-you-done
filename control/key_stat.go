package control

import "time"

type KeyStat struct {
	code      uint8
	minuteMap map[int]uint64
}

func (s *KeyStat) record() {
	minute := int(time.Now().Unix() / 60)
	if old, ok := s.minuteMap[minute]; ok {
		s.minuteMap[minute] = old + 1
	} else {
		s.minuteMap[minute] = 1
	}
}
