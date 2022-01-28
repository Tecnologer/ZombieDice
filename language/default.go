package lang

import (
	"strings"
	"sync"

	"github.com/Xuanwo/go-locale"
	"github.com/tecnologer/dicegame/language/eng"
	"github.com/tecnologer/dicegame/language/spa"
)

var (
	current DiceLanguage
	spanish spa.Es419
	english eng.UsEng

	m *sync.Mutex
)

func init() {
	m = &sync.Mutex{}
}

type DiceLanguage interface {
	String() string
	Printf(string, ...interface{})
	Printlnf(string, ...interface{})
	Sprintf(string, ...interface{}) string
}

func getDefault() DiceLanguage {
	s, err := locale.Detect()
	if err == nil && strings.HasPrefix(s.String(), "es") {
		return spanish
	}
	return english
}

func GetCurrent() DiceLanguage {
	//TODO: check this
	m.Lock()
	defer m.Unlock()

	if current == nil {
		return getDefault()
	}
	return current
}

func Load() {
	panic("ToDo")
}
