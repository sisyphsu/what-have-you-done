package control

var KeyStatMap = make(map[uint8]*KeyStat)

type KeyStats struct {
	KeyMap map[int]int
}

func Record(code uint8) {
	stat := KeyStatMap[code]
	if stat == nil {
		stat = &KeyStat{code: code, minuteMap: make(map[int]uint64)}
		KeyStatMap[code] = stat
	}
	stat.record()
}
