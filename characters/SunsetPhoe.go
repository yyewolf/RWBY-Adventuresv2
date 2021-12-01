package chars

var sunsetPhoe = CharacterStruct{
	Name:         "Sunset Phoe",
	Weapon:       "Fire Lily",
	Rarity:       5,
	ImageAuthors: "deboo#9678",
	Category:     "OC Contest 2",
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       12,
		Damage:      51,
		Healing:     20,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Gravity Propel",
			Speed:      9,
			StunChance: 31,
			Damages:    34,
			Heal:       30,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Punch",
			Speed:      16,
			StunChance: 22,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Lily Smash",
			Speed:      11,
			StunChance: 21,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Healing Fire",
			Speed:      25,
			StunChance: 27,
			Damages:    45,
			Heal:       60,
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
