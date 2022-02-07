package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
	"github.com/tecnologer/dicegame/server/models/gproto"
	dice "github.com/tecnologer/dicegame/src"
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/utils"
	"google.golang.org/grpc"
)

type cmdType struct {
	action cmdAction
	info   string
}

type cmdAction func()

var (
	exit   bool
	lFmt   lang.DiceLanguage
	cmds   map[string]cmdType
	port   int
	client gproto.GameClient
)

func main() {
	flag.IntVar(&port, "port", 8088, "Port of the server")
	flag.Parse()

	host := fmt.Sprintf(":%d", port)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	connection, err := grpc.Dial(host, opts...)
	if err != nil {
		logrus.Fatalf("connecting to %s. Error: %v", host, err)
	}
	client = gproto.NewGameClient(connection)
	logrus.Infof("connected to %s", host)

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
		lFmt.Sprintf("new-mp"): {
			action: newMpHost,
			info:   lFmt.Sprintf("Creates a new multiplayer game"),
		},
		lFmt.Sprintf("join-mp"): {
			action: joinMpHost,
			info:   lFmt.Sprintf("Joins to existing multiplayer game"),
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

func newMpHost() {
	req := &gproto.NewGameRequest{
		Player:       askPlayer(),
		Password:     utils.AskString("Password (Enter, to be empty):", ""),
		LimitPlayers: int32(utils.AskInt("Limit of players (Enter, to not limit): ", 0)),
	}

	res, err := client.NewGame(context.Background(), req)
	if err != nil {
		lFmt.Printlnf("Error on create new multiplayer host: %v", err)
		return
	}

	lFmt.Printlnf("The multiplayer code is: %s.", res.Code)
}

func joinMpHost() {
	req := &gproto.JoinRequest{
		Player:   askPlayer(),
		Code:     utils.AskRequiredString("Multiplayer Code: "),
		Password: utils.AskString("Password (Enter, to be empty):", ""),
	}

	res, err := client.Join(context.Background(), req)
	if err != nil {
		lFmt.Printlnf("Cannot join to code %s, try again: Error: %v", req.Code, err)
		return
	}

	lFmt.Printlnf("You joinen to %s as %s", req.Code, res.Player.Name)
}

func askPlayer() *gproto.Player {
	return &gproto.Player{
		Name: utils.AskRequiredString("Enter your name: "),
	}
}

func printRules() {
	lFmt.Printf(constants.Rules)
}

func start() {
	players := getPlayers()
	game := dice.NewGame(players...)
	game.Start()
}
