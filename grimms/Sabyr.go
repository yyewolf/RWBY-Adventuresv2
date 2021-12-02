package grimms

var sabyr = GrimmStruct{
	Name:         "Sabyr",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      160,
		Armor:       12,
		Damage:      51,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Biting",
			Speed:      29,
			StunChance: 21,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Clawing",
			Speed:      24,
			StunChance: 12,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Charging",
			Speed:      14,
			StunChance: 31,
			Damages:    37,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Howl",
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
