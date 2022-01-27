package eng

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	l = language.AmericanEnglish
	p = message.NewPrinter(l)
)

type UsEng byte

func (e UsEng) String() string {
	return e.Sprintf("American English (%v)", l)
}

func (e UsEng) Printf(format string, values ...interface{}) {
	p.Printf(format, values...)
}

func (e UsEng) Sprintf(format string, values ...interface{}) string {
	return p.Sprintf(format, values...)
}

func (e UsEng) Printlnf(format string, values ...interface{}) {
	e.Printf(format, values...)
	fmt.Println()
}
