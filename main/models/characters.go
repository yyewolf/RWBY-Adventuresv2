package models

import "github.com/jinzhu/gorm"

type PlayerCharacterStats struct {
	gorm.Model
	CharID      string  `gorm:"primary_key;column:char_id"`
	Value       float64 `gorm:"column:value;not null"`
	Health      int     `gorm:"column:health;not null"`
	Armor       int     `gorm:"column:armor;not null"`
	Damage      int     `gorm:"column:damage;not null"`
	Healing     int     `gorm:"column:healing;not null"`
	DodgeChance int     `gorm:"column:dodge_chance;not null"`
}

type PlayerCharacter struct {
	gorm.Model
	CharID        string `gorm:"primary_key;column:char_id"`
	UserID        string `gorm:"column:user_id;not null"`
	Name          string `gorm:"column:name;not null"`
	Level         int    `gorm:"column:level;not null"`
	XP            int64  `gorm:"column:xp;not null"`
	XPCap         int64  `gorm:"column:xp_max;not null"`
	Rarity        int    `gorm:"column:rarity;not null"`
	InMission     bool   `gorm:"column:in_mission;not null"`
	IsInFavorites bool   `gorm:"column:is_in_favorites;not null"`
	Buffs         int    `gorm:"column:buffs;not null"`
	// Foreign keys
	Stats PlayerCharacterStats `gorm:"foreignkey:CharID"`
}
