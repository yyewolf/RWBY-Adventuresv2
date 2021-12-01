package chars

var gwenDarcySummer = CharacterStruct{
	Name:         "Gwen Darcy Summer",
	Weapon:       "Throwing Knives",
	Skin:         "Summer",
	Parent:       "Gwen Darcy",
	Rarity:       5,
	ImageAuthors: "kyoukohorimiya\nいえすぱ",
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
			Name:       "Ballet Knives",
			Speed:      42,
			StunChance: 23,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Oops",
			Speed:      12,
			StunChance: 34,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Knife Burst",
			Speed:      10,
			StunChance: 34,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Assault",
			Speed:      18,
			StunChance: 21,
			Damages:    35,
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
