package models

type Player struct {
	Name   string
	Brains uint
	Shots  uint
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Brains: 0,
		Shots:  0,
	}
}
