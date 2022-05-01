package chars

var iliaAmitola = CharacterStruct{
	Name:         "Ilia Amitola",
	Weapon:       "Lightning Lash",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "White Fang",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       12,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Spinning Blade",
			Speed:      26,
			StunChance: 23,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Lightning Whip",
			Speed:      14,
			StunChance: 44,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Strike!",
			Speed:      22,
			StunChance: 28,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Thunder Bullets",
			Speed:      21,
			StunChance: 16,
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
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
