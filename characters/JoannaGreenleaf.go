package chars

var joannaGreenleaf = CharacterStruct{
	Name:         "Joanna Greenleaf",
	Weapon:       "Crossbow Staff",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Happy Huntresses",
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      43,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Last Rites",
			Speed:      45,
			StunChance: 20,
			Damages:    56,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Harrier",
			Speed:      11,
			StunChance: 20,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Whispering Blow",
			Speed:      13,
			StunChance: 12,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Lacerating Shot",
			Speed:      15,
			StunChance: 20,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 62
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
