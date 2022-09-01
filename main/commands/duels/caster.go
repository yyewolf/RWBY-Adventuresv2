package commands_duels

import (
	"fmt"
	"math/rand"
	chars "rwby-adventures/characters"
	"rwby-adventures/grimms"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

// BattlePersonaIn represents an abstract character in battle
type BattlePersonaIn interface {
	IsAlive() bool
	TakeDamage(from BattlePersona, attack int) (damageDealt int, killed bool, stunned bool, healed int)
}

type BattlePersona struct {
	name      string
	imagePath string
	iconURL   string
	winURL    string
	maxHealth int
	rarity    int
	stunned   bool
	stats     *BattleStats
	atks      []*BattleAtks
	priority  int
	semblance *BattleSemblance

	Type  int
	Grimm *models.Grimm
	Char  *models.Character
}

type BattleStats struct {
	PersonaID   string
	Health      int
	Armor       int
	Damage      int
	Healing     int
	DodgeChance int
}

type BattleAtks struct {
	Name       string `json:"-"`
	Speed      int    `json:"-"`
	StunChance int    `json:"-"`
	Damages    int    `json:"-"`
	Heal       int    `json:"-"`
	LastUsed   int    `json:"-"`
	Every      int    `json:"-"`
}

type BattleSemblance struct {
	Main           func(*BattleStats, *BattleSemblance) BattleSemblanceUsed //Called when used
	Attacked       func(*BattleStats, int, *BattleSemblance)                //Called when attacking
	GotAttacked    func(*BattleStats, int, *BattleSemblance)                //Called when attacked
	Passive        func(*BattleStats, *BattleSemblance)                     //Called each turn
	Every          int
	Desc           string
	CustomFunction map[string]interface{}
	CustomData     map[string]interface{}
}

type BattleSemblanceUsed struct {
	Damage     int
	Heal       int
	StunChance int
}

type BattlePlayer struct {
	User         *discordgo.User
	Chose        bool
	Attack       int
	Semblance    bool
	SelectedID   string
	SelectedType int
}

type BattleStruct struct {
	ID         string
	HasStarted bool
	TurnNumber int
	Finished   bool
	ChannelID  string
	PlayersID  map[string]int

	Players       []*BattlePlayer
	Chars         []*BattlePersona
	BattleMessage *discordgo.Message
	StartTime     time.Time

	Original *discord.CmdContext
	WonImg   string
	Won      *BattlePlayer
	Lost     *BattlePlayer
}

func (s *BattlePersona) IsAlive() bool {
	isAlive := s.stats.Health > 0
	return isAlive
}

func (s *BattlePersona) RarityToColor() int {
	if s.Type == models.CharType {
		return s.Char.RarityToColor()
	}
	return s.Grimm.RarityToColor()
}

func (s *BattlePersona) TakeDamage(from *BattlePersona, attackUsed *BattleAtks) (damageDealt int, killed bool, stunned bool, healed int, dodged bool) {
	if !from.stunned {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) <= s.stats.DodgeChance {
			dodged = true
			return
		}
		damageDealt = int((float64(attackUsed.Damages) / float64(100)) * float64(from.stats.Damage))
		damageDealt -= s.stats.Armor - from.stats.Damage
		if damageDealt < 0 {
			damageDealt = 0
		}

		s.stunned = rand.Intn(100) <= attackUsed.StunChance
		stunned = s.stunned

		s.stats.Health -= damageDealt
		BeforeHP := from.stats.Health

		from.stats.Health += int((float64(attackUsed.Heal) / float64(100)) * float64(from.stats.Healing))

		if from.stats.Health > from.maxHealth {
			from.stats.Health = from.maxHealth
		}

		if s.stats.Health <= 0 {
			s.stats.Health = 0
			killed = true
		}

		AfterHP := from.stats.Health
		healed = AfterHP - BeforeHP
		return
	}
	from.stunned = false
	return
}

func (b *BattleStruct) bothChose() bool {
	if b.Players[0].Chose && b.Players[1].Chose {
		return true
	}
	return false
}

func btnToAtk(str string) (i int) {
	i = -1
	switch str {
	case "1":
		i = 0
		return
	case "2":
		i = 1
		return
	case "3":
		i = 2
		return
	case "4":
		i = 3
		return
	case "s":
		i = 100
		return
	}
	return
}

func opponentPlayer(i int) int {
	//if i is not 1, i will be 0, else i will be 1
	b := !(i != 0)
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

/*
	Conversion
*/

func CharToPersona(char *models.Character) (persona *BattlePersona) {
	abs := char.ToRealChar()
	return &BattlePersona{
		name:      char.Name,
		maxHealth: char.Stats.Health,
		imagePath: abs.ImageFile,
		iconURL:   abs.IconURL,
		winURL:    abs.WinURL,
		stunned:   false,
		rarity:    char.Rarity,
		stats:     getBattlesStats(char.Stats),
		atks:      getBattleAtks(abs.Attacks),
		priority:  abs.SemblancePriority,
		semblance: getBattlesSemblance(abs.Semblance),

		Type: models.CharType,
		Char: char,
	}
}

func GrimmToPersona(grimm *models.Grimm) (persona *BattlePersona) {
	abs := grimm.ToRealGrimm()
	return &BattlePersona{
		name:      grimm.Name,
		maxHealth: grimm.Stats.Health,
		imagePath: abs.ImageFile,
		iconURL:   abs.IconURL,
		winURL:    abs.WinURL,
		stunned:   false,
		rarity:    grimm.Rarity,
		stats:     getBattlesStats(grimm.Stats),
		atks:      getBattleAtks(abs.Attacks),
		priority:  abs.SpePriority,
		semblance: getBattlesSemblance(abs.Special),

		Type:  models.GrimmType,
		Grimm: grimm,
	}
}

func getBattleAtks(in interface{}) []*BattleAtks {
	switch v := in.(type) {
	case []chars.CharacterAttacksStruct:
		ob2 := []*BattleAtks{}
		for _, c := range v {
			add := &BattleAtks{
				Name:       c.Name,
				Speed:      c.Speed,
				StunChance: c.StunChance,
				Damages:    c.Damages,
				Heal:       c.Heal,
				LastUsed:   c.LastUsed,
				Every:      c.Every,
			}
			ob2 = append(ob2, add)
		}
		return ob2
	case []grimms.GrimmAttacksStruct:
		ob2 := []*BattleAtks{}
		for _, c := range v {
			add := &BattleAtks{
				Name:       c.Name,
				Speed:      c.Speed,
				StunChance: c.StunChance,
				Damages:    c.Damages,
				Heal:       c.Heal,
				LastUsed:   c.LastUsed,
				Every:      c.Every,
			}
			ob2 = append(ob2, add)
		}
		return ob2
	}
	return []*BattleAtks{}
}

func getBattlesStats(in interface{}) *BattleStats {
	switch v := in.(type) {
	case models.CharacterStats:
		ob2 := &BattleStats{
			PersonaID:   v.CharID,
			Health:      v.Health,
			Armor:       v.Armor,
			Damage:      v.Damage,
			Healing:     v.Healing,
			DodgeChance: v.DodgeChance,
		}
		return ob2
	case models.GrimmStat:
		ob2 := &BattleStats{
			PersonaID:   v.GrimmID,
			Health:      v.Health,
			Armor:       v.Armor,
			Damage:      v.Damage,
			Healing:     v.Healing,
			DodgeChance: v.DodgeChance,
		}
		return ob2
	}
	return &BattleStats{}
}

func getBattlesSemblance(in interface{}) *BattleSemblance {
	switch v := in.(type) {
	case chars.CharacterSemblance:
		newMain := func(a *BattleStats, b *BattleSemblance) BattleSemblanceUsed {
			return BattleSemblanceUsed(v.Main(&chars.CharacterStatsStruct{
				CharID:      a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, &v))
		}
		newAttacked := func(a *BattleStats, b int, c *BattleSemblance) {
			v.Attacked(&chars.CharacterStatsStruct{
				CharID:      a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, b, &v)
		}
		newGotAttacked := func(a *BattleStats, b int, c *BattleSemblance) {
			v.GotAttacked(&chars.CharacterStatsStruct{
				CharID:      a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, b, &v)
		}
		newPassive := func(a *BattleStats, b *BattleSemblance) {
			fmt.Println(a)
			v.Passive(&chars.CharacterStatsStruct{
				CharID:      a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, &v)
		}
		ob2 := &BattleSemblance{
			Main:           newMain,
			Attacked:       newAttacked,
			GotAttacked:    newGotAttacked,
			Passive:        newPassive,
			Every:          v.Every,
			Desc:           v.Desc,
			CustomFunction: v.CustomFunction,
			CustomData:     v.CustomData,
		}
		return ob2
	case grimms.GrimmSpe:
		newMain := func(a *BattleStats, b *BattleSemblance) BattleSemblanceUsed {
			return BattleSemblanceUsed(v.Main(&grimms.GrimmStatsStruct{
				GrimmID:     a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, &v))
		}
		newAttacked := func(a *BattleStats, b int, c *BattleSemblance) {
			v.Attacked(&grimms.GrimmStatsStruct{
				GrimmID:     a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, b, &v)
		}
		newGotAttacked := func(a *BattleStats, b int, c *BattleSemblance) {
			v.GotAttacked(&grimms.GrimmStatsStruct{
				GrimmID:     a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, b, &v)
		}
		newPassive := func(a *BattleStats, b *BattleSemblance) {
			v.Passive(&grimms.GrimmStatsStruct{
				GrimmID:     a.PersonaID,
				Health:      a.Health,
				Armor:       a.Armor,
				Damage:      a.Damage,
				Healing:     a.Healing,
				DodgeChance: a.DodgeChance,
			}, &v)
		}
		ob2 := &BattleSemblance{
			Main:           newMain,
			Attacked:       newAttacked,
			GotAttacked:    newGotAttacked,
			Passive:        newPassive,
			Every:          v.Every,
			Desc:           v.Desc,
			CustomFunction: v.CustomFunction,
			CustomData:     v.CustomData,
		}
		return ob2
	}
	return &BattleSemblance{}
}

/*
	Semblance related
*/

func (s *BattleSemblance) callMain(stats *BattleStats) BattleSemblanceUsed {
	defer func() {
		if r := recover(); r != nil {
			//
		}
	}()
	r := s.Main(stats, s)
	return r
}

func (s *BattleSemblance) callAttacked(stats *BattleStats, damageDealt int) {
	defer func() {
		if r := recover(); r != nil {
			//
		}
	}()
	s.Attacked(stats, damageDealt, s)
}

func (s *BattleSemblance) callGotAttacked(stats *BattleStats, damageReceived int) {
	defer func() {
		if r := recover(); r != nil {
			//
		}
	}()
	s.GotAttacked(stats, damageReceived, s)
}

func (s *BattleSemblance) callPassive(stats *BattleStats) {
	defer func() {
		if r := recover(); r != nil {
			//
		}
	}()
	s.Passive(stats, s)
}

/*

	Buttons

*/

func newDuelComponent(menuID string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "1️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-1",
				},
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "2️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-2",
				},
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "3️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-3",
				},
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "4️⃣",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-4",
				},
				&discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: "🇸",
					},
					Style:    discordgo.SecondaryButton,
					CustomID: menuID + "-s",
				},
			},
		},
	}
}
