package grimms

var nuckelavee = GrimmStruct{
	Name:         "Nuckelavee",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      175,
		Armor:       20,
		Damage:      46,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Extendable Arms",
			Speed:      28,
			StunChance: 14,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Horns Attack",
			Speed:      24,
			StunChance: 32,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Sonic Roar",
			Speed:      12,
			StunChance: 21,
			Damages:    47,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Trample",
			Speed:      23,
			StunChance: 21,
			Damages:    38,
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
