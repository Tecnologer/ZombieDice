package models

import lang "github.com/tecnologer/dicegame/language"

type DiceLevel byte

const (
	LevelEasy DiceLevel = iota
	LevelMedium
	LevelHard
)

var (
	levelToString map[DiceLevel]string
)

func init() {
	lFmt = lang.GetCurrent()
	levelToString = map[DiceLevel]string{
		LevelEasy:   lFmt.Sprintf("easy"),
		LevelMedium: lFmt.Sprintf("medium"),
		LevelHard:   lFmt.Sprintf("hard"),
	}
}

func (dl DiceLevel) String() string {
	return levelToString[dl]
}
