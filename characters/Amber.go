package chars

var amber = CharacterStruct{
	Name:         "Amber",
	Weapon:       "Staff",
	Rarity:       5,
	ImageAuthors: "Amino Autumn Garnet",
	Category:     "Maiden",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      155,
		Armor:       10,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Falling Leaves",
			Speed:      15,
			StunChance: 18,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Reverse shield ",
			Speed:      28,
			StunChance: 27,
			Damages:    38,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Brown Fire",
			Speed:      14,
			StunChance: 21,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Blasts of Wind",
			Speed:      17,
			StunChance: 35,
			Damages:    44,
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
