package chars

var fennecAlbain = CharacterStruct{
	Name:         "Fennec Albain",
	Weapon:       "Cyclone and Inferno",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "White Fang",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      47,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Wind Tornado",
			Speed:      24,
			StunChance: 31,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Double daggers!",
			Speed:      14,
			StunChance: 24,
			Damages:    43,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Cyclone",
			Speed:      22,
			StunChance: 28,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Stab",
			Speed:      19,
			StunChance: 16,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 48
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
