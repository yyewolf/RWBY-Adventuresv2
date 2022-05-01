package chars

var anRenSummer = CharacterStruct{
	Name:         "An Ren Summer",
	Weapon:       "/",
	Skin:         "Summer",
	Parent:       "An Ren",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Parenting Team (Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       25,
		Damage:      40,
		Healing:     10,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Lotus Attack",
			Speed:      54,
			StunChance: 0,
			Damages:    34,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Blossom Spirit",
			Speed:      13,
			StunChance: 40,
			Damages:    53,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Bamboo hit",
			Speed:      32,
			StunChance: 0,
			Damages:    41,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Retaliate",
			Speed:      18,
			StunChance: 12,
			Damages:    34,
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
