package chars

var summerRose = CharacterStruct{
	Name:         "Summer Rose",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team STRQ (Limited)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       12,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rose Petal",
			Speed:      27,
			StunChance: 14,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Silver Eyes",
			Speed:      18,
			StunChance: 24,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Good Bye",
			Speed:      28,
			StunChance: 15,
			Damages:    49,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Ruby's Tears",
			Speed:      16,
			StunChance: 30,
			Damages:    44,
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
