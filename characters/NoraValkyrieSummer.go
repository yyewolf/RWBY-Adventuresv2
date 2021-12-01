package chars

var noraValkyrieSummer = CharacterStruct{
	Name:         "Nora Valkyrie Summer",
	Weapon:       "Myrtenaster",
	Skin:         "Summer",
	Parent:       "Nora Valkyrie",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team JNPR (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      185,
		Armor:       32,
		Damage:      47,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Head Breaker",
			Speed:      10,
			StunChance: 25,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Stomp!",
			Speed:      30,
			StunChance: 30,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},

		{
			Name:       "Lightning Strike",
			Speed:      10,
			StunChance: 10,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Heart Bomb!❤️",
			Speed:      30,
			StunChance: 20,
			Damages:    28,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 5,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = semblance.CustomData["taken"].(int)
			semblance.CustomData["taken"] = 0
			return
		},
		GotAttacked: func(stats *CharacterStatsStruct, dealt int, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) + dealt
			semblance.CustomData["taken"] = i
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals the damage you've received during the last time you used this semblance (or since the beginning).",
	},
}
