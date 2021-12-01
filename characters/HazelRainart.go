package chars

var hazelRainart = CharacterStruct{
	Name:         "Hazel Rainart",
	Weapon:       "Fists",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Salem's Inner Circle (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      190,
		Armor:       25,
		Damage:      40,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Knuckle Down",
			Speed:      12,
			StunChance: 40,
			Damages:    48,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Haymaker",
			Speed:      36,
			StunChance: 32,
			Damages:    31,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Facebreaker",
			Speed:      23,
			StunChance: 27,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "SMASH",
			Speed:      30,
			StunChance: 10,
			Damages:    43,
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
