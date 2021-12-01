package chars

var adrianCottaArc = CharacterStruct{
	Name:         "Adrian Cotta-Arc",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Arc",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       10,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Nom",
			Speed:      22,
			StunChance: 19,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Yum",
			Speed:      22,
			StunChance: 18,
			Damages:    46,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Bah",
			Speed:      21,
			StunChance: 25,
			Damages:    42,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Boohoo",
			Speed:      21,
			StunChance: 26,
			Damages:    38,
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
