package models

import (
	"fmt"

	"github.com/tecnologer/dicegame/src/utils"
)

type Dice struct {
	Level    DiceLevel
	Sides    [6]DiceSide
	InBucket bool
	Picked   bool
}

func NewDice(level DiceLevel) *Dice {
	return &Dice{
		Level:    level,
		Sides:    getDiceSides(level),
		InBucket: true,
		Picked:   false,
	}
}

func (d *Dice) Roll() DiceSide {
	return d.Sides[utils.GetRandInt(len(d.Sides)-1)]
}

func getDiceSides(level DiceLevel) [6]DiceSide {
	var sides [6]DiceSide

	switch level {
	case LevelEasy:
		sides[0] = Brain
		sides[1] = Footprints
		sides[2] = Brain
		sides[3] = Footprints
		sides[4] = Shotgun
		sides[5] = Brain
	case LevelMedium:
		sides[0] = Brain
		sides[1] = Shotgun
		sides[2] = Footprints
		sides[3] = Brain
		sides[4] = Shotgun
		sides[5] = Footprints
	case LevelHard:
		sides[0] = Shotgun
		sides[1] = Footprints
		sides[2] = Shotgun
		sides[3] = Shotgun
		sides[4] = Brain
		sides[5] = Shotgun
	}
	return sides
}

func (d *Dice) String() string {
	return fmt.Sprint(d.Level)
}

func (d *Dice) Println() {
	lFmt.Printlnf("%s", d.Level)
}

func (d *Dice) Print() {
	lFmt.Printf("%s", d.Level)
}
