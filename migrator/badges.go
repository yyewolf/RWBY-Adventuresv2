package main

import (
	"strings"

	runner "gopkg.in/mgutz/dat.v2/sqlx-runner"
)

type badge struct {
	value       int64
	name        string
	emoji       string
	description string
	check       func(*player) bool
}

//IF YOU EDIT THIS :
//EDIT ConfigBadges.go FOR THE MARKET AS WELL
var badges = []*badge{
	badgeDeveloper,
	badgeStaff,
	badgeEvent,
	badgeRich,
	badgeLucky,
	badgeArena,
	badgeRP,
	badgePartner,
	badgeBugHunter,
	badgeDuelWin,
	badgeDuelLose,
	badgeDuelPlayed,
	badgeMarketSeller,
	badgeMarketBuyer,
	badgeDungeoneer,
	badgeLootbox,
	badgeRWBY,
	badgeJNPR,
	badgeSTRQ,
	badgeFNKI,
	badgeCFVY,
	badgeNDGO,
	badgeSSSN,
}

var badgeRWBY = &badge{
	value:       1 << 1,
	name:        "RWBY Fan",
	emoji:       "<:Team_Ruby:712297998759100427>",
	description: "Has collected every team RWBY members.",
	check:       characterCheckBuilder("Ruby", "Weiss", "Yang", "Blake"),
}

var badgeJNPR = &badge{
	value:       1 << 2,
	name:        "JNPR Fan",
	emoji:       "<:Team_Jaune_Arc:712298492164571236>",
	description: "Has collected every team JNPR members.",
	check:       characterCheckBuilder("Jaune", "Nora", "Pyrrha", "Ren"),
}

var badgeSTRQ = &badge{
	value:       1 << 3,
	name:        "STRQ Fan",
	emoji:       "<:Team_Raven:831264170363584563>",
	description: "Has collected every team STRQ members.",
	check:       characterCheckBuilder("Summer Rose", "Taiyang", "Raven", "Qrow"),
}

var badgeDeveloper = &badge{
	value:       1 << 4,
	name:        "Developer",
	emoji:       "<:discord_dev:831264750615789578>",
	description: "Is a developer of the bot.",
	check: func(p *player) bool {
		if p.ID == "527218072264769547" {
			return true
		}
		if p.ID == "144472011924570113" {
			return true
		}
		return false
	},
}

var badgeRich = &badge{
	value:       1 << 5,
	name:        "Rich boi",
	emoji:       "<:Lien:745591379173965874>",
	description: "Has reached 1M Liens (a rich boi).",
	check: func(p *player) bool {
		return p.money() >= 1000000
	},
}

var badgeLucky = &badge{
	value:       1 << 6,
	name:        "Lucky boi",
	emoji:       "🍀",
	description: "Has earned a character through gambles.",
	check: func(p *player) bool {
		return false
	},
}

var badgeFNKI = &badge{
	value:       1 << 7,
	name:        "FNKI Fan",
	emoji:       "<:Team_FNKI:712299074795339798>",
	description: "Has collected every team FNKI members.",
	check:       characterCheckBuilder("Flynt Coal", "Neon Katt", "Kobalt", "Ivori"),
}

var badgeCFVY = &badge{
	value:       1 << 8,
	name:        "CFVY Fan",
	emoji:       "<:Team_Coco_Adel:712299075080290315>",
	description: "Has collected every team CFVY members.",
	check:       characterCheckBuilder("Coco Adel", "Fox Alistair", "Velvet Scarlatina", "Yatsuhashi Daichi"),
}

var badgeNDGO = &badge{
	value:       1 << 9,
	name:        "NDGO Fan",
	emoji:       "<:Team_NDGO:712299075025764414>",
	description: "Has collected every team NDGO members.",
	check:       characterCheckBuilder("Nebula Violette", "Dew Gayl", "Gwen Darcy", "Octavia Ember"),
}

var badgeArena = &badge{
	value:       1 << 10,
	name:        "Arena destroyer",
	emoji:       "<:grimm:831537835688984636>",
	description: "Has completed 1000 arenas.",
	check: func(p *player) bool {
		return p.Counter.ArenasCompleted >= 1000
	},
}

var badgeRP = &badge{
	value:       1 << 11,
	name:        "Roleplayer",
	emoji:       "<:stern_kingdom:831539377933910046>",
	description: "Has sent 10000 roleplay messages.",
	check: func(p *player) bool {
		return p.Counter.RoleplaySent >= 10000
	},
}

var badgeDuelWin = &badge{
	value:       1 << 12,
	name:        "Good Duelist",
	emoji:       "<:RWBY_GG:740204287136890960>",
	description: "Has won 250 duels.",
	check: func(p *player) bool {
		return p.Counter.BattlesWon >= 250
	},
}

var badgeDuelLose = &badge{
	value:       1 << 13,
	name:        "Bad Duelist",
	emoji:       "<:Lose:745580183813357599>",
	description: "Has lost 250 duels.",
	check: func(p *player) bool {
		return p.Counter.BattlesLost >= 250
	},
}

var badgeDuelPlayed = &badge{
	value:       1 << 14,
	name:        "Duelist",
	emoji:       "⚔️",
	description: "Has played 500 duels.",
	check: func(p *player) bool {
		return p.Counter.BattlesLost+p.Counter.BattlesWon >= 500
	},
}

var badgeMarketSeller = &badge{
	value:       1 << 15,
	name:        "Massive seller",
	emoji:       "🏪",
	description: "Has sold 500 characters.",
	check: func(p *player) bool {
		return p.Counter.MarketSold >= 500
	},
}

var badgeMarketBuyer = &badge{
	value:       1 << 16,
	name:        "Impulsive buyer",
	emoji:       "🛒",
	description: "Has bought 500 characters.",
	check: func(p *player) bool {
		return p.Counter.MarketBought >= 500
	},
}

var badgeDungeoneer = &badge{
	value:       1 << 17,
	name:        "Dungeon master",
	emoji:       "<:dungeon:831541339510276108>",
	description: "Has played in 750 dungeons.",
	check: func(p *player) bool {
		return p.Counter.DungeonDone >= 750
	},
}

var badgeLootbox = &badge{
	value:       1 << 18,
	name:        "Crate opener",
	emoji:       "<:Rare_Lootbox:745590749709467713>",
	description: "Has opened 2000 loot boxes.",
	check: func(p *player) bool {
		return p.Counter.LootboxOpened >= 2000
	},
}

var badgeEvent = &badge{
	value:       1 << 19,
	name:        "Event guest",
	emoji:       "<:Nora_WoW:745586695247757353>",
	description: "Has participated in an event.",
	check: func(p *player) bool {
		return false
	},
}

var badgeMaxLvl = &badge{
	value:       1 << 20,
	name:        "Good leveler",
	emoji:       "<:pyrrha:745587658230726686>",
	description: "Has 10 or more level 500 characters.",
	check: func(p *player) bool {
		amount := 0
		for _, c := range p.Characters {
			if c.Level == 500 {
				amount++
				if amount == 10 {
					return true
				}
			}
		}
		for _, c := range p.Grimms {
			if c.Level == 500 {
				amount++
				if amount == 10 {
					return true
				}
			}
		}
		return false
	},
}

var badgeInventoryFull = &badge{
	value:       1 << 21,
	name:        "Cat lady",
	emoji:       "<:OH:760102269365780480>",
	description: "Has had an inventory full of a single character.",
	check: func(p *player) bool {
		char := ""
		for _, c := range p.Characters {
			if char == "" {
				char = c.Name
			}
			if char != c.Name {
				return false
			}
		}
		return true
	},
}

var badgeStaff = &badge{
	value:       1 << 22,
	name:        "Ban hammerer",
	emoji:       "🚔",
	description: "Is part of RWBY Adventures' staff.",
	check: func(p *player) bool {
		if p.ID == "144472011924570113" {
			return true
		}
		if p.ID == "527218072264769547" {
			return true
		}
		if p.ID == "304308341151236096" {
			return true
		}
		if p.ID == "279644151006494720" {
			return true
		}
		if p.ID == "341610917345099777" {
			return true
		}
		if p.ID == "714566529517355069" {
			return true
		}
		if p.ID == "457269910549299207" {
			return true
		}
		return false
	},
}

var badgeSSSN = &badge{
	value:       1 << 23,
	name:        "SSSN Fan",
	emoji:       "<:Team_Sun_Wukong:712299074811854910>",
	description: "Has collected every team SSSN members.",
	check:       characterCheckBuilder("Sun Wukong", "Neptune Vasilias", "Scarlet David", "Sage Ayana"),
}

var badgePartner = &badge{
	value:       1 << 24,
	name:        "Partner",
	emoji:       "<:partner_logo:831981258241867808>",
	description: "Is a RWBY Adventures' partner.",
	check: func(p *player) bool {
		return false
	},
}

var badgeBugHunter = &badge{
	value:       1 << 25,
	name:        "Bug Hunter",
	emoji:       "<:bug_hunter:831982297019449394>",
	description: "Has earned this badge by helping the developers team.",
	check: func(p *player) bool {
		return false
	},
}

var badgeOC = &badge{
	value:       1 << 26,
	name:        "OC (WWZA) Fan",
	emoji:       "<:Team_WWZA:832685782245834812>",
	description: "Has collected every team WWZA members.",
	check: func(p *player) bool {
		return false
	},
}

var badgeGrimm = &badge{
	value:       1 << 27,
	name:        "1st Grimm Team Fan",
	emoji:       "<:beowolf:724200490631430205>",
	description: "Has collected every 1st Grimm Team.",
	check: func(p *player) bool {
		return false
	},
}

var badgePackFull = &badge{
	value:       1 << 28,
	name:        "Dog lady",
	emoji:       "<a:Rwby_Triggered:760098648658542592>",
	description: "Has had a pack full of a single grimm.",
	check: func(p *player) bool {
		char := ""
		for _, c := range p.Grimms {
			if char == "" {
				char = c.Name
			}
			if char != c.Name {
				return false
			}
		}
		return true
	},
}

func characterCheckBuilder(charsToCheck ...string) func(p *player) bool {
	return func(p *player) (r bool) {
		return characterCheck(p, charsToCheck)
	}
}

func characterCheck(p *player, charsToCheck []string) (r bool) {
	checkers := make(map[string]bool)
	for _, nameCheck := range charsToCheck {
		checkers[nameCheck] = false
	}
	for _, char := range p.Characters {
		for _, nameCheck := range charsToCheck {
			if strings.Contains(char.Name, nameCheck) && !checkers[nameCheck] {
				checkers[nameCheck] = true
				break
			}
		}
	}
	r = true
	for _, val := range checkers {
		if !val {
			r = false
			break
		}
	}
	return
}

func grimmCheckBuilder(grimmsToCheck ...string) func(p *player) bool {
	return func(p *player) (r bool) {
		return grimmCheck(p, grimmsToCheck)
	}
}

func grimmCheck(p *player, grimmsToCheck []string) (r bool) {
	checkers := make(map[string]bool)
	for _, nameCheck := range grimmsToCheck {
		checkers[nameCheck] = false
	}
	for _, grimm := range p.Grimms {
		for _, nameCheck := range grimmsToCheck {
			if strings.Contains(grimm.Name, nameCheck) && !checkers[nameCheck] {
				checkers[nameCheck] = true
				break
			}
		}
	}
	r = true
	for _, val := range checkers {
		if !val {
			r = false
			break
		}
	}
	return
}

func (p *player) hasBadge(value int64) bool {
	return p.Badges&value == value
}

func (p *player) addBadge(tx *runner.Tx, value int64) {
	if p.hasBadge(value) {
		return
	}
	p.Badges += value
	tx.Update("players").
		Set("badges", p.Badges).
		Where("id=$1", p.ID).
		Exec()
}

func (p *player) removeBadge(tx *runner.Tx, value int64) {
	p.Badges -= value
	tx.Update("players").
		Set("badges", p.Badges).
		Where("id=$1", p.ID).
		Exec()
}
