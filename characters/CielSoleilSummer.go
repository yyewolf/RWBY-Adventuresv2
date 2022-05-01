package chars

var cielSoleilSummer = CharacterStruct{
	Name:         "Ciel Soleil Summer",
	Weapon:       "/",
	Skin:         "Summer",
	Parent:       "Ciel Soleil",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Team Penny (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       12,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Persistent Look",
			Speed:      28,
			StunChance: 28,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Vault Breaker",
			Speed:      16,
			StunChance: 34,
			Damages:    44,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Smash",
			Speed:      22,
			StunChance: 42,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Chic Kick",
			Speed:      21,
			StunChance: 16,
			Damages:    43,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 53
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
