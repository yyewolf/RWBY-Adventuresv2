package chars

var neonKatt = CharacterStruct{
	Name:         "Neon Katt",
	Weapon:       "Nunchaku",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team FNKI",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       13,
		Damage:      50,
		Healing:     0,
		DodgeChance: 7,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rainbow Dash",
			Speed:      30,
			StunChance: 15,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Nunchaku!",
			Speed:      20,
			StunChance: 30,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Wink Wink",
			Speed:      10,
			StunChance: 20,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Never Miss a Beat",
			Speed:      15,
			StunChance: 32,
			Damages:    35,
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
