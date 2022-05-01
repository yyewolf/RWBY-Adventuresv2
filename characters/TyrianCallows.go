package chars

var tyrianCallows = CharacterStruct{
	Name:         "Tyrian Callows",
	Weapon:       "Stinger",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Salem's Inner Circle (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      185,
		Armor:       21,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Quick Dash",
			Speed:      48,
			StunChance: 10,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Sting",
			Speed:      32,
			StunChance: 23,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Surprise !",
			Speed:      22,
			StunChance: 42,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Side Strike",
			Speed:      31,
			StunChance: 12,
			Damages:    44,
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
