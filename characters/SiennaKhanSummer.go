package chars

var siennaKhanSummer = CharacterStruct{
	Name:         "Sienna Khan Summer",
	Weapon:       "Cerberus Whip",
	Skin:         "Summer",
	Parent:       "Sienna Khan",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.\nいえすぱ",
	Category:     "White Fang (Limited Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      175,
		Armor:       15,
		Damage:      50,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Throwning Blades",
			Speed:      27,
			StunChance: 21,
			Damages:    39,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Grab",
			Speed:      14,
			StunChance: 34,
			Damages:    35,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Fire Whip",
			Speed:      14,
			StunChance: 24,
			Damages:    40,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Cerberus's Bark",
			Speed:      19,
			StunChance: 28,
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
			d.Damage = 51
			return
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals more damage than a normal ability.",
	},
}
