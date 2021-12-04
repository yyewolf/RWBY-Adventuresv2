package models

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/grimms"
	"time"
)

type GrimmStat struct {
	GrimmID     string  `gorm:"primary_key;column:id"`
	Value       float64 `gorm:"column:value;not null"`
	Health      int     `gorm:"column:health;not null"`
	Armor       int     `gorm:"column:armor;not null"`
	Damage      int     `gorm:"column:damage;not null"`
	Healing     int     `gorm:"column:healing;not null"`
	DodgeChance int     `gorm:"column:dodge_chance;not null"`
}

type Grimm struct {
	GrimmID       string    `gorm:"primary_key;column:id"`
	UserID        string    `gorm:"column:user_id;not null"`
	Name          string    `gorm:"column:name;not null"`
	Level         int       `gorm:"column:level;not null"`
	XP            int64     `gorm:"column:xp;not null"`
	XPCap         int64     `gorm:"column:xp_max;not null"`
	Rarity        int       `gorm:"column:rarity;not null"`
	InHunt        bool      `gorm:"column:in_hunt;not null"`
	IsInFavorites bool      `gorm:"column:is_in_favorites;not null"`
	Buffs         int       `gorm:"column:buffs;not null"`
	OwnedAt       time.Time `gorm:"column:owned_at;not null"`
	// Foreign keys
	Stats GrimmStat `gorm:"foreignkey:GrimmID"`
}

func (c *Grimm) ToRealGrimm() grimms.GrimmStruct {
	i := 0
	for i = range config.BaseGrimms {
		if config.BaseGrimms[i].Name == c.Name {
			break
		}
	}

	returnGrimm := config.BaseGrimms[i]
	returnGrimm.CustomID = c.GrimmID
	returnGrimm.Stats.Armor = c.Stats.Armor
	returnGrimm.Stats.Damage = c.Stats.Damage
	returnGrimm.Stats.Healing = c.Stats.Healing
	returnGrimm.Stats.Health = c.Stats.Health
	returnGrimm.Stats.Value = c.Stats.Value
	returnGrimm.Rarity = c.Rarity
	returnGrimm.Level = c.Level

	for i := range returnGrimm.Attacks {
		returnGrimm.Attacks[i].LastUsed = -5
	}
	returnGrimm.Special.CustomData = make(map[string]interface{})
	returnGrimm.Special.CustomData["lastTime"] = 0

	return returnGrimm
}

func (c *Grimm) RarityString() (x string) {
	switch c.Rarity {
	case 0: // Normal
		x = "□ Normal"
	case 1: // Abnormal
		x = "◇ Abnormal"
	case 2: // Sparse
		x = "♡ Sparse"
	case 3: // Freaky
		x = "♤ Freaky"
	case 4: // Mysterious
		x = "♧ Mysterious"
	case 5: // Bloody
		x = "☆ Bloody"
	}

	for i := 0; i < c.Buffs; i++ {
		x += "+"
	}
	return x
}

func (g *Grimm) FullString() string {
	return fmt.Sprintf("`%s level %d (%d/%dXP) %s (%.2f%%)`", g.RarityString(), g.Level, g.XP, g.XPCap, g.Name, g.Stats.Value)
}
