package chars

var pyrrhaNikosEaster = CharacterStruct{
	Name:         "Pyrrha Nikos Easter",
	Weapon:       "Miló and Akoúo̱",
	Skin:         "Easter",
	Parent:       "Pyrrha Nikos",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team JNPR (Easter)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      200,
		Armor:       28,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Magnetic Slash",
			Speed:      23,
			StunChance: 38,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Snipe hit",
			Speed:      34,
			StunChance: 32,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Xiphos",
			Speed:      23,
			StunChance: 14,
			Damages:    25,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Gloden Javelin",
			Speed:      29,
			StunChance: 23,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = int(float64(semblance.CustomData["taken"].(int)) * 1.0 / 4.0)
			semblance.CustomData["taken"] = 0
			return
		},
		Attacked: func(stats *CharacterStatsStruct, dealt int, semblance *CharacterSemblance) {
			i := semblance.CustomData["taken"].(int) + dealt
			semblance.CustomData["taken"] = i
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals a fourth (1/4) of the damage you've received during the last time you used this semblance (or since the beginning).",
	},
}
