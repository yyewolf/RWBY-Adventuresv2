package models

import (
	"bytes"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/grimms"
	"rwby-adventures/main/static"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
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

func (c *Grimm) RarityToColor() int {
	EmbedColor := 0
	switch c.Rarity {
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

func (g *Grimm) MiniString() string {
	return fmt.Sprintf("%s level %d %s (%.2f%%)", g.RarityString(), g.Level, g.Name, g.Stats.Value)
}

func (g *Grimm) CheckConditions(f *InvFilters) bool {
	if f.Filtering {
		if f.Name != "" && !(strings.Contains(strings.ToLower(g.Name), strings.ToLower(f.Name))) {
			return false
		}
		if f.Level != 1000 && g.Level != f.Level {
			return false
		}
		if g.Stats.Value < f.ValAbove {
			return false
		}
		if g.Stats.Value > f.ValBelow {
			return false
		}
		if f.Rarity != -1 && g.Rarity != f.Rarity {
			return false
		}
		if f.Buffs != 0 && g.Buffs != f.Buffs {
			return false
		}
	}
	if f.Favorites && !g.IsInFavorites {
		return false
	}
	return true
}

func (c *Grimm) ToLootedEmbed(mention, menuID, box string, original *grimms.GrimmStruct) *discordgo.MessageSend {

	imgData, _ := static.DatabaseFS.ReadFile(original.ImageFile)
	imgDecoded := bytes.NewBuffer(imgData)

	return &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Reader: imgDecoded,
				Name:   "ch.png",
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("You looted a %s %s !", c.RarityString(), c.Name),
			Color:       c.RarityToColor(),
			Description: fmt.Sprintf("Congratulations %s, you found %s in a **%s** !\n", mention, c.Name, box),
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://ch.png",
			},
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Label: "More info",
						Emoji: discordgo.ComponentEmoji{
							Name: "ℹ️",
						},
						Style:    discordgo.SecondaryButton,
						CustomID: menuID + "-info",
					},
				},
			},
		},
	}
}

func (c *Grimm) CalcXPCap() int64 {
	return int64(50*c.Level*c.Level + 100)
}
