package chars

var cardinWinchester = CharacterStruct{
	Name:         "Cardin Winchester",
	Weapon:       "The Executioner",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CRDL",
	Stats: CharacterStatsStruct{
		Health:      185,
		Armor:       18,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Explosive Strike",
			Speed:      4,
			StunChance: 45,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Low slash",
			Speed:      20,
			StunChance: 12,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Brutal Crane Flush",
			Speed:      5,
			StunChance: 23,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Rage Binding",
			Speed:      5,
			StunChance: 23,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 4,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = 55
			d.StunChance = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability and can stun.",
	},
}
