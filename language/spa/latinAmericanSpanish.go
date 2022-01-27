package spa

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	p *message.Printer

	msgMap = map[string]string{
		`** Tip: you can type "help" to show the available options.`: `** Sugerencia: puedes escribir "help" para mostrar las opciones disponibles.`,
		"Choose an option: ":          "Selecciona una opcion: ",
		"Latin American Spanish (%v)": "Espaniol Latinoamerica (%v)",
	}
)

func init() {

	for key, msg := range msgMap {
		err := message.SetString(language.LatinAmericanSpanish, key, msg)

		if err != nil {
			panic(err)
		}
	}
	p = message.NewPrinter(language.LatinAmericanSpanish)
}

//Es419 is a type for printer for the language Latin American Spanish
type Es419 byte

func (s Es419) String() string {
	return s.Sprintf("Latin American Spanish (%v)", language.LatinAmericanSpanish)
}

func (s Es419) Printf(format string, values ...interface{}) {
	p.Printf(format, values...)
}

func (s Es419) Sprintf(format string, values ...interface{}) string {
	return p.Sprintf(format, values...)
}

func (s Es419) Printlnf(format string, values ...interface{}) {
	s.Printf(format, values...)
	fmt.Println()
}
