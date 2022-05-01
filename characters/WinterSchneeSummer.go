package chars

var winterSchneeSummer = CharacterStruct{
	Name:         "Winter Schnee Summer",
	Weapon:       "Sword",
	Skin:         "Summer",
	Parent:       "Winter Schnee",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.\nいえすぱ",
	Category:     "Schnee Family (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       12,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Frosted Sword",
			Speed:      26,
			StunChance: 24,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Duelist's Dance",
			Speed:      14,
			StunChance: 44,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Frozen Sword Slash",
			Speed:      24,
			StunChance: 26,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Ice Monster",
			Speed:      25,
			StunChance: 14,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
