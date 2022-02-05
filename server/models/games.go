package models

import (
	"errors"
	"fmt"
	"sync"

	lang "github.com/tecnologer/dicegame/language"
	dice "github.com/tecnologer/dicegame/src"
	dicemodels "github.com/tecnologer/dicegame/src/models"
	"github.com/tecnologer/dicegame/src/utils"
)

const keyLen = 6

var (
	tFmt        lang.DiceLanguage
	lock        = &sync.Mutex{}
	actualGames *games
)

type game struct {
	*dice.Game
	password     string
	limitPlayers int
}
type games struct {
	current map[string]*game
}

func (g *games) isThereGame(key string) bool {
	_, exists := g.current[key]
	return exists
}

func (g *games) newGame(key, pwd string, limitPlayers int) {
	//if the game is already created
	if g.isThereGame(key) {
		return
	}
	g.current[key] = &game{
		Game:         dice.NewGame(),
		password:     pwd,
		limitPlayers: limitPlayers,
	}

}

func (g *games) joinPlayer(key string, player *dicemodels.Player) error {
	return g.current[key].AddPlayer(player)
}
func (g *games) isLimitReached(key string) bool {
	limit := g.getLimitPlayers(key)
	//there is not limit
	if limit < 1 {
		return false
	}

	return len(g.current[key].Game.Players) >= limit
}

func (g *games) getLimitPlayers(key string) int {
	return g.current[key].limitPlayers
}

func InitGames() {
	lock.Lock()
	defer lock.Unlock()

	tFmt = lang.GetCurrent()
	actualGames = &games{
		current: make(map[string]*game),
	}
}

func JoinToGame(key string, player *dicemodels.Player) error {
	if !actualGames.isThereGame(key) {
		return errors.New(tFmt.Sprintf("There isn't a game with key '%s'", key))
	}

	if actualGames.isLimitReached(key) {
		return errors.New(tFmt.Sprintf("The game '%s' reached its limit of %d players.", key, actualGames.getLimitPlayers(key)))
	}

	return actualGames.joinPlayer(key, player)
}

func CreateGame(pwd string, limitPlayer int) string {
	lock.Lock()
	defer lock.Unlock()

	key := generateKeyGame()
	actualGames.newGame(key, pwd, limitPlayer)

	return key
}

func generateKeyGame() (key string) {
	for len(key) < keyLen {
		if utils.GetRandInt(9) >= 5 {
			r := rune(utils.GetRandIntRange(65, 90))
			key += string(r)
			continue
		}

		key += fmt.Sprint(utils.GetRandInt(9))
	}
	return key
}
