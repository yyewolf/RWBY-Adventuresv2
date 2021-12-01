package models

import (
	"rwby-adventures/config"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	DiscordID    string `gorm:"primary_key;column:discord_id"`
	IsNew        bool   `gorm:"column:is_new;not null"`
	Balance      int64  `gorm:"column:balance;not null"`
	Level        int64  `gorm:"column:level;not null"`
	CP           int64  `gorm:"column:cp;not null"`
	MaxCP        int64  `gorm:"column:max_cp;not null"`
	CharLimit    int    `gorm:"column:max_char;not null"`
	Maxlootbox   int    `gorm:"column:max_lootbox;not null"`
	SelectedID   string `gorm:"column:selected_id;not null"`
	SelectedType int    `gorm:"column:selected_type;not null"`
	Badges       int64  `gorm:"column:badges;not null"`
	Settings     int64  `gorm:"column:settings;not null"`
	Disabled     bool   `gorm:"column:disabled;not null"`

	// Foreign keys
	Missions      PlayerMission   `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Status        PlayerStatus    `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Shop          PlayerShop      `gorm:"foreignkey:DiscordID;references:DiscordID"`
	SelectedChar  PlayerCharacter `gorm:"foreignkey:UserID;references:SelectedID"`
	SelectedGrimm PlayerGrimm     `gorm:"foreignkey:UserID;references:SelectedID"`
	LastBoxes     PlayerLootTime  `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Gamble        PlayerGamble    `gorm:"foreignkey:DiscordID;references:DiscordID"`
	LimitedBoxes  []LimitedBoxes  `gorm:"foreignkey:DiscordID"`
	SpecialBoxes  []SpecialBoxes  `gorm:"foreignkey:DiscordID"`

	// Loaded later
	Characters    []PlayerCharacter `gorm:"foreignkey:DiscordID"`
	Grimms        []PlayerGrimm     `gorm:"foreignkey:DiscordID"`
	CharInMission PlayerCharacter
	GrimmInHunt   PlayerGrimm
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
		Find(p, id)
	if e.Error != nil {
		p = &Player{
			DiscordID: id,
			IsNew:     true,
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
