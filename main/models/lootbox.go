package models

import "github.com/jinzhu/gorm"

type PlayerLootTime struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`
	Amount    int    `gorm:"column:amount;not null"`
	Time      int64  `gorm:"column:time;not null"`
}

type PlayerGamble struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`
	Amount    int    `gorm:"column:amount;not null"`
	Time      int64  `gorm:"column:time;not null"`
}

type LimitedBoxes struct {
	gorm.Model
	DiscordID string `gorm:"column:discord_id"`
	For       string `gorm:"column:for;not null"`
	Type      int    `gorm:"column:type;not null"`
}

type SpecialBoxes struct {
	gorm.Model
	DiscordID string `gorm:"column:discord_id"`
	For       string `gorm:"column:for;not null"`
}
