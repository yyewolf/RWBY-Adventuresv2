package chars

var higanbanaWaitress = CharacterStruct{
	Name:         "Higanbana Waitress",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "ExtremeSoda",
	Category:     "Random",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       12,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Slap!",
			Speed:      19,
			StunChance: 18,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Want a drink?",
			Speed:      21,
			StunChance: 26,
			Damages:    42,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Kick",
			Speed:      25,
			StunChance: 24,
			Damages:    45,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "OwO",
			Speed:      29,
			StunChance: 31,
			Damages:    39,
			Heal:       0,
			Every:      2,
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
