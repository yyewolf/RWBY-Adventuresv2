package chars

var octaviaEmber = CharacterStruct{
	Name:         "Octavia Ember",
	Weapon:       "Dagger",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team NDGO",
	Stats: CharacterStatsStruct{
		Health:      140,
		Armor:       5,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Fire Dust",
			Speed:      15,
			StunChance: 20,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Sand Arch",
			Speed:      28,
			StunChance: 12,
			Damages:    31,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Sweep",
			Speed:      20,
			StunChance: 15,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Blinding Dust",
			Speed:      38,
			StunChance: 35,
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
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
