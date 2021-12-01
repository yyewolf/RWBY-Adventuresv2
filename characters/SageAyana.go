package chars

var sageAyana = CharacterStruct{
	Name:         "Sage Ayana",
	Weapon:       "Pilgrim",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team SSSN",
	Stats: CharacterStatsStruct{
		Health:      143,
		Armor:       12,
		Damage:      30,
		Healing:     20,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Shock burst",
			Speed:      21,
			StunChance: 32,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Heal up!",
			Speed:      54,
			StunChance: 0,
			Damages:    0,
			Heal:       40,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Spinning Slash",
			Speed:      19,
			StunChance: 26,
			Damages:    31,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Dash in",
			Speed:      32,
			StunChance: 31,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			d.Heal = 12
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability and heals you a little.",
	},
}
