package chars

var rubyRoseSummerV2 = CharacterStruct{
	Name:         "Ruby Rose Summer V2",
	Weapon:       "Crescent Rose",
	Skin:         "SummerV2",
	Parent:       "Ruby Rose",
	Rarity:       5,
	ImageAuthors: "@Hentaking\nいえすぱ",
	Category:     "Team RWBY (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Crescent Dash",
			Speed:      23,
			StunChance: 12,
			Damages:    55,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Sniper Shot",
			Speed:      31,
			StunChance: 22,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Hook",
			Speed:      12,
			StunChance: 18,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Dust Shots",
			Speed:      22,
			StunChance: 32,
			Damages:    33,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 5,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			stats.DodgeChance = 72
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
		Desc:       "Get dodge chance for this and the next turn.",
	},
}
