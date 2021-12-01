package chars

var bolinHori = CharacterStruct{
	Name:         "Bolin Hori",
	Weapon:       "Staff",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team ABRN",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      47,
		Healing:     0,
		DodgeChance: 6,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Staff Hit",
			Speed:      20,
			StunChance: 42,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Curved Blades",
			Speed:      0,
			StunChance: 30,
			Damages:    30,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rock Spear",
			Speed:      30,
			StunChance: 13,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fang Wave",
			Speed:      0,
			StunChance: 24,
			Damages:    22,
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
