package chars

var ivori = CharacterStruct{
	Name:         "Ivori",
	Weapon:       "Whip",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team FNKI",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       15,
		Damage:      43,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Kick",
			Speed:      11,
			StunChance: 62,
			Damages:    41,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Double whip",
			Speed:      21,
			StunChance: 20,
			Damages:    53,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Replica Slash",
			Speed:      31,
			StunChance: 20,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fierce Slap",
			Speed:      42,
			StunChance: 11,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
