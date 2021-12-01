package chars

var pennyPolendinaXmas = CharacterStruct{
	Name:         "Penny Polendina Xmas",
	Weapon:       "Floating Array",
	Skin:         "Xmas",
	Parent:       "Penny Polendina",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team Penny (Xmas)",
	NotLootable:  true,
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       30,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "10 Floating Blades",
			Speed:      37,
			StunChance: 21,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Gotcha!",
			Speed:      28,
			StunChance: 18,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Puppet Strings",
			Speed:      22,
			StunChance: 36,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Laser Beams",
			Speed:      31,
			StunChance: 24,
			Damages:    37,
			Heal:       0,
			Every:      3,
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
