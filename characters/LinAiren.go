package chars

var linAiren = CharacterStruct{
	Name:         "Lin Airen",
	Weapon:       "Strings",
	Rarity:       5,
	ImageAuthors: "Ellennnnnnnnnnnnnnnnnnn (Ellen) #4019",
	Category:     "OC Contest (Team WWZA)",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Dancing breeze",
			Speed:      29,
			StunChance: 25,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Swirling webs",
			Speed:      24,
			StunChance: 18,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fiery revenge",
			Speed:      18,
			StunChance: 19,
			Damages:    50,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Show time",
			Speed:      28,
			StunChance: 31,
			Damages:    27,
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
