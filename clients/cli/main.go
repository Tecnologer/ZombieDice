package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
	dice "github.com/tecnologer/dicegame/src"
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/utils"
)

type cmdType struct {
	action cmdAction
	info   string
}

type cmdAction func()

var (
	exit bool
	lFmt lang.DiceLanguage
	cmds map[string]cmdType
)

func main() {
	lFmt = lang.GetCurrent()
	cmds = map[string]cmdType{
		lFmt.Sprintf("exit"): {
			action: exitCmd,
			info:   lFmt.Sprintf("Closes the game"),
		},
		lFmt.Sprintf("start"): {
			action: start,
			info:   lFmt.Sprintf("Starts the game"),
		},
		lFmt.Sprintf("rules"): {
			action: printRules,
			info:   lFmt.Sprintf("Displays the rules"),
		},
	}

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose log")
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debug(lFmt)
	logrus.SetReportCaller(true)

	lFmt.Printlnf("** Tip: you can type \"help\" to show the available options.", lFmt.Sprintf("help"))
	for !exit {
		cmd := utils.AskString("Choose an option: ", lFmt.Sprintf("help"))
		callCmd(cmd)
	}
	logrus.Info(lFmt.Sprintf("Game over!"))
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

func printRules() {
	lFmt.Printf(constants.Rules)
}

func start() {
	players := getPlayers()
	game := dice.NewGame(players...)
	game.Start()
}
