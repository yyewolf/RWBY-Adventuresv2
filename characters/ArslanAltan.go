package chars

var arslanAltan = CharacterStruct{
	Name:         "Arslan Altan",
	Weapon:       "Dagger",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team ABRN",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      47,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rope Dart",
			Speed:      0,
			StunChance: 44,
			Damages:    29,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Dagger Attack",
			Speed:      0,
			StunChance: 32,
			Damages:    49,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Backstab",
			Speed:      0,
			StunChance: 10,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Launching Stab",
			Speed:      0,
			StunChance: 19,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 48
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
