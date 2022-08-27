package game

import (
	"fmt"
	"math/rand"

	"github.com/yyewolf/gosf"
)

const (
	choiceMoney = iota
	choiceCCBox
)

type DungeonChoice struct {
	Index         int    `json:"index"`
	Type          int    `json:"type,omitempty"`
	Amount        int    `json:"-"`
	ButtonMessage string `json:"message,omitempty"`
	ChoiceMessage string `json:"-"`
}

func (c *DungeonChoice) Text() {
	switch c.Type {
	case choiceMoney:
		c.ButtonMessage = "Lien(s)"
		c.ChoiceMessage = "Ambrosius gave you %dⱠ (Liens) !"
	case choiceCCBox:
		c.ButtonMessage = "Classic Character Box(es)"
		c.ChoiceMessage = "Ambrosius gave you %d classic character boxes !"
	}
}

func (c *DungeonChoice) Amounts() {
	switch c.Type {
	case choiceMoney:
		c.Amount = rand.Intn(350) + 100
	case choiceCCBox:
		c.Amount = rand.Intn(3) + 1
	}
}

func generateChoices(amount int) (choices []*DungeonChoice) {
	for i := 0; i < amount; i++ {
		rdmChoice := rand.Intn(2)
		choice := &DungeonChoice{
			Type: rdmChoice,
		}
		choice.Text()
		choice.Amounts()
		choice.Index = i
		choices = append(choices, choice)
	}
	return
}

func (d *Dungeon) MakeChoice(choice int) *gosf.Message {
	currentCell := d.Grid[d.Position.Y][d.Position.X]

	if currentCell.Type != tileAmbrosius {
		return gosf.NewFailureMessage("nice try...")
	}

	choosenChoice := d.Grid[d.Position.Y][d.Position.X].Choices[choice]

	response := &gosf.Message{
		Success: true,
		Text:    fmt.Sprintf(choosenChoice.ChoiceMessage, choosenChoice.Amount),
	}

	switch choosenChoice.Type {
	case choiceMoney:
		d.Rewards.Lien += choosenChoice.Amount
	case choiceCCBox:
		d.Rewards.CCBox += choosenChoice.Amount
	}

	// reset cell
	currentCell.Type = tileFloor
	currentCell.Choices = nil

	return response
}
