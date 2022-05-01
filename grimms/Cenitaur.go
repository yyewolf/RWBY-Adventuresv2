package grimms

var cenitaur = GrimmStruct{
	Name:         "Cenitaur",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Limited",
	Limited:      true,
	Stats: GrimmStatsStruct{
		Health:      170,
		Armor:       15,
		Damage:      45,
		Healing:     0,
		DodgeChance: 3,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Smash",
			Speed:      29,
			StunChance: 21,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Chant",
			Speed:      24,
			StunChance: 12,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Intimidate",
			Speed:      14,
			StunChance: 31,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rupture",
			Speed:      17,
			StunChance: 23,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SpePriority: 6,
	Special: GrimmSpe{
		Every: 5,
		Main: func(stats *GrimmStatsStruct, semblance *GrimmSpe) (d GrimmSpeUsed) {
			d.Damage = semblance.CustomData["taken"].(int)
			semblance.CustomData["taken"] = 0
			return
		},
		GotAttacked: func(stats *GrimmStatsStruct, dealt int, semblance *GrimmSpe) {
			i := semblance.CustomData["resetIn"].(int) + int(float64(dealt)/1.7)
			semblance.CustomData["taken"] = i
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals the damage you've received during the last time you used this semblance (or since the beginning).",
	},
}
