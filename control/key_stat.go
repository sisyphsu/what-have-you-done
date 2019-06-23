package control

import (
	"github.com/go-eden/goxui"
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
	Minutes  int
	maxCount int
	keyMap   map[string]int
	statMap  map[string]*KeyItem
}

func NewKeyStat() *KeyStat {
	ks := &KeyStat{
		Minutes:  60,
		maxCount: 20,
		keyMap:   make(map[string]int),
		statMap:  make(map[string]*KeyItem),
	}
	go func() {
		for range time.Tick(time.Second * 20) {
			ks.flush()
		}
	}()
	return ks
}

func (k *KeyStat) record(key string) {
	stat := k.statMap[key]
	if stat == nil {
		stat = &KeyItem{name: key, minuteMap: make(map[int]int)}
		k.statMap[key] = stat
	}
	stat.record()
	// flush count
	newCount := stat.culCount(k.Minutes)
	k.keyMap[key] = newCount
	if k.maxCount < newCount {
		k.maxCount = newCount
		k.flush()
	} else {
		goxui.TriggerEvent(key, float64(newCount)/float64(k.maxCount))
	}
}

func (k *KeyStat) flush() {
	var maxCount int
	for _, count := range k.keyMap {
		if maxCount < count {
			maxCount = count
		}
	}
	for key, count := range k.keyMap {
		goxui.TriggerEvent(key, float64(count)/float64(k.maxCount))
	}
}

func (k *KeyStat) GetScore(key string) float64 {
	count := k.keyMap[key]
	if k.maxCount <= 0 {
		return 0
	}
	return float64(count) / float64(k.maxCount)
}
