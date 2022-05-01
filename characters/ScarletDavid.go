package chars

var scarletDavid = CharacterStruct{
	Name:         "Scarlet David",
	Weapon:       "Hook and Darling",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team SSSN",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       10,
		Damage:      44,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Grapple Hook",
			Speed:      14,
			StunChance: 31,
			Damages:    25,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Darling!",
			Speed:      21,
			StunChance: 17,
			Damages:    56,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fencing",
			Speed:      31,
			StunChance: 23,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Cutlass",
			Speed:      12,
			StunChance: 36,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
