package models

import "github.com/tecnologer/dicegame/server/models/gproto"

func newJoinResponse(player *gproto.Player, status bool) *gproto.Response {
	return &gproto.Response{
		Content: &gproto.Response_JoinResponse{
			JoinResponse: &gproto.PlayerUpdate{
				Player: player,
				Status: status,
			},
		},
	}
}

func notifyNewJoin(code string, player *gproto.Player) *notification {
	return &notification{
		code: code,
		response: &gproto.Response{
			Content: &gproto.Response_Notification{
				Notification: &gproto.NotificationResponse{
					Type: gproto.NotificationResponse_JOIN,
					Content: &gproto.NotificationResponse_PlayerJoin{
						PlayerJoin: &gproto.PlayerUpdate{
							Player: player,
						},
					},
				},
			},
		},
	}
}

func newGameResponse(code string) *gproto.Response {
	return &gproto.Response{
		Content: &gproto.Response_NewGameResponse{
			NewGameResponse: &gproto.NewGameResponse{
				Code: code,
			},
		},
	}
}
