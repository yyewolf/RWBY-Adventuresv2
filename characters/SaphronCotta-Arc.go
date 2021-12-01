package chars

var saphronCottaArc = CharacterStruct{
	Name:         "Saphron Cotta-Arc",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Arc",
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
			Name:       "Sunlight",
			Speed:      21,
			StunChance: 28,
			Damages:    44,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Sunrise",
			Speed:      24,
			StunChance: 22,
			Damages:    42,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Solar Flare",
			Speed:      28,
			StunChance: 31,
			Damages:    39,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Eclipse",
			Speed:      19,
			StunChance: 12,
			Damages:    47,
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
