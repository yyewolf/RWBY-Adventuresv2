package chars

var terraCottaArc = CharacterStruct{
	Name:         "Terra Cotta-Arc",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Arc",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       10,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Moonlight",
			Speed:      21,
			StunChance: 19,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Lunar Rush",
			Speed:      28,
			StunChance: 22,
			Damages:    40,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Moonfall",
			Speed:      26,
			StunChance: 24,
			Damages:    42,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Eclipse",
			Speed:      28,
			StunChance: 23,
			Damages:    45,
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
