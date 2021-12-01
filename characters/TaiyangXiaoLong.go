package chars

var taiyangXiaoLong = CharacterStruct{
	Name:         "Taiyang Xiao Long",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Team STRQ (Limited)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Knuckle Down",
			Speed:      18,
			StunChance: 28,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Charge",
			Speed:      28,
			StunChance: 29,
			Damages:    34,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Smash !",
			Speed:      0,
			StunChance: 17,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "High Kick",
			Speed:      21,
			StunChance: 17,
			Damages:    48,
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
