package dice

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/models"
	"github.com/tecnologer/dicegame/src/utils"
)

type Game struct {
	turn      *turn
	Players   []*models.Player
	Dices     [constants.GameDiceCount]*models.Dice
	Bucket    *models.Bucket
	IsStopped bool
}

func NewGame(playersName ...string) (game *Game) {
	var players []*models.Player
	if len(playersName) < 2 {
		panic("se requieren al menos dos jugadores")
	}
	if len(playersName) > 0 {
		players = make([]*models.Player, len(playersName))
		for i, player := range playersName {
			players[i] = models.NewPlayer(player)
		}
	} else {
		players = []*models.Player{models.NewPlayer("Player1"), models.NewPlayer("Player2")}
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
		fmt.Printf("Turno del jugador: %s\n", g.turn.Player.Name)

		g.PickDices()

		fmt.Printf("dices in bucket: %d\n", len(*g.Bucket))

		fmt.Println("Dados seleccionados:")
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
		fmt.Printf("El ganador es %s\n con %d cerebros en %d turnos\n",
			g.turn.Player.Name,
			g.turn.getPlayerBrains(),
			g.turn.number,
		)
		g.IsStopped = true
		return true
	}

	if g.turn.Lost() {
		color.Red("perdiste el turno con %d disparos", g.turn.Shots)
		return true
	}

	// playerWantsContinue := true
	// fmt.Printf("en este turno tienes %d cerebros y %d disparos.\n",
	// 	g.turn.Brains,
	// 	g.turn.Shots,
	// )
	playerWantsContinue := utils.AskBoolf("%s: en este turno llevas %d cerebros y %d disparos. quieres continuar? (Si,No): ",
		true,
		g.turn.Player.Name,
		g.turn.Brains,
		g.turn.Shots,
	)

	if !playerWantsContinue {
		g.turn.save()
	}

	return !playerWantsContinue

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
		fmt.Printf("dado #%d obtuvo %s\n", i+1, side)
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
