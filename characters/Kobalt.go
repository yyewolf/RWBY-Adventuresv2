package chars

var kobalt = CharacterStruct{
	Name:         "Kobalt",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team FNKI",
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      53,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Blue Kick",
			Speed:      22,
			StunChance: 44,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Punch!",
			Speed:      12,
			StunChance: 26,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Blitz",
			Speed:      15,
			StunChance: 12,
			Damages:    43,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Reap",
			Speed:      12,
			StunChance: 33,
			Damages:    27,
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
