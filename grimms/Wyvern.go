package grimms

var wyvern = GrimmStruct{
	Name:         "Wyvern",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      170,
		Armor:       10,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Flight,",
			Speed:      21,
			StunChance: 28,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Spawning Grimm",
			Speed:      23,
			StunChance: 12,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Ramming",
			Speed:      26,
			StunChance: 24,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Bite",
			Speed:      12,
			StunChance: 31,
			Damages:    37,
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
