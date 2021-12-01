package models

import (
	"rwby-adventures/config"
	"time"
)

type Player struct {
	DiscordID     string `gorm:"primary_key;column:discord_id"`
	IsNew         bool   `gorm:"column:is_new;not null"`
	Balance       int64  `gorm:"column:balance;not null"`
	BiddedBalance int64  `gorm:"column:bidded_balance;not null"`
	Level         int64  `gorm:"column:level;not null"`
	CP            int64  `gorm:"column:cp;not null"`
	MaxCP         int64  `gorm:"column:max_cp;not null"`
	CharLimit     int    `gorm:"column:max_char;not null"`
	Maxlootbox    int    `gorm:"column:max_lootbox;not null"`
	SelectedID    string `gorm:"column:selected_id;not null"`
	SelectedType  int    `gorm:"column:selected_type;not null"`
	Badges        int64  `gorm:"column:badges;not null"`
	Settings      int64  `gorm:"column:settings;not null"`
	Disabled      bool   `gorm:"column:disabled;not null"`
	Arms          int    `gorm:"column:arms;not null"`
	Minions       int    `gorm:"column:minions;not null"`
	Jar           int64  `gorm:"column:jar;not null"`

	// Foreign keys
	Missions      PlayerMission  `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Status        PlayerStatus   `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Shop          PlayerShop     `gorm:"foreignkey:DiscordID;references:DiscordID"`
	SelectedChar  Character      `gorm:"foreignkey:UserID;references:SelectedID"`
	SelectedGrimm Grimm          `gorm:"foreignkey:UserID;references:SelectedID"`
	LastBoxes     PlayerLootTime `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Gamble        PlayerGamble   `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Boxes         PlayerBoxes    `gorm:"foreignkey:DiscordID;references:DiscordID"`
	LimitedBoxes  []LimitedBoxes `gorm:"foreignkey:DiscordID"`
	SpecialBoxes  []SpecialBoxes `gorm:"foreignkey:DiscordID"`

	// Loaded later
	Characters    []Character `gorm:"foreignkey:DiscordID"`
	Grimms        []Grimm     `gorm:"foreignkey:DiscordID"`
	CharInMission Character
	GrimmInHunt   Grimm
}

func GetPlayer(id string) *Player {
	p := &Player{
		DiscordID: id,
	}
	e := config.Database.
		Preload("Status").
		Preload("Missions").
		Preload("Shop").
		Preload("SelectedChar").
		Preload("SelectedGrimm").
		Preload("LastBoxes").
		Preload("Gamble").
		Preload("LimitedBoxes").
		Preload("SpecialBoxes").
		Preload("Boxes").
		Find(p, id)
	if e.Error != nil {
		p = &Player{
			DiscordID:  id,
			IsNew:      true,
			CharLimit:  30,
			Maxlootbox: 3,
			Boxes: PlayerBoxes{
				Boxes: 1,
			},
			Missions: PlayerMission{
				DiscordID: id,
			},
			Status: PlayerStatus{
				DiscordID: id,
			},
			Shop: PlayerShop{
				DiscordID: id,
			},
		}
		config.Database.Create(p)
	}
	config.Database.Order(p.Status.OrderBy).Find(&p.Characters, "discord_id=?", p.DiscordID)
	config.Database.Order(p.Status.OrderBy).Find(&p.Grimms, "discord_id=?", p.DiscordID)
	config.Database.Order(p.Status.OrderBy).Find(&p.CharInMission, "discord_id=? and mission_in", p.DiscordID)
	config.Database.Order(p.Status.OrderBy).Find(&p.GrimmInHunt, "discord_id=? and hunt_in", p.DiscordID)
	return p
}

func (p *Player) CanDropLootBox() (canHe bool, reset bool) {
	lastTime := time.Unix(p.LastBoxes.Time, 0)
	if p.LastBoxes.Amount < p.Maxlootbox {
		canHe = true
		reset = false
		return
	} else if time.Since(lastTime).Hours() > 24 && p.LastBoxes.Amount == p.Maxlootbox {
		canHe = true
		reset = true
		return
	}
	return
}

func (p *Player) CanGamble() (canHe, reset bool) {
	lastTime := time.Unix(p.Gamble.Time, 0)
	if p.Gamble.Amount < 3 {
		canHe = true
		reset = false
		return
	} else if time.Since(lastTime).Hours() > 24 && p.Gamble.Amount >= 3 {
		canHe = true
		reset = true
		return
	}
	return
}

func (p *Player) CanDungeon() bool {
	return time.Now().Unix()-p.Status.LastDungeon > 18000
}

func (p *Player) MaxChar() int {
	return p.CharLimit + p.Shop.Extensions
}

func (p *Player) TotalBalance() int64 {
	return p.Balance - p.BiddedBalance
}

func (p *Player) Lootbox() (s int) {
	s += p.Boxes.Boxes
	s += p.Boxes.RareBoxes
	s += p.Boxes.GrimmBoxes
	s += p.Boxes.RareGrimmBoxes
	s += len(p.LimitedBoxes)
	s += len(p.SpecialBoxes)
	return
}

func (p *Player) CharAmount() int {
	return len(p.Characters)
}

func (p *Player) GrimmAmount() int {
	return len(p.Grimms)
}
