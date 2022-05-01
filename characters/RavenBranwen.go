package chars

var ravenBranwen = CharacterStruct{
	Name:         "Raven Branwen",
	Weapon:       "Omen",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team STRQ (Limited)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Ōdachi",
			Speed:      18,
			StunChance: 19,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Icy Sword !",
			Speed:      22,
			StunChance: 28,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Eyes",
			Speed:      29,
			StunChance: 31,
			Damages:    35,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Slash of Death",
			Speed:      15,
			StunChance: 21,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 58
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
