package chars

var robynHill = CharacterStruct{
	Name:         "Robyn Hill",
	Weapon:       "Fan Crossbow",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Happy Huntresses",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Crossbow Bolts",
			Speed:      38,
			StunChance: 21,
			Damages:    55,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Vault",
			Speed:      19,
			StunChance: 39,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Eye of Fools",
			Speed:      23,
			StunChance: 24,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Blinding Assault",
			Speed:      15,
			StunChance: 28,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 54
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
