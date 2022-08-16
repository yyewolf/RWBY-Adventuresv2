package models

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"rwby-adventures/main/static"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	CharType = iota
	GrimmType
)

type InvFilters struct {
	Name      string
	Level     int
	ValAbove  float64
	ValBelow  float64
	Buffs     int
	Rarity    int
	Favorites bool

	Filtering bool
}

type CharacterStats struct {
	CharID      string `gorm:"primary_key;column:id"`
	Health      int    `gorm:"column:health;not null"`
	Armor       int    `gorm:"column:armor;not null"`
	Damage      int    `gorm:"column:damage;not null"`
	Healing     int    `gorm:"column:healing;not null"`
	DodgeChance int    `gorm:"column:dodge_chance;not null"`
}

type Character struct {
	CharID        string    `gorm:"primary_key;column:id"`
	UserID        string    `gorm:"column:user_id;not null" json:"-"`
	Name          string    `gorm:"column:name;not null"`
	Level         int       `gorm:"column:level;not null"`
	XP            int64     `gorm:"column:xp;not null"`
	XPCap         int64     `gorm:"column:xp_max;not null"`
	Rarity        int       `gorm:"column:rarity;not null"`
	Value         float64   `gorm:"column:value;not null"`
	InMission     bool      `gorm:"column:in_mission;not null"`
	IsInFavorites bool      `gorm:"column:is_in_favorites;not null"`
	Buffs         int       `gorm:"column:buffs;not null"`
	OwnedAt       time.Time `gorm:"column:owned_at;not null"`
	// Foreign keys
	Stats CharacterStats `gorm:"foreignkey:CharID"`
}

func GetChar(id string) (*Character, error) {
	c := &Character{
		CharID: id,
	}
	e := config.Database.Joins("Stats").Find(c)
	if e.Error != nil || e.RowsAffected == 0 {
		return nil, errors.New("oof")
	}
	return c, nil
}

func (c Character) ToRealChar() chars.CharacterStruct {
	i := -1
	for i = range config.BaseCharacters {
		if config.BaseCharacters[i].Name == c.Name {
			break
		}
	}

	if i == -1 {
		fmt.Println("Error: Character not found")
		return chars.CharacterStruct{}
	}

	returnChar := config.BaseCharacters[i]
	returnChar.CustomID = c.CharID
	returnChar.Stats.Armor = c.Stats.Armor
	returnChar.Stats.Damage = c.Stats.Damage
	returnChar.Stats.Healing = c.Stats.Healing
	returnChar.Stats.Health = c.Stats.Health
	returnChar.Value = c.Value
	returnChar.Rarity = c.Rarity
	returnChar.Level = c.Level

	for i := range returnChar.Attacks {
		returnChar.Attacks[i].LastUsed = -5
	}
	returnChar.Semblance.CustomData = make(map[string]interface{})
	returnChar.Semblance.CustomData["lastTime"] = 0

	return returnChar
}

func (c *Character) RarityToColor() int {
	EmbedColor := 0
	switch c.Rarity {
	case 0: // Common
		EmbedColor = 0x808080
	case 1: // Uncommon
		EmbedColor = 0x7CFC00
	case 2: // Rare
		EmbedColor = 0x87CEEB
	case 3: // Very Rare
		EmbedColor = 0xBA55D3
	case 4: // Legendary
		EmbedColor = 0xFFD700
	case 5: // Collector
		EmbedColor = 0xFF0000
	}
	return EmbedColor
}

func (c *Character) RarityString() (x string) {
	switch c.Rarity {
	case 0: // Common
		x = "□ Common"
	case 1: // Uncommon
		x = "◇ Uncommon"
	case 2: // Rare
		x = "♡ Rare"
	case 3: // Very Rare
		x = "♤ Very Rare"
	case 4: // Legendary
		x = "♧ Legendary"
	case 5: // Collector
		x = "☆ Collector"
	}

	for i := 0; i < c.Buffs; i++ {
		x += "+"
	}
	return x
}

func (c *Character) FullString() string {
	return fmt.Sprintf("%s level %d (%d/%dXP) %s (%.2f%%)", c.RarityString(), c.Level, c.XP, c.XPCap, c.Name, c.Value)
}

func (c *Character) MiniString() string {
	return fmt.Sprintf("%s level %d %s (%.2f%%)", c.RarityString(), c.Level, c.Name, c.Value)
}

func (c *Character) CheckConditions(f *InvFilters) bool {
	if f.Filtering {
		if f.Name != "" && !(strings.Contains(strings.ToLower(c.Name), strings.ToLower(f.Name))) {
			return false
		}
		if f.Level != 1000 && c.Level != f.Level {
			return false
		}
		if c.Value < f.ValAbove {
			return false
		}
		if c.Value > f.ValBelow {
			return false
		}
		if f.Rarity != -1 && c.Rarity != f.Rarity {
			return false
		}
		if f.Buffs != 0 && c.Buffs != f.Buffs {
			return false
		}
	}
	if f.Favorites && !c.IsInFavorites {
		return false
	}
	return true
}

func (c *Character) ToLootedEmbed(mention, menuID, box string, original *chars.CharacterStruct) *discordgo.MessageSend {

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

func (c *Character) CalcXPCap() int64 {
	return int64(50*c.Level*c.Level + 100)
}

func (c *Character) CalcXP(multiplier int, boost bool) int64 {
	if c.Level >= 500 {
		return 0
	}
	rand.Seed(time.Now().UTC().UnixNano())
	rint := multiplier * (c.Level)
	add := int64(float64((rand.Intn(26+rint))+15) * (math.Pow(float64(c.Level), 0.72) + 1))
	if boost {
		rint = ((3 / 2) * multiplier) * (c.Level)
		add = int64(float64((rand.Intn(33+rint))+25) * (math.Pow(float64(c.Level), 0.84) + 1))
	}
	return add
}

func (c *Character) GiveXP(XP int64) (levelUp bool) {
	for c.XP+XP > c.XPCap {
		levelUp = true
		//if level up
		XP -= c.XPCap - c.XP
		c.Level++
		c.XP = 0
		c.XPCap = c.CalcXPCap()
		c.CalcStats()
	}
	c.XP += XP
	c.XPCap = c.CalcXPCap()
	return levelUp
}

func (c *Character) CalcStats() {
	def := c.ToRealChar()
	c.Stats.Damage = int(float64(def.Stats.Damage) + float64(9*c.Level)*float64(c.Value/100.0)*math.Pow(2, float64(c.Rarity)/4.6)*math.Pow(3, float64(c.Buffs)/7.0))
	c.Stats.Healing = int(float64(def.Stats.Healing) + float64(11*c.Level)*float64(c.Value/100.0)*math.Pow(2, float64(c.Rarity)/9.0)*math.Pow(3, float64(c.Buffs)/10.0))
	c.Stats.Armor = int(float64(def.Stats.Armor) + float64(8*c.Level)*float64(c.Value/100.0)*math.Pow(2, float64(c.Rarity)/11.8)*math.Pow(3, float64(c.Buffs)/14.0))
	c.Stats.Health = int(float64(def.Stats.Health) + float64(18*c.Level)*float64(c.Value/100.0)*math.Pow(2, float64(c.Rarity)/4.6)*math.Pow(3, float64(c.Buffs)/7.0))
}

func (c *Character) Save() (err error) {
	return config.Database.Save(c).Error
}
