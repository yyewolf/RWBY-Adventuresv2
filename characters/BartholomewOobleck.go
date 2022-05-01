package chars

var bartholomewOobleck = CharacterStruct{
	Name:         "Bartholomew Oobleck",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "Teachers",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       12,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Flamethrower",
			Speed:      12,
			StunChance: 18,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Blast",
			Speed:      28,
			StunChance: 25,
			Damages:    42,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Spear Throw",
			Speed:      14,
			StunChance: 31,
			Damages:    38,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Go Zwei Go!",
			Speed:      25,
			StunChance: 24,
			Damages:    46,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 58
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
