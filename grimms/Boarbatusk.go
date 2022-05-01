package grimms

var boarbatusk = GrimmStruct{
	Name:         "Boarbatusk",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      190,
		Armor:       30,
		Damage:      38,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Tusks",
			Speed:      22,
			StunChance: 26,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Charges",
			Speed:      24,
			StunChance: 31,
			Damages:    23,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Hoof",
			Speed:      18,
			StunChance: 15,
			Damages:    47,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Body Rolls",
			Speed:      23,
			StunChance: 21,
			Damages:    35,
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
