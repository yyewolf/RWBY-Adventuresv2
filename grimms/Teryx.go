package grimms

var teryx = GrimmStruct{
	Name:         "Teryx",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      160,
		Armor:       10,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Lunges",
			Speed:      13,
			StunChance: 24,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Claw Strikes",
			Speed:      24,
			StunChance: 26,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Bites",
			Speed:      27,
			StunChance: 19,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Flight",
			Speed:      28,
			StunChance: 16,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
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
