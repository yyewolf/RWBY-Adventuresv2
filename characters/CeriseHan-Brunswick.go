package chars

var ceriseHanBrunswick = CharacterStruct{
	Name:         "Cerise Han-Brunswick",
	Weapon:       "Strings",
	Rarity:       5,
	ImageAuthors: "mcmystery#0971",
	Category:     "OC Contest 2",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       10,
		Damage:      53,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Repeaters",
			Speed:      25,
			StunChance: 21,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Dancing Whips",
			Speed:      18,
			StunChance: 12,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Summer-Assaults",
			Speed:      19,
			StunChance: 22,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Elemental Attacks",
			Speed:      16,
			StunChance: 14,
			Damages:    46,
			Heal:       0,
			Every:      3,
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
