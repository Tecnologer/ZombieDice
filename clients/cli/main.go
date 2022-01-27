package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	dice "github.com/tecnologer/dicegame/src"
	"github.com/tecnologer/dicegame/src/utils"
)

type cmdType struct {
	action cmdAction
	info   string
}
type cmdAction func()

var (
	exit bool
	cmds = map[string]cmdType{
		"salir": {
			action: exitCmd,
			info:   "Cierra el juego",
		},
		"iniciar": {
			action: start,
			info:   "Inicia el juego",
		},
	}
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose log")
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetReportCaller(true)

	fmt.Println("Tip: puedes escribir \"help\" para mostrar las opciones disponibles")
	for !exit {
		cmd := utils.AskString("Selecciona una opcion: ", "help")
		callCmd(cmd)
	}
	logrus.Info("Fin del juego!")
}

func exitCmd() {
	exit = true
}

func callCmd(cmdKey string) {
	cmd, ok := cmds[cmdKey]
	if !ok {
		helpCmd()
		return
	}
	cmd.action()
}

func helpCmd() {
	for key, cmd := range cmds {
		fmt.Printf("• %s: %s\n", key, cmd.info)
	}
	fmt.Println()
}

func getPlayers() []string {
	playersNum := utils.AskInt("Cuantos jugadores seran? ", 1)
	players := make([]string, playersNum)
	for i := 0; i < playersNum; i++ {
		players[i] = utils.AskRequiredStringf("Nombre jugador %d: ", i+1)
	}

	return players
}

func start() {
	players := getPlayers()
	game := dice.NewGame(players...)
	game.Start()
}
