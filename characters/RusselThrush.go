package chars

var russelThrush = CharacterStruct{
	Name:         "Russel Thrush",
	Weapon:       "Shortwings",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CRDL",
	Stats: CharacterStatsStruct{
		Health:      185,
		Armor:       18,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Fast Throw",
			Speed:      31,
			StunChance: 22,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Aura Slicer",
			Speed:      22,
			StunChance: 12,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Quick Slash",
			Speed:      31,
			StunChance: 18,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Blade of Dust",
			Speed:      18,
			StunChance: 31,
			Damages:    31,
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
