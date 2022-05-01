package chars

var trifa = CharacterStruct{
	Name:         "Trifa",
	Weapon:       "Knife",
	Rarity:       5,
	ImageAuthors: "Zentics",
	Category:     "White Fang",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Cobweb",
			Speed:      21,
			StunChance: 25,
			Damages:    44,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Venomous Bite",
			Speed:      18,
			StunChance: 17,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Cocoon",
			Speed:      20,
			StunChance: 31,
			Damages:    46,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "It's a trap!",
			Speed:      29,
			StunChance: 28,
			Damages:    36,
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
