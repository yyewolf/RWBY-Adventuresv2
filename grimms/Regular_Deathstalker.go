package grimms

var regulardeathstalker = GrimmStruct{
	Name:         "Regular Deathstalker",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      170,
		Armor:       25,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Sting",
			Speed:      18,
			StunChance: 31,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Pinches",
			Speed:      24,
			StunChance: 24,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Deadly Bite",
			Speed:      14,
			StunChance: 24,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Venom",
			Speed:      19,
			StunChance: 28,
			Damages:    41,
			Heal:       0,
			Every:      4,
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
