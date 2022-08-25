package models

import "rwby-adventures/config"

type PlayerStats struct {
	DiscordID string `gorm:"primary_key;column:id"`

	ArenasCompleted int `gorm:"column:arenas_completed;not null"`
	BattlesWon      int `gorm:"column:battles_won;not null"`
	BattlesLost     int `gorm:"column:battles_lost;not null"`
	LootboxOpened   int `gorm:"column:lootbox_opened;not null"`
	RoleplaySent    int `gorm:"column:roleplay_sent;not null"`
	DungeonDone     int `gorm:"column:dungeon_done;not null"`
	MarketSold      int `gorm:"column:market_sold;not null"`
	MarketBought    int `gorm:"column:market_bought;not null"`
}

func (p *PlayerStats) Save() {
	config.Database.Save(p)
}
