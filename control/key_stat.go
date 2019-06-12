package control

import "time"

type KeyItem struct {
	code      uint16
	minuteMap map[int]int
}

func (s *KeyItem) record() {
	minute := int(time.Now().Unix() / 60)
	if old, ok := s.minuteMap[minute]; ok {
		s.minuteMap[minute] = old + 1
	} else {
		s.minuteMap[minute] = 1
	}
}

func (s *KeyItem) culCount(minute int) (count int) {
	to := int(time.Now().Unix() / 60)
	from := to - minute
	for i := from; i <= to; i++ {
		count += s.minuteMap[i]
	}
	return count
}

type KeyStat struct {
	Minutes int
	keyMap  map[uint16]int
	statMap map[uint16]*KeyItem
}

func NewKeyStat() *KeyStat {
	return &KeyStat{
		Minutes: 60,
		keyMap:  make(map[uint16]int),
		statMap: make(map[uint16]*KeyItem),
	}
}

func (k *KeyStat) Record(code uint16) {
	stat := k.statMap[code]
	if stat == nil {
		stat = &KeyItem{code: code, minuteMap: make(map[int]int)}
		k.statMap[code] = stat
	}
	stat.record()

	k.keyMap[code] = stat.culCount(k.Minutes)
}

func (k *KeyStat) GetStat(keycode uint16) int {
	return k.keyMap[keycode]
}
