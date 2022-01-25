package models

type DiceLevel byte

const (
	LevelEasy DiceLevel = iota
	LevelMedium
	LevelHard
)

var (
	levelToString = map[DiceLevel]string{
		LevelEasy:   "easy",
		LevelMedium: "medium",
		LevelHard:   "hard",
	}
)

func (dl DiceLevel) String() string {
	return levelToString[dl]
}
