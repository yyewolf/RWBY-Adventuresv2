package models

import "rwby-adventures/config"

type PlayerStatus struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	LastXP      int64 `gorm:"column:last_xp;not null"`
	Voted       bool  `gorm:"column:has_voted;not null"`
	DailyStreak int   `gorm:"column:daily_streak;not null"`
	LastOpening int64 `gorm:"column:last_opening;not null"`

	LastReport  int64 `gorm:"column:last_report;not null"`
	LastDungeon int64 `gorm:"column:last_dungeon;not null"`
}

func (s *PlayerStatus) Save() {
	config.Database.Save(s)
}
