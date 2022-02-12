package models

import (
	"github.com/tecnologer/dicegame/server/models/gproto"
	"github.com/tecnologer/dicegame/src/models"
)

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

func diceToProtoDice(dice *models.Dice) *gproto.Dice {
	sides := make([]gproto.DiceSide, 6)
	for i, side := range dice.Sides {
		sides[i] = gproto.DiceSide(side)
	}
	return &gproto.Dice{
		Level: gproto.Dice_DiceLevel(dice.Level),
		Sides: sides,
	}
}

func protoDiceToDice(dice *gproto.Dice) *models.Dice {
	var sides [6]models.DiceSide
	for i, side := range dice.Sides {
		sides[i] = models.DiceSide(side)
	}
	return &models.Dice{
		Level: models.DiceLevel(dice.Level),
		Sides: sides,
	}
}

func newDiceResponse(dice *gproto.Dice) *gproto.Response {
	return &gproto.Response{
		Content: &gproto.Response_Movement{
			Movement: &gproto.MovementResponse{
				Type: gproto.MovementType_PICK,
				Movement: &gproto.MovementResponse_Dice{
					Dice: dice,
				},
			},
		},
	}
}

func notifyPickDice(code string, dice *gproto.Dice) *notification {
	return &notification{
		code: code,
		response: &gproto.Response{
			Content: &gproto.Response_Notification{
				Notification: &gproto.NotificationResponse{
					Type: gproto.NotificationResponse_PICK,
					Content: &gproto.NotificationResponse_DicePick{
						DicePick: dice,
					},
				},
			},
		},
	}
}

func newRollDiceResponse(side models.DiceSide) *gproto.Response {
	return &gproto.Response{
		Content: &gproto.Response_Movement{
			Movement: &gproto.MovementResponse{
				Type: gproto.MovementType_ROLL,
				Movement: &gproto.MovementResponse_RollDice{
					RollDice: &gproto.RollDice{
						Side: gproto.DiceSide(side),
					},
				},
			},
		},
	}
}

func notifyRollDice(code string, dice *gproto.Dice, side models.DiceSide) *notification {
	return &notification{
		code: code,
		response: &gproto.Response{
			Content: &gproto.Response_Notification{
				Notification: &gproto.NotificationResponse{
					Type: gproto.NotificationResponse_PICK,
					Content: &gproto.NotificationResponse_RollDice{
						RollDice: &gproto.RollDice{
							Dice: dice,
							Side: gproto.DiceSide(side),
						},
					},
				},
			},
		},
	}
}
