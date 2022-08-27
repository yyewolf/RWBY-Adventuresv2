package models

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/grimms"
	"rwby-adventures/main/static"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type GrimmStat struct {
	GrimmID     string `gorm:"primary_key;column:id"`
	Health      int    `gorm:"column:health;not null"`
	Armor       int    `gorm:"column:armor;not null"`
	Damage      int    `gorm:"column:damage;not null"`
	Healing     int    `gorm:"column:healing;not null"`
	DodgeChance int    `gorm:"column:dodge_chance;not null"`
}

type Grimm struct {
	GrimmID       string    `gorm:"primary_key;column:id"`
	UserID        string    `gorm:"column:user_id;not null"`
	Name          string    `gorm:"column:name;not null"`
	Level         int       `gorm:"column:level;not null"`
	XP            int64     `gorm:"column:xp;not null"`
	XPCap         int64     `gorm:"column:xp_max;not null"`
	Rarity        int       `gorm:"column:rarity;not null"`
	Value         float64   `gorm:"column:value;not null"`
	InHunt        bool      `gorm:"column:in_hunt;not null"`
	IsInFavorites bool      `gorm:"column:is_in_favorites;not null"`
	Buffs         int       `gorm:"column:buffs;not null"`
	OwnedAt       time.Time `gorm:"column:owned_at;not null"`
	// Foreign keys
	Stats GrimmStat `gorm:"foreignkey:GrimmID"`
}

func GetGrimm(id string) (*Grimm, error) {
	g := &Grimm{
		GrimmID: id,
	}
	e := config.Database.Joins("Stats").Find(g)
	if e.Error != nil || e.RowsAffected == 0 {
		return nil, errors.New("oof")
	}
	return g, nil
}

func (g *Grimm) Valid() bool {
	if g == nil {
		return false
	}
	return g.Name != ""
}

func (g *Grimm) ToRealGrimm() grimms.GrimmStruct {
	i := 0
	for i = range config.BaseGrimms {
		if config.BaseGrimms[i].Name == g.Name {
			break
		}
	}

	returnGrimm := config.BaseGrimms[i]
	returnGrimm.CustomID = g.GrimmID
	returnGrimm.Stats.Armor = g.Stats.Armor
	returnGrimm.Stats.Damage = g.Stats.Damage
	returnGrimm.Stats.Healing = g.Stats.Healing
	returnGrimm.Stats.Health = g.Stats.Health
	returnGrimm.Value = g.Value
	returnGrimm.Rarity = g.Rarity
	returnGrimm.Level = g.Level

	for i := range returnGrimm.Attacks {
		returnGrimm.Attacks[i].LastUsed = -5
	}
	returnGrimm.Special.CustomData = make(map[string]interface{})
	returnGrimm.Special.CustomData["lastTime"] = 0

	return returnGrimm
}

func (g *Grimm) RarityToColor() int {
	EmbedColor := 0
	switch g.Rarity {
	case 0: // Normal
		EmbedColor = 0x808080
	case 1: // Abnormal
		EmbedColor = 0x285300
	case 2: // Sparse
		EmbedColor = 0x00008b
	case 3: // Freaky
		EmbedColor = 0xB22222
	case 4: // Mysterious
		EmbedColor = 0x800080
	case 5: // Bloody
		EmbedColor = 0x121212
	}
	return EmbedColor
}

func (g *Grimm) RarityString() (x string) {
	switch g.Rarity {
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

	for i := 0; i < g.Buffs; i++ {
		x += "+"
	}
	return x
}

func (g *Grimm) FullString() string {
	return fmt.Sprintf("`%s level %d (%d/%dXP) %s (%.2f%%)`", g.RarityString(), g.Level, g.XP, g.XPCap, g.Name, g.Value)
}

func (g *Grimm) MiniString() string {
	return fmt.Sprintf("%s level %d %s (%.2f%%)", g.RarityString(), g.Level, g.Name, g.Value)
}

func (g *Grimm) CheckConditions(f *InvFilters) bool {
	if f.Filtering {
		if f.Name != "" && !(strings.Contains(strings.ToLower(g.Name), strings.ToLower(f.Name))) {
			return false
		}
		if f.Level != 1000 && g.Level != f.Level {
			return false
		}
		if g.Value < f.ValAbove {
			return false
		}
		if g.Value > f.ValBelow {
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

func (g *Grimm) ToLootedEmbed(mention, menuID, box string, original *grimms.GrimmStruct) *discordgo.MessageSend {

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
			Title:       fmt.Sprintf("You looted a %s %s !", g.RarityString(), g.Name),
			Color:       g.RarityToColor(),
			Description: fmt.Sprintf("Congratulations %s, you found %s in a **%s** !\n", mention, g.Name, box),
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

func (g *Grimm) CalcXPCap() int64 {
	return int64(50*g.Level*g.Level + 100)
}

func (g *Grimm) CalcStats() {
	def := g.ToRealGrimm()
	g.Stats.Damage = int(float64(def.Stats.Damage) + float64(9*g.Level)*float64(g.Value/100.0)*math.Pow(2, float64(g.Rarity)/4.6)*math.Pow(3, float64(g.Buffs)/7.0))
	g.Stats.Healing = int(float64(def.Stats.Healing) + float64(11*g.Level)*float64(g.Value/100.0)*math.Pow(2, float64(g.Rarity)/9.0)*math.Pow(3, float64(g.Buffs)/10.0))
	g.Stats.Armor = int(float64(def.Stats.Armor) + float64(8*g.Level)*float64(g.Value/100.0)*math.Pow(2, float64(g.Rarity)/11.8)*math.Pow(3, float64(g.Buffs)/14.0))
	g.Stats.Health = int(float64(def.Stats.Health) + float64(18*g.Level)*float64(g.Value/100.0)*math.Pow(2, float64(g.Rarity)/4.6)*math.Pow(3, float64(g.Buffs)/7.0))
}

func (g *Grimm) CalcXP(multiplier int, boost bool) int64 {
	if g.Level >= 500 {
		return 0
	}
	rand.Seed(time.Now().UTC().UnixNano())
	rint := multiplier * (g.Level)
	add := int64(float64((rand.Intn(26+rint))+15) * (math.Pow(float64(g.Level), 0.72) + 1))
	if boost {
		rint = ((3 / 2) * multiplier) * (g.Level)
		add = int64(float64((rand.Intn(33+rint))+25) * (math.Pow(float64(g.Level), 0.84) + 1))
	}
	return add
}

func (g *Grimm) GiveXP(XP int64) (levelUp bool) {
	for g.XP+XP > g.XPCap {
		levelUp = true
		//if level up
		XP -= g.XPCap - g.XP
		g.Level++
		g.XP = 0
		g.XPCap = g.CalcXPCap()
		g.CalcStats()
	}
	g.XP += XP
	g.XPCap = g.CalcXPCap()
	return levelUp
}

func (g *Grimm) Save() (err error) {
	return config.Database.Save(g).Error
}
