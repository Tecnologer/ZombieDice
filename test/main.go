package main

import (
	dice "github.com/tecnologer/dicegame/src"
)

func main() {
	game := dice.NewGame("player1", "player2")
	game.Start()
}
