package chars

var blakeBelladonna = CharacterStruct{
	Name:         "Blake Belladonna",
	Weapon:       "Gambol Shroud",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team RWBY",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       25,
		Damage:      50,
		Healing:     10,
		DodgeChance: 7,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Dash attack",
			Speed:      55,
			StunChance: 20,
			Damages:    25,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "From behind",
			Speed:      12,
			StunChance: 22,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Leap Dive",
			Speed:      35,
			StunChance: 28,
			Damages:    43,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Nova Burst",
			Speed:      12,
			StunChance: 33,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			stats.DodgeChance = 72
			d.Heal = 10
			semblance.CustomData["resetIn"] = 1
			return
		},
		Passive: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) - 1
			if i >= 0 {
				semblance.CustomData["resetIn"] = i
			}
			if i == 0 {
				stats.DodgeChance = 0
			}
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Gives you dodge chance for the next turns and heals you during this turn.",
	},
}
