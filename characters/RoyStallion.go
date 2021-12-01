package chars

var royStallion = CharacterStruct{
	Name:         "Roy Stallion",
	Weapon:       "Saws",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team BRNZ",
	Stats: CharacterStatsStruct{
		Health:      149,
		Armor:       20,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Quick Saw Dash",
			Speed:      31,
			StunChance: 11,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Saw throw",
			Speed:      22,
			StunChance: 25,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Ruthless Predator",
			Speed:      31,
			StunChance: 23,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Cull the Meek",
			Speed:      22,
			StunChance: 31,
			Damages:    32,
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
