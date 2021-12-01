package chars

var cocoAdelSummer = CharacterStruct{
	Name:         "Coco Adel Summer",
	Weapon:       "Gianduja",
	Skin:         "Summer",
	Parent:       "Coco Adel",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.\nいえすぱ",
	Category:     "Team CFVY (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      162,
		Armor:       12,
		Damage:      60,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Mini-Gun",
			Speed:      22,
			StunChance: 0,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Suitcase slap",
			Speed:      45,
			StunChance: 54,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Aura crush",
			Speed:      34,
			StunChance: 18,
			Damages:    42,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Jaw Slash",
			Speed:      45,
			StunChance: 10,
			Damages:    34,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			semblance.CustomData["oldDamages"] = stats.Damage
			stats.Damage = int(float64(stats.Damage) * float64(1.2))
			semblance.CustomData["resetIn"] = 2
			return
		},
		Passive: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) {
			i := semblance.CustomData["resetIn"].(int) - 1
			if i >= 0 {
				semblance.CustomData["resetIn"] = i
			}
			if i == 0 {
				stats.Damage = semblance.CustomData["oldDamages"].(int)
			}
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Gets a damage buff for the next turn.",
	},
}
