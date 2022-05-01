package chars

var brawnzNi = CharacterStruct{
	Name:         "Brawnz Ni",
	Weapon:       "Claws",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team BRNZ",
	Stats: CharacterStatsStruct{
		Health:      157,
		Armor:       13,
		Damage:      42,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Quick Dash",
			Speed:      50,
			StunChance: 11,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Claw Strike",
			Speed:      0,
			StunChance: 32,
			Damages:    38,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Claw Eruption",
			Speed:      0,
			StunChance: 54,
			Damages:    21,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Leaf Stomp",
			Speed:      0,
			StunChance: 44,
			Damages:    35,
			Heal:       0,
			Every:      3,
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
