package chars

var foxAlistair = CharacterStruct{
	Name:         "Fox Alistair",
	Weapon:       "Sharp Retribution",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CFVY",
	Stats: CharacterStatsStruct{
		Health:      144,
		Armor:       10,
		Damage:      60,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Flash Slash",
			Speed:      42,
			StunChance: 12,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Axe Kick",
			Speed:      13,
			StunChance: 34,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Cross",
			Speed:      21,
			StunChance: 21,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Jab",
			Speed:      23,
			StunChance: 12,
			Damages:    25,
			Heal:       0,
			Every:      2,
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
