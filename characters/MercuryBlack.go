package chars

var mercuryBlack = CharacterStruct{
	Name:         "Mercury Black",
	Weapon:       "Talaria",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "Cinder's Faction (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      195,
		Armor:       20,
		Damage:      40,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Roundhouse Kick",
			Speed:      17,
			StunChance: 29,
			Damages:    55,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Flurry",
			Speed:      31,
			StunChance: 19,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Parrrley",
			Speed:      22,
			StunChance: 41,
			Damages:    26,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Cripple",
			Speed:      25,
			StunChance: 24,
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
