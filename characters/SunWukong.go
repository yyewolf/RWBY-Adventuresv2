package chars

var sunWukong = CharacterStruct{
	Name:         "Sun Wukong",
	Weapon:       "Ruyi Bang and Jingu Bang",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team SSSN",
	Stats: CharacterStatsStruct{
		Health:      165,
		Armor:       15,
		Damage:      40,
		Healing:     0,
		DodgeChance: 7,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Crushing Blow",
			Speed:      23,
			StunChance: 32,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Sun Shot",
			Speed:      27,
			StunChance: 28,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Jumping Tail",
			Speed:      35,
			StunChance: 35,
			Damages:    26,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Sunny Staff",
			Speed:      18,
			StunChance: 21,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 62
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
