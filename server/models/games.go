package models

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	lang "github.com/tecnologer/dicegame/language"
	"github.com/tecnologer/dicegame/server/models/gproto"
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
	streams      map[string]gproto.Game_NotificationsServer
}
type games struct {
	current map[string]*game
}

type notification struct {
	response *gproto.Response
	code     string
	player   *gproto.Player
}

func (g *games) isThereGame(key string) bool {
	_, exists := g.current[key]
	return exists
}
func (g *games) existsGame(key string) error {
	if !g.isThereGame(key) {
		return errors.Errorf("there is not game with code %s", key)
	}
	return nil
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
		streams:      make(map[string]gproto.Game_NotificationsServer),
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

func (g *games) checkPassword(key, pwd string) bool {
	return g.current[key].password == pwd
}

func (g *games) sendNotification(notif *notification) {
	if err := g.existsGame(notif.code); err != nil {
		logrus.WithError(err).Warn("notification couldn't send")
		return
	}

	for playerName, stream := range g.current[notif.code].streams {
		if playerName == notif.player.Name {
			continue
		}
		e := stream.Send(notif.response)
		if e != nil {
			logrus.WithError(e).Warn("the notification couldn't send.")
		}
	}
}

func (g *games) addStream(key string, player *gproto.Player, stream gproto.Game_NotificationsServer) {
	if !g.isThereGame(key) {
		logrus.Warnf("there is not game with code %s, couldn't register for notif.")
		return
	}

	g.current[key].streams[player.Name] = stream
}

func (g *games) removeStream(key string, player *gproto.Player) {
	if err := g.existsGame(key); err != nil {
		logrus.WithError(err).Warn("notification couldn't be removed")
		return
	}

	delete(g.current[key].streams, player.Name)
}

func (g *games) pickDice(key string) (*gproto.Dice, error) {
	key = strings.ToUpper(key)
	if err := g.existsGame(key); err != nil {
		return nil, errors.Wrap(err, "dice not picked")
	}

	dice := g.current[key].Game.Bucket.PickRandomDice()
	return diceToProtoDice(dice), nil
}

func (g *games) nextPlayer(key string) (end *dice.Turn, next *dice.Turn, err error) {
	key = strings.ToUpper(key)
	if err := g.existsGame(key); err != nil {
		return nil, nil, errors.Wrap(err, "cannot select next player")
	}

	//get score of the curren player
	end = g.current[key].GetTurn()

	//select next player
	g.current[key].Game.NextPlayer()
	next = g.current[key].GetTurn()
	return
}

func InitGames() {
	lock.Lock()
	defer lock.Unlock()

	tFmt = lang.GetCurrent()
	actualGames = &games{
		current: make(map[string]*game),
	}
}

func JoinToGame(key, pwd string, player *gproto.Player) error {
	if player == nil {
		return errors.New(tFmt.Sprintf("Player is required"))
	}

	if err := actualGames.existsGame(key); err != nil {
		return errors.Wrap(err, "cannot join")
	}

	if !actualGames.checkPassword(key, pwd) {
		return errors.New(tFmt.Sprintf("The key '%s' or the password don't match.", key))
	}

	if actualGames.isLimitReached(key) {
		return errors.New(tFmt.Sprintf("The game '%s' reached its limit of %d players.", key, actualGames.getLimitPlayers(key)))
	}

	err := actualGames.joinPlayer(key, parseProtoPlayerToPlayer(player))
	if err != nil {
		return err
	}
	logrus.Infof("Player '%s' joined to %s", player.Name, key)
	return nil
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

func parseProtoPlayerToPlayer(p *gproto.Player) *dicemodels.Player {
	return dicemodels.NewPlayer(p.Name)
}
