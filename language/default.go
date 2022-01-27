package lang

import "github.com/tecnologer/dicegame/language/spa"

var current DiceLanguage

type DiceLanguage interface {
	String() string
	Printf(string, ...interface{})
	Printlnf(string, ...interface{})
	Sprintf(string, ...interface{}) string
}

func getDefault() DiceLanguage {
	// var english eng.UsEng
	var english spa.Es419
	current = english
	return english
}

func GetCurrent() DiceLanguage {
	if current == nil {
		return getDefault()
	}
	return current
}

func Load() {
	panic("ToDo")
}
