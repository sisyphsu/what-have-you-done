package control

import (
	"github.com/sisyphsu/goxui"
	"time"
)

type KeyItem struct {
	name      string
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
	keyMap  map[string]int
	statMap map[string]*KeyItem
}

func NewKeyStat() *KeyStat {
	return &KeyStat{
		Minutes: 60,
		keyMap:  make(map[string]int),
		statMap: make(map[string]*KeyItem),
	}
}

func (k *KeyStat) Record(key string) {
	stat := k.statMap[key]
	if stat == nil {
		stat = &KeyItem{name: key, minuteMap: make(map[int]int)}
		k.statMap[key] = stat
	}
	stat.record()
	// flush count
	newCount := stat.culCount(k.Minutes)
	k.keyMap[key] = newCount
	goxui.TriggerEvent("count_change_"+stat.name, newCount)
}

func (k *KeyStat) GetStat(key string) int {
	return k.keyMap[key]
}
