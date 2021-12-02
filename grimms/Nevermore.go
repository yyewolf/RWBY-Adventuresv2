package grimms

var nevermore = GrimmStruct{
	Name:         "Nevermore",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Spearlike Feather Quills",
			Speed:      18,
			StunChance: 12,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Gotcha",
			Speed:      21,
			StunChance: 28,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Beak Attack",
			Speed:      14,
			StunChance: 26,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Headbutt Charge",
			Speed:      22,
			StunChance: 28,
			Damages:    37,
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
