package chars

var youngSalem = CharacterStruct{
	Name:         "Young Salem",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "RoosterTeeth Anim.",
	Category:     "Salem's Inner Circle",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      53,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Revenge",
			Speed:      25,
			StunChance: 32,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Hatred",
			Speed:      21,
			StunChance: 15,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Salem's tears",
			Speed:      29,
			StunChance: 40,
			Damages:    38,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Heartache",
			Speed:      18,
			StunChance: 27,
			Damages:    49,
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
