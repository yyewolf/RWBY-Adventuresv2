package chars

var jauneArcHalloween = CharacterStruct{
	Name:         "Jaune Arc Halloween",
	Weapon:       "Crocea Mors",
	Skin:         "Halloween",
	Parent:       "Jaune Arc",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team JNPR (Halloween)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      200,
		Armor:       50,
		Damage:      20,
		Healing:     20,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Patch up",
			Speed:      80,
			StunChance: 0,
			Damages:    0,
			Heal:       42,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Crippling Slash",
			Speed:      0,
			StunChance: 23,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rive",
			Speed:      0,
			StunChance: 12,
			Damages:    32,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Riposte",
			Speed:      0,
			StunChance: 23,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 50
			d.Heal = 14
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability and heals you this turn.",
	},
}
