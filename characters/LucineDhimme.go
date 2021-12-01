package chars

var lucineDhimme = CharacterStruct{
	Name:         "Lucine Dhimme",
	Weapon:       "Dual Hookswords",
	Rarity:       5,
	ImageAuthors: "Lacinyc#0912",
	Category:     "OC Contest 2",
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Slash",
			Speed:      18,
			StunChance: 16,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Dash",
			Speed:      21,
			StunChance: 29,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Don't crash",
			Speed:      23,
			StunChance: 12,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Lash",
			Speed:      14,
			StunChance: 18,
			Damages:    44,
			Heal:       0,
			Every:      3,
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
