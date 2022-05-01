package chars

var whitleySchnee = CharacterStruct{
	Name:         "Whitley Schnee",
	Weapon:       "Arrogance",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Schnee Family",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       10,
		Damage:      56,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Arrogant Look",
			Speed:      26,
			StunChance: 24,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Cold Greeting",
			Speed:      29,
			StunChance: 28,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "What a pity",
			Speed:      24,
			StunChance: 26,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Mischievous Smile",
			Speed:      25,
			StunChance: 14,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
