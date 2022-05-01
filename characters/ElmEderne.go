package chars

var elmEderne = CharacterStruct{
	Name:         "Elm Ederne",
	Weapon:       "Timber",
	Rarity:       5,
	ImageAuthors: "Jimmiek Rankin",
	Category:     "Ace Operatives",
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       40,
		Damage:      38,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rockets",
			Speed:      20,
			StunChance: 30,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Hammer Knock",
			Speed:      10,
			StunChance: 35,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Hard Pound",
			Speed:      4,
			StunChance: 33,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Chop chop",
			Speed:      12,
			StunChance: 54,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			stats.Armor += stats.Armor
			semblance.CustomData["resetIn"] = 1
			return
		},
		Passive: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) - 1
			if i >= 0 {
				semblance.CustomData["resetIn"] = i
			}
			if i == 0 {
				stats.Armor -= stats.Armor
			}
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Double your armor for the next turn.",
	},
}
