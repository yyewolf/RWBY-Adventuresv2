package chars

var adamTaurus = CharacterStruct{
	Name:         "Adam Taurus",
	Weapon:       "Wilt and Blush",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "White Fang (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Fire Dash",
			Speed:      26,
			StunChance: 23,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Blush's Bullet",
			Speed:      21,
			StunChance: 31,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Chokutō Slashu",
			Speed:      14,
			StunChance: 24,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Omae wa mou Shindeiru!",
			Speed:      19,
			StunChance: 28,
			Damages:    45,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 6,
	Semblance: CharacterSemblance{
		Every: 5,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = semblance.CustomData["taken"].(int)
			semblance.CustomData["taken"] = 0
			return
		},
		GotAttacked: func(stats *CharacterStatsStruct, dealt int, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) + int(float64(dealt)/1.7)
			semblance.CustomData["taken"] = i
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals the damage you've received during the last time you used this semblance (or since the beginning).",
	},
}
