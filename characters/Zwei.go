package chars

//Zwei is a special character in the game
var Zwei = CharacterStruct{
	Name:         "Zwei",
	Weapon:       "Paws",
	Rarity:       5,
	NotLootable:  true,
	Limited:      true,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Zwei",
	Stats: CharacterStatsStruct{
		Health:      165,
		Armor:       16,
		Damage:      58,
		Healing:     11,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Pat pat",
			Speed:      32,
			StunChance: 20,
			Damages:    35,
			Heal:       5,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Cute pose",
			Speed:      54,
			StunChance: 65,
			Damages:    31,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Distant woofing",
			Speed:      31,
			StunChance: 5,
			Damages:    42,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Munch munch",
			Speed:      16,
			StunChance: 12,
			Damages:    31,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			stats.DodgeChance = 81
			semblance.CustomData["resetIn"] = 1
			return
		},
		Passive: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) - 1
			if i >= 0 {
				semblance.CustomData["resetIn"] = i
			}
			if i == 0 {
				stats.DodgeChance = 0
			}
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Run around a dodge everything.",
	},
}
