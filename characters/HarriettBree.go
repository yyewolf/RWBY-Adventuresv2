package chars

var harrietBree = CharacterStruct{
	Name:         "Harriet Bree",
	Weapon:       "Fast Knuckles",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Ace Operatives",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       20,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Fast Knuckles",
			Speed:      11,
			StunChance: 23,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Electric Speed",
			Speed:      34,
			StunChance: 32,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Wombo Combo",
			Speed:      56,
			StunChance: 21,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Crab Fists",
			Speed:      34,
			StunChance: 32,
			Damages:    36,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			stats.DodgeChance = 62
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
		Desc:       "Earn more dodge chance for the next turn.",
	},
}
