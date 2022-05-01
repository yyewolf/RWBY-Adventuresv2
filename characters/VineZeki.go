package chars

var vineZeki = CharacterStruct{
	Name:         "Vine Zeki",
	Weapon:       "Thorn",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Ace Operatives",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       15,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Vine Slap",
			Speed:      12,
			StunChance: 50,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Boomerang",
			Speed:      31,
			StunChance: 28,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Lance",
			Speed:      27,
			StunChance: 35,
			Damages:    27,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Genesis",
			Speed:      52,
			StunChance: 12,
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
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
