package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
	"github.com/tecnologer/dicegame/server/models/gproto"
	dice "github.com/tecnologer/dicegame/src"
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	you    *gproto.Player
)

func main() {
	flag.IntVar(&port, "port", 8088, "Port of the server")
	flag.Parse()

	host := fmt.Sprintf(":%d", port)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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

	newR, ok := res.GetContent().(*gproto.Response_NewGameResponse)
	if !ok {
		lFmt.Printlnf("Invalid response for new game.")
		return
	}

	go registerForNotifications(newR.NewGameResponse.Code)
	lFmt.Printlnf("The multiplayer code is: %s.", newR.NewGameResponse.Code)
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
	newR, ok := res.GetContent().(*gproto.Response_JoinResponse)
	if !ok {
		lFmt.Printlnf("Invalid response for join game.")
		return
	}

	go registerForNotifications(req.Code)
	lFmt.Printlnf("You joinen to %s as %s", req.Code, newR.JoinResponse.Player.Name)
}

func askPlayer() *gproto.Player {
	if you == nil {
		you = &gproto.Player{
			Name: utils.AskRequiredString("Enter your name: "),
		}
	}
	return you
}

func printRules() {
	lFmt.Printf(constants.Rules)
}

func start() {
	players := getPlayers()
	game := dice.NewGame(players...)
	game.Start()
}

func registerForNotifications(code string) {
	req := &gproto.RegisterNotifications{
		Code: code,
	}
	stream, err := client.Notifications(context.Background(), req)
	if err != nil {
		panic(err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			logrus.Error(err)
			os.Exit(0)
		}

		resNotif, ok := response.GetContent().(*gproto.Response_Notification)
		if !ok {
			logrus.Warn("response is not notification")
			continue
		}
		notif := resNotif.Notification

		switch notif.Type {
		case gproto.NotificationResponse_JOIN:
			playerJoin, ok := notif.Content.(*gproto.NotificationResponse_PlayerJoin)
			if !ok {
				continue
			}
			player := playerJoin.PlayerJoin.Player
			lFmt.Printlnf("The player %s joined to the game.", player.Name)
		case gproto.NotificationResponse_LEFT:
			playerLeft, ok := notif.Content.(*gproto.NotificationResponse_PlayerLeft)
			if !ok {
				continue
			}
			player := playerLeft.PlayerLeft.Player
			lFmt.Printlnf("The player %s left the game.", player.Name)
		case gproto.NotificationResponse_OVER:
			gameOver, ok := notif.Content.(*gproto.NotificationResponse_GameOver)
			if !ok {
				continue
			}
			winner := gameOver.GameOver.Winner
			lFmt.Printlnf("The player %s won the game with %d brains.", winner.Name, gameOver.GameOver.Brains)
		case gproto.NotificationResponse_TURN:
			playerTurn, ok := notif.Content.(*gproto.NotificationResponse_PlayerTurn)
			if !ok {
				continue
			}
			player := playerTurn.PlayerTurn.Player
			if player.Name != you.Name {
				lFmt.Printlnf("It's turn of player %s.", player.Name)
			} else {
				lFmt.Printlnf("It's your turn.")
			}
		}
	}
}

// func stopNotifications() {
// 	//TODO
// }
