package dice

import (
	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/models"
	"github.com/tecnologer/dicegame/src/utils"
)

var lFmt lang.DiceLanguage

type Game struct {
	turn      *turn
	Players   []*models.Player
	Dices     [constants.GameDiceCount]*models.Dice
	Bucket    *models.Bucket
	IsStopped bool
}

func NewGame(playersName ...string) (game *Game) {
	lFmt = lang.GetCurrent()

	var players []*models.Player
	if len(playersName) < 1 {
		panic(lFmt.Sprintf("It's necessary at least one player."))
	}
	if len(playersName) > 0 {
		players = make([]*models.Player, len(playersName))
		for i, player := range playersName {
			players[i] = models.NewPlayer(player)
		}
	}

	if len(players) == 1 {
		players = append(players, models.NewPlayerIA(lFmt.Sprintf("computer")))
	}

	game = &Game{
		Players: players,
		Dices:   getDicesNewGame(),
		turn:    newTurn(),
	}

	game.Bucket = models.NewBucket(game.Dices)
	game.NextPlayer()
	return
}

func getDicesNewGame() [constants.GameDiceCount]*models.Dice {
	var dices [constants.GameDiceCount]*models.Dice
	levels := [constants.GameDiceCount]models.DiceLevel{
		models.LevelEasy,
		models.LevelEasy,
		models.LevelEasy,
		models.LevelEasy,
		models.LevelEasy,
		models.LevelEasy,
		models.LevelMedium,
		models.LevelMedium,
		models.LevelMedium,
		models.LevelMedium,
		models.LevelHard,
		models.LevelHard,
		models.LevelHard,
	}

	for i, level := range levels {
		dices[i] = models.NewDice(level)
	}
	return dices
}

func (g *Game) PickDices() {
	if !g.Bucket.HasEnougthDices() {
		g.resetBucket()
	}
	for i := len(g.turn.Dices); i < constants.DicePerRoll; i++ {
		dice := g.Bucket.PickRandomDice()

		if dice == nil {

			dice = g.Bucket.PickRandomDice()
		}
		g.turn.Dices = append(g.turn.Dices, dice)
	}
}

func (g *Game) GetDicesPicked() [3]*models.Dice {
	picked := [3]*models.Dice{}
	i := 0
	for _, dice := range g.Dices {
		if !dice.Picked {
			continue
		}

		picked[i] = dice
		i++
	}

	return picked
}

func (g *Game) Start() {
	for !g.IsOver() {
		lFmt.Printf("Turn of player: %s\n", g.turn.Player.Name)

		g.PickDices()

		logrus.Debugf("dices in bucket: %d\n", len(*g.Bucket))

		lFmt.Printlnf("Selected dices:")
		printDices(g.turn.Dices...)

		g.rollDices()

		if g.IsNextPlayer() {
			g.NextPlayer()
		}
	}
}

func (g *Game) SetBrain() {
	g.turn.Brains++
}

func (g *Game) SetShotgun() {
	g.turn.Shots++
}

func (g *Game) IsOver() bool {
	return g.turn.Won() || g.IsStopped
}

func (g *Game) IsNextPlayer() bool {
	if g.turn.Won() {
		lFmt.Printf("The winner is %s\n with %d brains in %d turns\n",
			g.turn.Player.Name,
			g.turn.getPlayerBrains(),
			g.turn.number,
		)
		g.IsStopped = true
		return true
	}

	if g.turn.Lost() {
		lFmt.Printf("Your turn ends with %d shotguns\n", g.turn.Shots)
		utils.AskEnter("Press enter to end your turn...")
		return true
	}

	playerWantsContinue := true
	// lFmt.Printf("en este turno tienes %d cerebros y %d disparos.\n",
	// 	g.turn.Brains,
	// 	g.turn.Shots,
	// )
	lFmt.Printf("%s: on this turn, you've %d brains and %d shotguns. Total brains: %d.\n",
		g.turn.Player.Name,
		g.turn.Brains,
		g.turn.Shots,
		g.turn.Player.Brains,
	)

	if g.turn.isComputer() {
		playerWantsContinue = wanstAIEndTurn(g.turn.Player)
	} else {
		playerWantsContinue = utils.AskBoolf("Do you want to end your turn? (Yes, Default: No): ",
			false,
		)
	}

	if playerWantsContinue {
		g.turn.save()
	}

	return playerWantsContinue

}

func (g *Game) NextPlayer() {
	defer g.resetBucket()

	if g.turn.Player == nil {
		g.turn.setPlayer(g.Players[0])
		return
	}

	for i, player := range g.Players {
		if g.turn.Player == player {
			if i+1 <= len(g.Players)-1 {
				g.turn.setPlayer(g.Players[i+1])
			} else {
				g.turn.setPlayer(g.Players[0])
			}
			break
		}
	}
}

func (g *Game) UnpickDice(i int) {
	g.turn.Dices = append(g.turn.Dices[:i], g.turn.Dices[i+1:]...)
}

//rollDices rolls the picked dices
func (g *Game) rollDices() {
	newPicked := make([]*models.Dice, 0)
	for i, dice := range g.turn.Dices {
		if dice == nil {
			continue
		}

		side := dice.Roll()
		lFmt.Printf("Dice #%d gets %s\n", i+1, side)
		if side == models.Brain {
			g.SetBrain()
		} else if side == models.Shotgun {
			g.SetShotgun()
		} else if side == models.Footprints {
			newPicked = append(newPicked, dice)
		}
	}
	g.turn.Dices = newPicked
}

func (g *Game) resetBucket() {
	g.Bucket.Clear()
	g.Bucket.AddDice(g.Dices[:]...)
}
func printDices(dices ...*models.Dice) {
	for _, dice := range dices {
		dice.Println()
	}
}
