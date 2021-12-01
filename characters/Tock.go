package chars

var tock = CharacterStruct{
	Name:         "Tock",
	Weapon:       "Tock's Swords",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Salem's Inner Circle (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      172,
		Armor:       25,
		Damage:      55,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Chronobreak",
			Speed:      12,
			StunChance: 27,
			Damages:    56,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Timewinder",
			Speed:      38,
			StunChance: 32,
			Damages:    29,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Whirling Death",
			Speed:      21,
			StunChance: 24,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rewind",
			Speed:      29,
			StunChance: 12,
			Damages:    48,
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
		Desc:       "Get dodge chance for this and the next turn.",
	},
}
