package control

type Root struct {
	Keyboard Keyboard
}

type Keyboard struct {
	KeyMap     map[int]int
	keyStatMap map[uint8]*KeyStat
}
