package chars

var dewGaylSummer = CharacterStruct{
	Name:         "Dew Gayl Summer",
	Weapon:       "Spear",
	Skin:         "Summer",
	Parent:       "Dew Gayl",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team NDGO (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      140,
		Armor:       5,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Dust Whirlwind",
			Speed:      0,
			StunChance: 44,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Spear-Hit",
			Speed:      45,
			StunChance: 38,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Bane Release",
			Speed:      12,
			StunChance: 8,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Angry Jolt",
			Speed:      15,
			StunChance: 8,
			Damages:    25,
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
