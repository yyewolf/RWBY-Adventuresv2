package chars

var marrowAmin = CharacterStruct{
	Name:         "Marrow Amin",
	Weapon:       "Fetch",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Ace Operatives",
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Rifle Shot",
			Speed:      32,
			StunChance: 30,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Quick Throw",
			Speed:      28,
			StunChance: 44,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Come back",
			Speed:      21,
			StunChance: 32,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Pierce",
			Speed:      26,
			StunChance: 34,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 56
			d.StunChance = 75
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability and can stun.",
	},
}
