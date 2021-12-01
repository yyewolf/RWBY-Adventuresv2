package chars

var reeseChloris = CharacterStruct{
	Name:         "Reese Chloris",
	Weapon:       "Hoverboard",
	Rarity:       5,
	ImageAuthors: "@mojojoj27827860",
	Category:     "Team ABRN",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Dual Revolvers",
			Speed:      32,
			StunChance: 25,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Bludgeon",
			Speed:      28,
			StunChance: 31,
			Damages:    29,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fiery Blasts",
			Speed:      23,
			StunChance: 22,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Get Frozen!",
			Speed:      17,
			StunChance: 31,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 54
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
