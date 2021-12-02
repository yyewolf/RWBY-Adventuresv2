package grimms

var beowolf = GrimmStruct{
	Name:         "Beowolf",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Lunge",
			Speed:      26,
			StunChance: 23,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Claws Strike",
			Speed:      21,
			StunChance: 31,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Bite",
			Speed:      14,
			StunChance: 24,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Leap",
			Speed:      19,
			StunChance: 28,
			Damages:    45,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
	},
	SpePriority: 1,
	Special: GrimmSpe{
		Every: 4,
		Main: func(stats *GrimmStatsStruct, semblance *GrimmSpe) (d GrimmSpeUsed) {
			d.Damage = 53
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
