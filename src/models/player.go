package models

type Player struct {
	Name   string
	Brains uint
	Shots  uint
	IsAI   bool
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Brains: 0,
		Shots:  0,
	}
}

func NewPlayerIA(name string) *Player {
	return &Player{
		Name:   name,
		Brains: 0,
		Shots:  0,
		IsAI:   true,
	}
}
