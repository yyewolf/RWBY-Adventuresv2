package models

import "strings"

var DefaultBadges = []*Badges{
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
	badgeMaxLvl,
	badgeInventoryFull,
	badgePackFull,
	badgeOC,
}

var badgeRWBY = &Badges{
	Name:        "RWBY Fan",
	Emoji:       "<:Team_Ruby:712297998759100427>",
	Description: "Has collected every team RWBY members.",
	Check:       characterCheckBuilder("Ruby", "Weiss", "Yang", "Blake"),
}

var badgeJNPR = &Badges{
	Name:        "JNPR Fan",
	Emoji:       "<:Team_Jaune_Arc:712298492164571236>",
	Description: "Has collected every team JNPR members.",
	Check:       characterCheckBuilder("Jaune", "Nora", "Pyrrha", "Ren"),
}

var badgeSTRQ = &Badges{
	Name:        "STRQ Fan",
	Emoji:       "<:Team_Raven:831264170363584563>",
	Description: "Has collected every team STRQ members.",
	Check:       characterCheckBuilder("Summer Rose", "Taiyang", "Raven", "Qrow"),
}

var badgeDeveloper = &Badges{
	Name:        "Developer",
	Emoji:       "<:discord_dev:831264750615789578>",
	Description: "Is a developer of the bot.",
	Check: func(p *Player) bool {
		if p.DiscordID == "527218072264769547" {
			return true
		}
		if p.DiscordID == "144472011924570113" {
			return true
		}
		return false
	},
}

var badgeRich = &Badges{
	Name:        "Rich boi",
	Emoji:       "<:Lien:745591379173965874>",
	Description: "Has reached 1M Liens (a rich boi).",
	Check: func(p *Player) bool {
		return p.TotalBalance() >= 1000000
	},
}

var badgeLucky = &Badges{
	Name:        "Lucky boi",
	Emoji:       "🍀",
	Description: "Has earned a character through gambles.",
	Check: func(p *Player) bool {
		return false
	},
}

var badgeFNKI = &Badges{
	Name:        "FNKI Fan",
	Emoji:       "<:Team_FNKI:712299074795339798>",
	Description: "Has collected every team FNKI members.",
	Check:       characterCheckBuilder("Flynt Coal", "Neon Katt", "Kobalt", "Ivori"),
}

var badgeCFVY = &Badges{
	Name:        "CFVY Fan",
	Emoji:       "<:Team_Coco_Adel:712299075080290315>",
	Description: "Has collected every team CFVY members.",
	Check:       characterCheckBuilder("Coco Adel", "Fox Alistair", "Velvet Scarlatina", "Yatsuhashi Daichi"),
}

var badgeNDGO = &Badges{
	Name:        "NDGO Fan",
	Emoji:       "<:Team_NDGO:712299075025764414>",
	Description: "Has collected every team NDGO members.",
	Check:       characterCheckBuilder("Nebula Violette", "Dew Gayl", "Gwen Darcy", "Octavia Ember"),
}

var badgeArena = &Badges{
	Name:        "Arena destroyer",
	Emoji:       "<:grimm:831537835688984636>",
	Description: "Has completed 1000 arenas.",
	Check: func(p *Player) bool {
		return p.Stats.ArenasCompleted >= 1000
	},
}

var badgeRP = &Badges{
	Name:        "Roleplayer",
	Emoji:       "<:stern_kingdom:831539377933910046>",
	Description: "Has sent 10000 roleplay messages.",
	Check: func(p *Player) bool {
		return p.Stats.RoleplaySent >= 10000
	},
}

var badgeDuelWin = &Badges{
	Name:        "Good Duelist",
	Emoji:       "<:RWBY_GG:740204287136890960>",
	Description: "Has won 250 duels.",
	Check: func(p *Player) bool {
		return p.Stats.BattlesWon >= 250
	},
}

var badgeDuelLose = &Badges{
	Name:        "Bad Duelist",
	Emoji:       "<:Lose:745580183813357599>",
	Description: "Has lost 250 duels.",
	Check: func(p *Player) bool {
		return p.Stats.BattlesLost >= 250
	},
}

var badgeDuelPlayed = &Badges{
	Name:        "Duelist",
	Emoji:       "⚔️",
	Description: "Has played 500 duels.",
	Check: func(p *Player) bool {
		return p.Stats.BattlesLost+p.Stats.BattlesWon >= 500
	},
}

var badgeMarketSeller = &Badges{
	Name:        "Massive seller",
	Emoji:       "🏪",
	Description: "Has sold 500 characters.",
	Check: func(p *Player) bool {
		return p.Stats.MarketSold >= 500
	},
}

var badgeMarketBuyer = &Badges{
	Name:        "Impulsive buyer",
	Emoji:       "🛒",
	Description: "Has bought 500 characters.",
	Check: func(p *Player) bool {
		return p.Stats.MarketBought >= 500
	},
}

var badgeDungeoneer = &Badges{
	Name:        "Dungeon master",
	Emoji:       "<:dungeon:831541339510276108>",
	Description: "Has played in 750 dungeons.",
	Check: func(p *Player) bool {
		return p.Stats.DungeonDone >= 750
	},
}

var badgeLootbox = &Badges{
	Name:        "Crate opener",
	Emoji:       "<:Rare_Lootbox:745590749709467713>",
	Description: "Has opened 2000 loot boxes.",
	Check: func(p *Player) bool {
		return p.Stats.LootboxOpened >= 2000
	},
}

var badgeEvent = &Badges{
	Name:        "Event guest",
	Emoji:       "<:Nora_WoW:745586695247757353>",
	Description: "Has participated in an event.",
	Check: func(p *Player) bool {
		return false
	},
}

var badgeMaxLvl = &Badges{
	Name:        "Good leveler",
	Emoji:       "<:pyrrha:745587658230726686>",
	Description: "Has 10 or more level 500 characters.",
	Check: func(p *Player) bool {
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

var badgeInventoryFull = &Badges{
	Name:        "Cat lady",
	Emoji:       "<:OH:760102269365780480>",
	Description: "Has had an inventory full of a single character.",
	Check: func(p *Player) bool {
		char := ""
		for _, c := range p.Characters {
			realc := c.ToRealChar()
			if char == "" {
				if realc.Parent != "" {
					char = realc.Parent
				} else {
					char = c.Name
				}
			}
			if char != c.Name && char != realc.Parent {
				return false
			}
		}
		return true
	},
}

var badgeStaff = &Badges{
	Name:        "Ban hammerer",
	Emoji:       "🚔",
	Description: "Is part of RWBY Adventures' staff.",
	Check: func(p *Player) bool {
		if p.DiscordID == "144472011924570113" {
			return true
		}
		if p.DiscordID == "527218072264769547" {
			return true
		}
		if p.DiscordID == "304308341151236096" {
			return true
		}
		if p.DiscordID == "279644151006494720" {
			return true
		}
		if p.DiscordID == "341610917345099777" {
			return true
		}
		if p.DiscordID == "714566529517355069" {
			return true
		}
		if p.DiscordID == "457269910549299207" {
			return true
		}
		return false
	},
}

var badgeSSSN = &Badges{
	Name:        "SSSN Fan",
	Emoji:       "<:Team_Sun_Wukong:712299074811854910>",
	Description: "Has collected every team SSSN members.",
	Check:       characterCheckBuilder("Sun Wukong", "Neptune Vasilias", "Scarlet David", "Sage Ayana"),
}

var badgePartner = &Badges{
	Name:        "Partner",
	Emoji:       "<:partner_logo:831981258241867808>",
	Description: "Is a RWBY Adventures' partner.",
	Check: func(p *Player) bool {
		return false
	},
}

var badgeBugHunter = &Badges{
	Name:        "Bug Hunter",
	Emoji:       "<:bug_hunter:831982297019449394>",
	Description: "Has earned this Badges by helping the developers team.",
	Check: func(p *Player) bool {
		return false
	},
}

var badgeOC = &Badges{
	Name:        "OC (WWZA) Fan",
	Emoji:       "<:Team_WWZA:832685782245834812>",
	Description: "Has collected every team WWZA members.",
	Check:       grimmCheckBuilder("Janina Wolf", "Bolva Wolfmoth", "Zon Bi", "Lin Airen"),
}

var badgePackFull = &Badges{
	Name:        "Dog lady",
	Emoji:       "<a:Rwby_Triggered:760098648658542592>",
	Description: "Has had a pack full of a single grimm.",
	Check: func(p *Player) bool {
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

func characterCheckBuilder(charsToCheck ...string) func(p *Player) bool {
	return func(p *Player) (r bool) {
		return characterCheck(p, charsToCheck)
	}
}

func characterCheck(p *Player, charsToCheck []string) (r bool) {
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

func grimmCheckBuilder(grimmsToCheck ...string) func(p *Player) bool {
	return func(p *Player) (r bool) {
		return grimmCheck(p, grimmsToCheck)
	}
}

func grimmCheck(p *Player, grimmsToCheck []string) (r bool) {
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
