package models

type DiceSide byte

const (
	Brain DiceSide = iota
	Shotgun
	Footprints
)

var (
	sideToString = map[DiceSide]string{
		Brain:      "brain",
		Shotgun:    "shotgun",
		Footprints: "footprints",
	}
)

func (ds DiceSide) String() string {
	return sideToString[ds]
}
