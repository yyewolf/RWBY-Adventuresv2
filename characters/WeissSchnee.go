package chars

var weissSchnee = CharacterStruct{
	Name:         "Weiss Schnee",
	Weapon:       "Myrtenaster",
	Rarity:       5,
	ImageAuthors: "kyoukohorimiya\nLavenderRare#3812",
	Category:     "Team RWBY",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       13,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rush-in",
			Speed:      32,
			StunChance: 20,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Parade",
			Speed:      54,
			StunChance: 48,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Pure Strike",
			Speed:      28,
			StunChance: 31,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Steelfang Slash",
			Speed:      12,
			StunChance: 28,
			Damages:    31,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 5,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 61
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
