package dice

import (
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/models"
)

type turn struct {
	number     uint
	Player     *models.Player
	Brains     uint
	Shots      uint
	Footprints uint
	Dices      []*models.Dice
}

func newTurn() *turn {
	return &turn{
		Player:     nil,
		Brains:     0,
		Shots:      0,
		Footprints: 0,
		Dices:      make([]*models.Dice, 0),
	}
}

//save sets the scores to the player
func (t *turn) save() {
	t.Player.Brains += t.Brains
	lFmt.Printf("%s tu turno termina, sumaste %d cerebros, tienes un total de %d\n",
		t.Player.Name,
		t.Brains,
		t.Player.Brains,
	)
}

//setPlayer set the player's turn and reset the turn's scores
func (t *turn) setPlayer(player *models.Player) {
	t.number++
	t.Player = player
	t.Brains = 0
	t.Shots = 0
	t.Footprints = 0
	t.Dices = []*models.Dice{}
}

func (t *turn) Won() bool {
	return t.getPlayerBrains() >= constants.BrainCount
}

func (t *turn) Lost() bool {
	return t.Shots >= constants.ShotgunsCount
}

func (t *turn) getPlayerBrains() uint {
	return t.Player.Brains + t.Brains
}

func (t *turn) isComputer() bool {
	return t.Player.IsAI
}
