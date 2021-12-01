package chars

var flyntCoal = CharacterStruct{
	Name:         "Flynt Coal",
	Weapon:       "Trumpet",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team FNKI",
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       12,
		Damage:      42,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Sound Waves",
			Speed:      63,
			StunChance: 45,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Clones",
			Speed:      0,
			StunChance: 30,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Sound Blast",
			Speed:      12,
			StunChance: 22,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fire wave",
			Speed:      14,
			StunChance: 34,
			Damages:    28,
			Heal:       0,
			Every:      2,
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
