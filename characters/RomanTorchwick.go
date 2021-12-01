package chars

var romanTorchwick = CharacterStruct{
	Name:         "Roman Torchwick",
	Weapon:       "Melodic Cudgel",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Cinder's Faction (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      190,
		Armor:       18,
		Damage:      49,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Flare Gun",
			Speed:      29,
			StunChance: 24,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Sword Cane",
			Speed:      18,
			StunChance: 23,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Come here!",
			Speed:      22,
			StunChance: 42,
			Damages:    21,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Guess who?",
			Speed:      31,
			StunChance: 12,
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
