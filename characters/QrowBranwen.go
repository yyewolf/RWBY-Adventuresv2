package chars

var qrowBranwen = CharacterStruct{
	Name:         "Qrow Branwen",
	Weapon:       "Harbinger",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team STRQ (Limited)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      48,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "War Scythe",
			Speed:      18,
			StunChance: 24,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Croak",
			Speed:      24,
			StunChance: 31,
			Damages:    30,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Retract Sword",
			Speed:      22,
			StunChance: 29,
			Damages:    48,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Unlucky Bullets",
			Speed:      29,
			StunChance: 21,
			Damages:    46,
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
