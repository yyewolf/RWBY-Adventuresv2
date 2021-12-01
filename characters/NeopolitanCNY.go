package chars

var neopolitanChineseNewYear = CharacterStruct{
	Name:         "Neopolitan Chinese New Year",
	Weapon:       "Hush",
	Skin:         "Chinese New Year",
	Parent:       "Neopolitan",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Cinder's Faction (Limited CNY)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      187,
		Armor:       15,
		Damage:      55,
		Healing:     0,
		DodgeChance: 9,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Dolon Cane",
			Speed:      23,
			StunChance: 26,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Extended Hush",
			Speed:      32,
			StunChance: 24,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Guess who?",
			Speed:      21,
			StunChance: 21,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Heavy Kick",
			Speed:      19,
			StunChance: 38,
			Damages:    43,
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
