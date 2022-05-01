package chars

var nadirShiko = CharacterStruct{
	Name:         "Nadir Shiko",
	Weapon:       "Submachine Gun",
	Rarity:       5,
	ImageAuthors: "doodlebugarts",
	Category:     "Team ABRN",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       20,
		Damage:      42,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Submachine Gun",
			Speed:      8,
			StunChance: 30,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Sword Slash",
			Speed:      21,
			StunChance: 49,
			Damages:    30,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Pierce",
			Speed:      12,
			StunChance: 26,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Sundering Shot",
			Speed:      44,
			StunChance: 32,
			Damages:    28,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 49
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
