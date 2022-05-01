package chars

var neptuneVasiliasHalloween = CharacterStruct{
	Name:         "Neptune Vasilias Halloween",
	Weapon:       "Tri-Hard",
	Skin:         "Halloween",
	Parent:       "Neptune Vasilias",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team SSSN (Halloween)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      155,
		Armor:       10,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Trident Mode",
			Speed:      20,
			StunChance: 25,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Guandao Mode",
			Speed:      15,
			StunChance: 20,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Rifle Mode",
			Speed:      25,
			StunChance: 25,
			Damages:    30,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Compact Mode",
			Speed:      10,
			StunChance: 10,
			Damages:    26,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
