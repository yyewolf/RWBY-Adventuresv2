package chars

var yatsuhashiDaichi = CharacterStruct{
	Name:         "Yatsuhashi Daichi",
	Weapon:       "Fulcrum",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CFVY",
	Stats: CharacterStatsStruct{
		Health:      220,
		Armor:       60,
		Damage:      30,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Knock out",
			Speed:      21,
			StunChance: 48,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Sword Slash",
			Speed:      12,
			StunChance: 0,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Savage Slash",
			Speed:      14,
			StunChance: 11,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Silverwing Slash",
			Speed:      52,
			StunChance: 11,
			Damages:    28,
			Heal:       0,
			Every:      2,
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
