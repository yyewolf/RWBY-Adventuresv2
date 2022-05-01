package chars

var ozma = CharacterStruct{
	Name:         "Ozma",
	Weapon:       "Ozma's Staff",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Ozpin",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      47,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Hope",
			Speed:      26,
			StunChance: 22,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Aphrodite",
			Speed:      24,
			StunChance: 18,
			Damages:    49,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Goodbye.",
			Speed:      21,
			StunChance: 24,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Staff Smash",
			Speed:      28,
			StunChance: 35,
			Damages:    40,
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
