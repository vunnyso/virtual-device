package keyboard

import "github.com/vunnyso/virtual-device/linux"

type Repeat struct {
	delay  int32
	period int32
}

type KeyMap map[rune]struct {
	keyCode       linux.Key
	shiftRequired bool
	altGrRequired bool
}
