package chars

var skyLark = CharacterStruct{
	Name:         "Sky Lark",
	Weapon:       "Feather's Edge",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CRDL",
	Stats: CharacterStatsStruct{
		Health:      185,
		Armor:       18,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Gun shot",
			Speed:      32,
			StunChance: 11,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Slow slash",
			Speed:      12,
			StunChance: 31,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Crippling Strike",
			Speed:      21,
			StunChance: 22,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Guillotine",
			Speed:      24,
			StunChance: 41,
			Damages:    28,
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
