package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
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
			info:   "Closes the game",
		},
		"iniciar": {
			action: start,
			info:   "Starts the game",
		},
	}
	lFmt lang.DiceLanguage
)

func main() {
	lFmt = lang.GetCurrent()

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose log")
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debug(lFmt)
	logrus.SetReportCaller(true)

	lFmt.Printlnf("** Tip: you can type \"help\" to show the available options.")
	for !exit {
		cmd := utils.AskString("Choose an option: ", "help")
		callCmd(cmd)
	}
	logrus.Info("Game over!")
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
		lFmt.Printf("• %s: %s\n", key, cmd.info)
	}
	fmt.Println()
}

func getPlayers() []string {
	playersNum := utils.AskInt("How many players? ", 1)
	players := make([]string, playersNum)
	for i := 0; i < playersNum; i++ {
		players[i] = utils.AskRequiredStringf("Name of player #%d: ", i+1)
	}

	return players
}

func start() {
	players := getPlayers()
	game := dice.NewGame(players...)
	game.Start()
}
