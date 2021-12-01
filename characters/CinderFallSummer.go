package chars

var cinderFallSummer = CharacterStruct{
	Name:         "Cinder Fall Summer",
	Weapon:       "Fire",
	Skin:         "Summer",
	Parent:       "Cinder Fall",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Cinder's Faction (Limited, Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      182,
		Armor:       15,
		Damage:      55,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Fire Sword Slash 🔥",
			Speed:      12,
			StunChance: 24,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Ball 🔥",
			Speed:      38,
			StunChance: 20,
			Damages:    27,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Spear Throw",
			Speed:      21,
			StunChance: 24,
			Damages:    38,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Knife stab",
			Speed:      24,
			StunChance: 12,
			Damages:    42,
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
