package chars

var mariaCalavera = CharacterStruct{
	Name:         "Maria Calavera",
	Weapon:       "Life and Death",
	Skin:         "",
	Parent:       "",
	Rarity:       5,
	ImageAuthors: "RWBY Amity Arena",
	Category:     "/",
	Limited:      true,
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       12,
		Damage:      52,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Silver Eyes",
			Speed:      11,
			StunChance: 31,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Purple Bullet",
			Speed:      24,
			StunChance: 21,
			Damages:    31,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Your Death",
			Speed:      21,
			StunChance: 18,
			Damages:    46,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Life Staff",
			Speed:      23,
			StunChance: 12,
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
			d.Damage = 58
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
