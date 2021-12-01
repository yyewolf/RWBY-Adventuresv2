package chars

var peterPort = CharacterStruct{
	Name:         "Peter Port",
	Weapon:       "Blowhard",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "Teachers",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      200,
		Armor:       20,
		Damage:      41,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Blunderbuss Shot",
			Speed:      15,
			StunChance: 24,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Furious Axe",
			Speed:      24,
			StunChance: 18,
			Damages:    43,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Nice Shot!",
			Speed:      18,
			StunChance: 15,
			Damages:    46,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Undertow",
			Speed:      26,
			StunChance: 28,
			Damages:    40,
			Heal:       0,
			Every:      2,
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
