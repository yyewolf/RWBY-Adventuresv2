package grimms

var petrageist = GrimmStruct{
	Name:         "Petra Geist",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      190,
		Armor:       35,
		Damage:      38,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Rock Ball",
			Speed:      16,
			StunChance: 26,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Earthquake",
			Speed:      21,
			StunChance: 25,
			Damages:    46,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Destruction",
			Speed:      26,
			StunChance: 14,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Charge",
			Speed:      23,
			StunChance: 31,
			Damages:    26,
			Heal:       0,
			Every:      2,
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
