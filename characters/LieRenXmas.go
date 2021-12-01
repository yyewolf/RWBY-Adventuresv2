package chars

var lieRenXmas = CharacterStruct{
	Name:         "Lie Ren Xmas",
	Weapon:       "StormFlower",
	Skin:         "Xmas",
	Parent:       "Lie Ren",
	Rarity:       5,
	NotLootable:  true,
	Limited:      true,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team JNPR (Xmas)",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       25,
		Damage:      40,
		Healing:     10,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Head Shot",
			Speed:      54,
			StunChance: 0,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Backstab",
			Speed:      13,
			StunChance: 40,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Crush",
			Speed:      32,
			StunChance: 0,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Retaliate",
			Speed:      18,
			StunChance: 12,
			Damages:    34,
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
