package models

import (
	"errors"
	"rwby-adventures/config"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Player struct {
	DiscordID     string `gorm:"primary_key;column:discord_id"`
	IsNew         bool   `gorm:"column:is_new;not null"`
	Balance       int64  `gorm:"column:balance;not null"`
	BiddedBalance int64  `gorm:"column:bidded_balance;not null"`
	Level         int64  `gorm:"column:level;not null"`
	CP            int64  `gorm:"column:cp;not null"`
	MaxCP         int64  `gorm:"column:max_cp;not null"`
	CharLimit     int    `gorm:"column:max_char;not null"`
	Maxlootbox    int    `gorm:"column:max_lootbox;not null"`
	SelectedID    string `gorm:"column:selected_id;not null"`
	SelectedType  int    `gorm:"column:selected_type;not null"`
	Badges        int64  `gorm:"column:badges;not null"`
	Settings      int64  `gorm:"column:settings;not null"`
	Disabled      bool   `gorm:"column:disabled;not null"`
	Arms          int    `gorm:"column:arms;not null"`
	Minions       int    `gorm:"column:minions;not null"`
	Jar           int64  `gorm:"column:jar;not null"`

	// Foreign keys
	Missions     *PlayerMission  `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Status       *PlayerStatus   `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Shop         *PlayerShop     `gorm:"foreignkey:DiscordID;references:DiscordID"`
	LastBoxes    *PlayerLootTime `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Gamble       *PlayerGamble   `gorm:"foreignkey:DiscordID;references:DiscordID"`
	Boxes        *PlayerBoxes    `gorm:"foreignkey:DiscordID;references:DiscordID"`
	LimitedBoxes []*LimitedBoxes `gorm:"foreignkey:DiscordID"`
	SpecialBoxes []*SpecialBoxes `gorm:"foreignkey:DiscordID"`

	// Loaded later
	Characters    []*Character  `gorm:"-"`
	Grimms        []*Grimm      `gorm:"-"`
	SelectedChar  *Character    `gorm:"-"`
	SelectedGrimm *Grimm        `gorm:"-"`
	CharInMission *Character    `gorm:"-"`
	GrimmInHunt   *Grimm        `gorm:"-"`
	Market        *PlayerMarket `gorm:"-"`
	TradeSent     int           `gorm:"-"`
	TradeReceived int           `gorm:"-"`
}

func GetPlayer(id string) *Player {
	p := &Player{
		DiscordID: id,
	}
	e := config.Database.
		Joins("Status").
		Joins("Missions").
		Joins("Shop").
		Joins("LastBoxes").
		Joins("Gamble").
		Joins("Boxes").
		Preload("LimitedBoxes").
		Preload("SpecialBoxes").
		Find(p, id)
	if e.Error != nil || e.RowsAffected == 0 {
		p = &Player{
			DiscordID:  id,
			IsNew:      true,
			Level:      1,
			CharLimit:  30,
			Maxlootbox: 3,
			Balance:    500,
			Boxes: &PlayerBoxes{
				DiscordID: id,
				Boxes:     1,
			},
			Status: &PlayerStatus{
				DiscordID: id,
			},
			Missions: &PlayerMission{
				DiscordID: id,
			},
			Shop: &PlayerShop{
				DiscordID: id,
			},
			LastBoxes: &PlayerLootTime{
				DiscordID: id,
			},
			Gamble: &PlayerGamble{
				DiscordID: id,
			},
		}
		return p
	}
	e = config.Database.Find(&Trade{}, "sender_id = ?", p.DiscordID)
	p.TradeSent = int(e.RowsAffected)
	e = config.Database.Find(&Trade{}, "receiver_id = ?", p.DiscordID)
	p.TradeReceived = int(e.RowsAffected)

	config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.Characters, "user_id = ? and not in_mission", p.DiscordID)
	config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.Grimms, "user_id = ? and not in_hunt", p.DiscordID)
	config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.CharInMission, "user_id = ? and in_mission", p.DiscordID)
	config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.GrimmInHunt, "user_id = ? and in_hunt", p.DiscordID)
	if p.SelectedID != "" {
		config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.SelectedChar, "user_id = ? and \"characters\".id = ?", p.DiscordID, p.SelectedID)
		config.Database.Joins("Stats").Order(p.Status.OrderBy).Find(&p.SelectedGrimm, "user_id = ? and \"grimms\".id = ?", p.DiscordID, p.SelectedID)
	}
	if p.SelectedChar.Name == "" {
		p.SelectedChar = nil
	}
	if p.SelectedGrimm.Name == "" {
		p.SelectedGrimm = nil
	}
	if p.CharInMission.Name == "" {
		p.CharInMission = nil
	}
	if p.GrimmInHunt.Name == "" {
		p.GrimmInHunt = nil
	}
	return p
}

func (p *Player) CanDropLootBox() (canHe bool, reset bool) {
	lastTime := time.Unix(p.LastBoxes.Time, 0)
	if p.LastBoxes.Amount < p.Maxlootbox {
		canHe = true
		reset = false
		return
	} else if time.Since(lastTime).Hours() > 24 && p.LastBoxes.Amount == p.Maxlootbox {
		canHe = true
		reset = true
		return
	}
	return
}

func (p *Player) CanGamble() (canHe, reset bool) {
	lastTime := time.Unix(p.Gamble.Time, 0)
	if p.Gamble.Amount < 3 {
		canHe = true
		reset = false
		return
	} else if time.Since(lastTime).Hours() > 24 && p.Gamble.Amount >= 3 {
		canHe = true
		reset = true
		return
	}
	return
}

func (p *Player) CanDungeon() bool {
	return time.Now().Unix()-p.Status.LastDungeon > 18000
}

func (p *Player) MaxChar() int {
	return p.CharLimit + p.Shop.Extensions
}

func (p *Player) TotalBalance() int64 {
	return p.Balance - p.BiddedBalance
}

func (p *Player) Lootbox() (s int) {
	s += p.Boxes.Boxes
	s += p.Boxes.RareBoxes
	s += p.Boxes.GrimmBoxes
	s += p.Boxes.RareGrimmBoxes
	s += len(p.LimitedBoxes)
	s += len(p.SpecialBoxes)
	return
}

func (p *Player) CharAmount() int {
	r := len(p.Characters)
	if p.Missions.IsInMission {
		r += 1
	}
	return r
}

func (p *Player) GrimmAmount() int {
	r := len(p.Grimms)
	if p.Missions.IsInHunt {
		r += 1
	}
	return r
}

func (p *Player) GetLatestPersona() (bool, *Character, *Grimm, int, error) {
	var indexc, indexg int
	c := &Character{}
	g := &Grimm{}
	config.Database.
		Preload("Stats").
		Order("owned_at desc").
		First(c, "user_id=? and not in_mission", p.DiscordID)
	config.Database.
		Preload("Stats").
		Order("owned_at desc").
		First(g, "user_id=? and not in_hunt", p.DiscordID)
	for i, char := range p.Characters {
		if char.CharID == c.CharID {
			indexc = i
			break
		}
	}
	for i, grimm := range p.Grimms {
		if grimm.GrimmID == g.GrimmID {
			indexg = i
			break
		}
	}

	if g.Name == "" && c.Name == "" {
		return false, nil, nil, 0, errors.New("not found")
	}
	if c.OwnedAt.After(g.OwnedAt) {
		return false, c, g, indexc, nil
	}
	return true, c, g, indexg, nil
}

func (p *Player) VerifyChars(s []string) bool {
	for _, checkID := range s {
		found := false
		for _, char := range p.Characters {
			if char.CharID == checkID {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (p *Player) VerifyGrimms(s []string) bool {
	for _, checkID := range s {
		found := false
		for _, char := range p.Grimms {
			if char.GrimmID == checkID {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (p *Player) SendLuckNotice(s *discordgo.Session) {
	if p.Shop.LuckBoostTime != 1 {
		return
	}
	channel, err := s.UserChannelCreate(p.DiscordID)
	if err != nil {
		return
	}
	embed := &discordgo.MessageEmbed{
		Title: "Your Luck Boost is running out !",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pm1.narvii.com/6719/01c7db349c4c882b866c06aeb1e3784c6e0c30fc_hq.jpg",
		},
		Color: config.Botcolor,
	}
	s.ChannelMessageSendEmbed(channel.ID, embed)
}

func (p *Player) AddSelectedXP(multiplier int, boost bool) int64 {
	if p.SelectedChar != nil {
		return p.SelectedChar.AddXP(multiplier, boost)
	}
	if p.SelectedGrimm != nil {
		return p.SelectedGrimm.AddXP(multiplier, boost)
	}
	return 0
}
