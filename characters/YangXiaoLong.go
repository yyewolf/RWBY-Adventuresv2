package chars

var yangXiaoLong = CharacterStruct{
	Name:         "Yang Xiao Long",
	Weapon:       "Ember Celica",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.\nLavenderRare#3812",
	Category:     "Team RWBY",
	Stats: CharacterStatsStruct{
		Health:      170,
		Armor:       30,
		Damage:      40,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Skull cracker",
			Speed:      12,
			StunChance: 42,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Upper cut",
			Speed:      38,
			StunChance: 12,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "I'm on fire ! 🔥",
			Speed:      33,
			StunChance: 24,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Low punch",
			Speed:      54,
			StunChance: 12,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 6,
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
