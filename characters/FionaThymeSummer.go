package chars

var fionaThymeSummer = CharacterStruct{
	Name:         "Fiona Thyme Summer",
	Weapon:       "Crossbow Staff",
	Skin:         "Summer",
	Parent:       "Fiona Thyme",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Happy Huntresses (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      165,
		Armor:       10,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Staff Hit",
			Speed:      42,
			StunChance: 20,
			Damages:    58,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Bayonet",
			Speed:      23,
			StunChance: 20,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Swift Strike",
			Speed:      30,
			StunChance: 13,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Frenzy",
			Speed:      12,
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
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
