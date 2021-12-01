package chars

var jamesIronwood = CharacterStruct{
	Name:         "James Ironwood",
	Weapon:       "Due Process",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "Headmasters",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      190,
		Armor:       15,
		Damage:      45,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Double Shot",
			Speed:      19,
			StunChance: 26,
			Damages:    47,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Shadow Bullet",
			Speed:      24,
			StunChance: 28,
			Damages:    40,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Iron Arm Hit",
			Speed:      21,
			StunChance: 18,
			Damages:    51,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Light Bullet",
			Speed:      24,
			StunChance: 28,
			Damages:    40,
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
