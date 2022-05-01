package grimms

var queenlancer = GrimmStruct{
	Name:         "Queen Lancer",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Classic",
	Stats: GrimmStatsStruct{
		Health:      150,
		Armor:       12,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []GrimmAttacksStruct{
		{
			Name:       "Firing Stinger",
			Speed:      18,
			StunChance: 16,
			Damages:    48,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Projectiles",
			Speed:      21,
			StunChance: 31,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Ramming",
			Speed:      32,
			StunChance: 12,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Toxik Venom",
			Speed:      19,
			StunChance: 25,
			Damages:    39,
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
