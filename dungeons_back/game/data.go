package game

import (
	"rwby-adventures/microservices"
)

const (
	tileFloor = iota
	tileWall
	tileFow
	tilePlayer
	tileMoney
	tileEscape
	tileVoid
	tileEnnemy
	tileAmbrosius
	tileArm
	tileMinion
	tileTPWall
)

const (
	findMoney  = "You found %dⱠ (Liens) !"
	findEnnemy = "You found an ennemy and lost %d HP !"
	findArm    = "You found an arm !"
	findMinion = "You found a minion !"
	findWall   = "What the ?"
)

var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

type PlayerPosition struct {
	X int
	Y int
}

type DungeonCell struct {
	ID      string           `json:"id"`
	Type    int              `json:"type"`
	Amount  int              `json:"amount,omitempty"`
	Message string           `json:"message,omitempty"`
	Damages int              `json:"damages,omitempty"`
	Choices []*DungeonChoice `json:"choices,omitempty"`
}

type Dungeon struct {
	Grid     [][]*DungeonCell
	Rewards  *microservices.DungeonReward
	Height   int
	Width    int
	Position *PlayerPosition
	Health   int

	SecretRoom   [][]*DungeonCell
	Temp         [][]*DungeonCell
	Win          bool
	InSecretRoom bool
	PreviousPos  *PlayerPosition
}
