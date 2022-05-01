package chars

var nebulaViolette = CharacterStruct{
	Name:         "Nebula Violette",
	Weapon:       "Crossbow",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team NDGO",
	Stats: CharacterStatsStruct{
		Health:      140,
		Armor:       5,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Lilac Slash",
			Speed:      10,
			StunChance: 30,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Galactic Arrow",
			Speed:      18,
			StunChance: 52,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Pluton",
			Speed:      5,
			StunChance: 20,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Stars shots",
			Speed:      20,
			StunChance: 32,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
