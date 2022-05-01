package chars

var glyndaGoodwitchSummer = CharacterStruct{
	Name:         "Glynda Goodwitch Summer",
	Weapon:       "The Disciplinarian",
	Skin:         "Summer",
	Parent:       "Glynda Goodwitch",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "Ozpin (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Telekinesis!",
			Speed:      18,
			StunChance: 21,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Weather Manipulation",
			Speed:      26,
			StunChance: 25,
			Damages:    43,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Domination",
			Speed:      24,
			StunChance: 22,
			Damages:    46,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Levitate",
			Speed:      25,
			StunChance: 36,
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
