package chars

var arthurWatts = CharacterStruct{
	Name:         "Arthur Watts",
	Weapon:       "Revolver",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Salem's Inner Circle (Limited)",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      165,
		Armor:       22,
		Damage:      58,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Frenzy Strikes",
			Speed:      54,
			StunChance: 8,
			Damages:    43,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Whirling Shot",
			Speed:      35,
			StunChance: 12,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Blocking Force",
			Speed:      22,
			StunChance: 42,
			Damages:    45,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Headshot",
			Speed:      31,
			StunChance: 12,
			Damages:    49,
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
