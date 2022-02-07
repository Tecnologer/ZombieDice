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

type DiceServer struct {
	gproto.UnimplementedGameServer
	host string
}

func (s DiceServer) NewGame(ctx context.Context, req *gproto.NewGameRequest) (*gproto.NewGameResponse, error) {
	code := CreateGame(req.Password, int(req.LimitPlayers))
	res := &gproto.NewGameResponse{
		Code: code,
		Host: s.host,
	}

	msg := "encrypted"
	if req.Password == "" {
		msg = "uncrypted"
	}
	logrus.Infof("New %s game created with code %s", msg, code)
	JoinToGame(code, req.Password, req.Player)
	return res, nil
}

func (s DiceServer) Join(ctx context.Context, req *gproto.JoinRequest) (*gproto.JoinResponse, error) {
	err := JoinToGame(req.Code, req.Password, req.Player)
	if err != nil {
		return nil, err
	}
	res := &gproto.JoinResponse{
		Player: req.Player,
		Status: true,
	}
	return res, nil
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
