package chars

var cloverEbi = CharacterStruct{
	Name:         "Clover Ebi",
	Weapon:       "Kingfisher",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Ace Operatives",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Grappling hook",
			Speed:      3,
			StunChance: 40,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fishing Pole",
			Speed:      0,
			StunChance: 35,
			Damages:    28,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Overflowing Wave",
			Speed:      5,
			StunChance: 23,
			Damages:    33,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Harpooned",
			Speed:      12,
			StunChance: 23,
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
			d.Damage = 52
			d.StunChance = 44
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability and has stun chance.",
	},
}
