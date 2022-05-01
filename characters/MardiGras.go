package chars

var mardiGras = CharacterStruct{
	Name:         "Mardi Gras",
	Weapon:       "Necklace",
	Rarity:       5,
	ImageAuthors: "@RKD_ART",
	Category:     "OC Contest 2",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       25,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Strike",
			Speed:      15,
			StunChance: 18,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Roguish Guile",
			Speed:      19,
			StunChance: 29,
			Damages:    36,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Dust Blast",
			Speed:      21,
			StunChance: 16,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Party Starter",
			Speed:      14,
			StunChance: 18,
			Damages:    46,
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
