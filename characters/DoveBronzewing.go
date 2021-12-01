package chars

var doveBronzewing = CharacterStruct{
	Name:         "Dove Bronzewing",
	Weapon:       "Hallshott",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CRDL",
	Stats: CharacterStatsStruct{
		Health:      172,
		Armor:       13,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Headshot !",
			Speed:      52,
			StunChance: 12,
			Damages:    48,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Quick Slash",
			Speed:      0,
			StunChance: 32,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Crescent Slam",
			Speed:      0,
			StunChance: 15,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Outburst",
			Speed:      12,
			StunChance: 24,
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
			d.Damage = 43
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
