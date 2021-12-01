package chars

var lilMissMalachite = CharacterStruct{
	Name:         "Lil Miss Malachite",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "/",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       15,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "A beer plz!",
			Speed:      19,
			StunChance: 24,
			Damages:    48,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Pipe",
			Speed:      23,
			StunChance: 28,
			Damages:    37,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Bodyguards!",
			Speed:      19,
			StunChance: 22,
			Damages:    45,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Miss Muffet",
			Speed:      22,
			StunChance: 24,
			Damages:    42,
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
