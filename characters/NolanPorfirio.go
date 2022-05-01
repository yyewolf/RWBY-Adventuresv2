package chars

var nolanPorfirio = CharacterStruct{
	Name:         "Nolan Porfirio",
	Weapon:       "Cattle Prod",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team BRNZ",
	Stats: CharacterStatsStruct{
		Health:      152,
		Armor:       14,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Quick Slash",
			Speed:      12,
			StunChance: 11,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Slice him up !",
			Speed:      20,
			StunChance: 35,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Electric Shock",
			Speed:      30,
			StunChance: 20,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "U can't run out",
			Speed:      10,
			StunChance: 40,
			Damages:    25,
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
