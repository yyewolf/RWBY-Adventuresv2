package chars

var janinaWolf = CharacterStruct{
	Name:         "Janina Wolf",
	Weapon:       "Emmie and Brandie",
	Rarity:       5,
	ImageAuthors: "GhostlyWeeb#5594\nUltimate Spice, Simp Hunter#9325",
	Category:     "OC Contest (Team WWZA)",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Scythe Trip",
			Speed:      26,
			StunChance: 21,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rip and Tear",
			Speed:      14,
			StunChance: 32,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Run n' Gun",
			Speed:      28,
			StunChance: 18,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Die hard",
			Speed:      21,
			StunChance: 28,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			semblance.CustomData["originalDmg"] = stats.Damage
			stats.Damage = int(float64(1.25) * float64(stats.Damage))
			semblance.CustomData["resetIn"] = 3
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
		Desc:       "Buffs your damages for a few turns.",
	},
}
