package chars

var mayMarigold = CharacterStruct{
	Name:         "May Marigold",
	Weapon:       "Crossbow Staff",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Happy Huntresses",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       20,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Leap Strike",
			Speed:      42,
			StunChance: 33,
			Damages:    49,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Counter Strike",
			Speed:      12,
			StunChance: 20,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Lacerating Shot",
			Speed:      11,
			StunChance: 20,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Darkness Strike",
			Speed:      13,
			StunChance: 23,
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
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
