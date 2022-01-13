package models

import (
	"fmt"
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"strings"
	"time"
)

const (
	CharType = iota
	GrimmType
)

type CharacterStats struct {
	CharID      string  `gorm:"primary_key;column:id"`
	Value       float64 `gorm:"column:value;not null"`
	Health      int     `gorm:"column:health;not null"`
	Armor       int     `gorm:"column:armor;not null"`
	Damage      int     `gorm:"column:damage;not null"`
	Healing     int     `gorm:"column:healing;not null"`
	DodgeChance int     `gorm:"column:dodge_chance;not null"`
}

type Character struct {
	CharID        string    `gorm:"primary_key;column:id"`
	UserID        string    `gorm:"column:user_id;not null"`
	Name          string    `gorm:"column:name;not null"`
	Level         int       `gorm:"column:level;not null"`
	XP            int64     `gorm:"column:xp;not null"`
	XPCap         int64     `gorm:"column:xp_max;not null"`
	Rarity        int       `gorm:"column:rarity;not null"`
	InMission     bool      `gorm:"column:in_mission;not null"`
	IsInFavorites bool      `gorm:"column:is_in_favorites;not null"`
	Buffs         int       `gorm:"column:buffs;not null"`
	OwnedAt       time.Time `gorm:"column:owned_at;not null"`
	// Foreign keys
	Stats CharacterStats `gorm:"foreignkey:CharID"`
}

func (c Character) ToRealChar() chars.CharacterStruct {
	i := 0
	for i = range config.BaseCharacters {
		if config.BaseCharacters[i].Name == c.Name {
			break
		}
	}

	returnChar := config.BaseCharacters[i]
	returnChar.CustomID = c.CharID
	returnChar.Stats.Armor = c.Stats.Armor
	returnChar.Stats.Damage = c.Stats.Damage
	returnChar.Stats.Healing = c.Stats.Healing
	returnChar.Stats.Health = c.Stats.Health
	returnChar.Stats.Value = c.Stats.Value
	returnChar.Rarity = c.Rarity
	returnChar.Level = c.Level

	for i := range returnChar.Attacks {
		returnChar.Attacks[i].LastUsed = -5
	}
	returnChar.Semblance.CustomData = make(map[string]interface{})
	returnChar.Semblance.CustomData["lastTime"] = 0

	return returnChar
}

func CharRarityToColor(Rarity int) int {
	EmbedColor := 0
	switch Rarity {
	case 0: // Common
		EmbedColor = 0x808080
		break
	case 1: // Uncommon
		EmbedColor = 0x7CFC00
		break
	case 2: // Rare
		EmbedColor = 0x87CEEB
		break
	case 3: // Very Rare
		EmbedColor = 0xBA55D3
		break
	case 4: // Legendary
		EmbedColor = 0xFFD700
		break
	case 5: // Collector
		EmbedColor = 0xFF0000
		break
	}
	return EmbedColor
}

func (c *Character) RarityString() (x string) {
	switch c.Rarity {
	case 0: // Common
		x = "□ Common"
		break
	case 1: // Uncommon
		x = "◇ Uncommon"
		break
	case 2: // Rare
		x = "♡ Rare"
		break
	case 3: // Very Rare
		x = "♤ Very Rare"
		break
	case 4: // Legendary
		x = "♧ Legendary"
		break
	case 5: // Collector
		x = "☆ Collector"
		break
	}

	for i := 0; i < c.Buffs; i++ {
		x += "+"
	}
	return x
}

func (c *Character) FullString() string {
	return fmt.Sprintf("%s level %d (%d/%dXP) %s (%.2f%%)", c.RarityString(), c.Level, c.XP, c.XPCap, c.Name, c.Stats.Value)
}

func (c *Character) CheckConditions(charname string, level int, valueabove float64, valuebelow float64, rarity int, buffs int) bool {
	if charname != "" && !(strings.Contains(strings.ToLower(c.Name), strings.ToLower(charname))) {
		return false
	}
	if level != 1000 && c.Level != level {
		return false
	}
	if c.Stats.Value < valueabove {
		return false
	}
	if c.Stats.Value > valuebelow {
		return false
	}
	if rarity != -1 && c.Rarity != rarity {
		return false
	}
	if buffs != 0 && c.Buffs != buffs {
		return false
	}
	return true
}
