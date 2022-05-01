package chars

var melanieMalachite = CharacterStruct{
	Name:         "Melanie Malachite",
	Weapon:       "Bladed Heels",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Junior's Club",
	Stats: CharacterStatsStruct{
		Health:      148,
		Armor:       16,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "High Kick",
			Speed:      50,
			StunChance: 11,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Backflip Slash",
			Speed:      0,
			StunChance: 32,
			Damages:    38,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Back Kick",
			Speed:      0,
			StunChance: 54,
			Damages:    21,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Low Kick",
			Speed:      0,
			StunChance: 44,
			Damages:    35,
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
