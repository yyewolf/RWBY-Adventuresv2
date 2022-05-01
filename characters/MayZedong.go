package chars

var mayZedong = CharacterStruct{
	Name:         "May Zedong",
	Weapon:       "Sniper Rifle",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team BRNZ",
	Stats: CharacterStatsStruct{
		Health:      157,
		Armor:       9,
		Damage:      42,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Snipe SHot",
			Speed:      35,
			StunChance: 11,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Slow Aim",
			Speed:      5,
			StunChance: 44,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Hilt slap",
			Speed:      18,
			StunChance: 23,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Slow Aim",
			Speed:      14,
			StunChance: 11,
			Damages:    32,
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
