package spa

import (
	"fmt"

	"github.com/tecnologer/dicegame/src/constants"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	p *message.Printer

	msgMap = map[string]string{
		`** Tip: you can type "help" to show the available options.`:              `** Sugerencia: puedes escribir "ayuda" para mostrar las opciones disponibles.`,
		"Do you want to end your turn? (Yes, Default: No): ":                      "¿Quieres terminar tu turno? (Si,Default: No): ",
		"%s: on this turn, you've %d brains and %d shotguns. Total brains: %d.\n": "%s: en este turno llevas %d cerebros y %d disparos. Cerebros totales: %d.\n",
		"The winner is %s\n with %d brains in %d turns\n":                         "El ganador es %s\n con %d cerebros en %d turnos\n",
		"Choose an option: ":                  "Selecciona una opcion: ",
		"Latin American Spanish (%v)":         "Espaniol Latinoamerica (%v)",
		"exit":                                "salir",
		"start":                               "iniciar",
		"rules":                               "reglas",
		"Game over!":                          "Juego finalizado!",
		"help":                                "ayuda",
		"Closes the game":                     "Cierra el juego",
		"Starts the game":                     "Inicia el juego",
		"Displays the rules":                  "Muestra las reglas.",
		"How many players? ":                  "¿Cuántos jugadores? ",
		"Name of player #%d: ":                "Nombre del jugador #%d: ",
		"Turn of player: %s\n":                "Turno del jugador: %s\n",
		"Selected dices:":                     "Dados seleccionados:",
		"easy":                                "facil",
		"medium":                              "medio",
		"hard":                                "dificil",
		"brain":                               "cerebro",
		"shotgun":                             "disparo",
		"footprints":                          "huellas",
		"Dice #%d gets %s\n":                  "Dado #%d obtuvo %s\n",
		"Invalid input, try again.":           "Entrada no valida, intenta de nuevo.",
		"no":                                  "no",
		"yes":                                 "si",
		"Your turn ends with %d shotguns\n":   "Perdiste el turno con %d disparos\n",
		"Press enter to end your turn...":     "Presiona enter para terminar tu turno...",
		"computer":                            "computadora",
		"It's necessary at least one player.": "Se requiere al menos un jugador",
		constants.Rules:                       rules,
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

const rules = `# Zombie Dice game

## Objectivo del juego

Intentar ser el primer jugador en acumular 13 cerebros.

## Como jugar

El primer jugador es el que gano el juego anterior, o el que pueda decir "Cerebrooooos!" con mas sentimiento.

En tu turno, sacude el vaso, toma tres dados del vaso sin mirar, y lanzalos. Cada uno es una victima. El dado rojo es el mas dificil. El verde es el mas facil, y el amarillo es medio dificil.
Los dados tienen 3 simbolos:

### Cerebro

Te comiste el cerebro de la victima. Pon el cerebro a tu izquierda.

### Disparo

La victima se defendio, pon tu dado de disparo a tu derecha.

### Huellas

Tu victima escapo. Manten tu dado de huellas en frente de ti. Si decides tirar de nuevo, debes de tirar de nuevo este dado, junto con nuevos tomados del vaso hasta completar un total de tres.

Si tiras **tres disparos**, tu turno termina. De lo contrario, tu puedes decidir si detenerte y sumar tus cerebros, o continuar tirando. Si decides detenerte, suma 1 punto por cada cerebro que tu tengas, y pon todos los dados de regreso en el vaso. Es el turno del siguiente jugador.

Si decides continuar, deja todas tus huellas en la mesa. A menos que los tres dados sean huellas, toma nuevos dados al azar del vaso hasta completar un total de tres, y tira de nuevp. Siempre que tires, deberas tirar tres dados.

Despues de tomar un nuevo dado, tu no puedes decidir detenerte... debes de tirar. Pon lo cerebros y los disparos como se indica arriba. Si obtienes mas de 3 disparos tu turno termina y no sumas ningun punto.

De lo contrario, puedes detenerte y sumar puntos, o tirar de nuevo.

## Cerebroooos!

Si no hay suficientes dados en el vaso, anota cuantos cerebros llevas y coloca todos los datos en el vaso (manten los disparos en frente de ti). Despues continua.

## Fin del juego

Juega hasta que alguno alcance 13 cerebros. Entonces finaliza la ronda. Quien tenga mas cerebros al final de la ronda es el ganador. Si hay un empate, el leader (solamente) juega una ronda de desempate.
`
