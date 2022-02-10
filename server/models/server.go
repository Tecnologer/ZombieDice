package models

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/dicegame/server/models/gproto"
	"google.golang.org/grpc"
)

var (
	notifications = make(chan *notification, 5)
)

type DiceServer struct {
	gproto.UnimplementedGameServer
	host string
}

func (s DiceServer) NewGame(ctx context.Context, req *gproto.NewGameRequest) (*gproto.Response, error) {
	code := CreateGame(req.Password, int(req.LimitPlayers))
	res := newGameResponse(code)

	msg := "encrypted"
	if req.Password == "" {
		msg = "uncrypted"
	}
	logrus.Infof("New %s game created with code %s", msg, code)
	JoinToGame(code, req.Password, req.Player)
	return res, nil
}

func (s DiceServer) Join(ctx context.Context, req *gproto.JoinRequest) (*gproto.Response, error) {
	err := JoinToGame(req.Code, req.Password, req.Player)
	if err != nil {
		return nil, err
	}

	res := newJoinResponse(req.Player, true)
	// actualGames.sendNotification(notifyNewJoin(req.Code, req.Player))
	notifications <- notifyNewJoin(req.Code, req.Player)
	return res, nil
}

func (DiceServer) Notifications(req *gproto.RegisterNotifications, stream gproto.Game_NotificationsServer) error {
	logrus.Debug("register for notifications")

	actualGames.addStream(req.Code, stream)
	// <-stream.Context().Done()
	for notif := range notifications {
		actualGames.sendNotification(notif)
	}
	return nil
}

func (DiceServer) Movement(ctx context.Context, req *gproto.MovementRequest) (*gproto.Response, error) {
	switch req.Type {
	case gproto.MovementType_PICK:
		actualGames.pickDice(req.Code)
	}
	return nil, nil
}

func NewServer(port int) {
	server := DiceServer{
		host: fmt.Sprintf(":%d", port),
	}
	listener, err := net.Listen("tcp", server.host)
	if err != nil {
		logrus.Fatalf("Unable to listen '%s': %v\n", server.host, err)
	}

	var (
		opts = []grpc.ServerOption{}
		s    = grpc.NewServer(opts...)
	)
	gproto.RegisterGameServer(s, server)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	logrus.Infof("Server started at host %s\n", server.host)
	c := make(chan os.Signal, 1)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	// fmt.Println("Closing MongoDB connection")
	// db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
