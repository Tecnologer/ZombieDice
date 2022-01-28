package models

import lang "github.com/tecnologer/dicegame/language"

type DiceSide byte

const (
	Brain DiceSide = iota
	Shotgun
	Footprints
)

var (
	sideToString map[DiceSide]string
)

func init() {
	lFmt = lang.GetCurrent()

	sideToString = map[DiceSide]string{
		Brain:      lFmt.Sprintf("brain"),
		Shotgun:    lFmt.Sprintf("shotgun"),
		Footprints: lFmt.Sprintf("footprints"),
	}
}

func (ds DiceSide) String() string {
	return sideToString[ds]
}
