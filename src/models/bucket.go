package models

import (
	"fmt"
	"strings"

	"github.com/tecnologer/dicegame/src/constants"
	"github.com/tecnologer/dicegame/src/utils"
)

type Bucket []*Dice

func NewBucket(dices [constants.GameDiceCount]*Dice) *Bucket {
	bucket := new(Bucket)
	(*bucket) = append(*bucket, dices[0:]...)
	return bucket
}

func (b *Bucket) Clear(dices ...Dice) {
	*b = []*Dice{}
}

func (b *Bucket) HasEnougthDices() bool {
	return len(*b) >= constants.DicePerRoll
}

func (b *Bucket) PickDice(i int) (dice *Dice) {
	dice = (*b)[i]
	*b = append((*b)[:i], (*b)[i+1:]...)
	return
}

func (b *Bucket) PickRandomDice() *Dice {
	if len(*b) < 1 {
		return nil
	}

	limit := len(*b) - 1
	if limit < 0 {
		return nil
	} else if limit == 0 {
		return b.PickDice(0)
	}

	pickedIndex := utils.GetRandInt(limit)
	return b.PickDice(pickedIndex)
}

func (b *Bucket) AddDice(dices ...*Dice) {
	*b = append(*b, dices...)
}

func (b Bucket) String() string {
	output := make([]string, len(b))
	for i, d := range b {
		output[i] = fmt.Sprint(d)
	}

	return strings.Join(output, ", ")
}
