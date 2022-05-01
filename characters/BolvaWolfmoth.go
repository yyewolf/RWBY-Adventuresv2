package chars

var bolvaWolfmoth = CharacterStruct{
	Name:         "Bolva Wolfmoth",
	Weapon:       "Forslashe Druhen",
	Rarity:       5,
	ImageAuthors: "月蚀 #9999",
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
			Name:       "Gravity Boost",
			Speed:      12,
			StunChance: 28,
			Damages:    49,
			Heal:       0,
			Every:      43,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Flames",
			Speed:      18,
			StunChance: 22,
			Damages:    29,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Ew~ Frozen !",
			Speed:      21,
			StunChance: 31,
			Damages:    47,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Iron Scratch",
			Speed:      25,
			StunChance: 17,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			semblance.CustomData["originalDmg"] = stats.Damage
			stats.Damage = int(float64(2.5) * float64(stats.Damage))
			semblance.CustomData["resetIn"] = 1
			return
		},
		Passive: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) - 1
			if i >= 0 {
				semblance.CustomData["resetIn"] = i
			}
			if i == 0 {
				stats.Damage = semblance.CustomData["originalDmg"].(int)
			}
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Summons a clone and double your next attack's damage.",
	},
}
