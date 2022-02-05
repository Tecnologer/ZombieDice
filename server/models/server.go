package models

import (
	"context"
	"fmt"

	"github.com/tecnologer/dicegame/server/models/gproto"
	"google.golang.org/grpc"
)

type DiceServer struct {
	host string
}

func (s *DiceServer) NewGame(ctx context.Context, in *gproto.NewGameRequest, opts ...grpc.CallOption) (*gproto.NewGameResponse, error) {
	return nil, nil
}
func (s *DiceServer) Join(ctx context.Context, in *gproto.JoinRequest, opts ...grpc.CallOption) (*gproto.JoinResponse, error) {
	return nil, nil
}

func NewServer(port int) *DiceServer {

	server := &DiceServer{
		host: fmt.Sprintf(":%d", port),
	}
	return server
}
