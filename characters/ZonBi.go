package chars

var zonBi = CharacterStruct{
	Name:         "Zon Bi",
	Weapon:       "/",
	Rarity:       5,
	ImageAuthors: "• 𝙉𝙎𝘼𝙉𝙄𝙏𝙔 •#7109",
	Category:     "OC Contest (Team WWZA)",
	Limited:      false,
	Stats: CharacterStatsStruct{
		Health:      150,
		Armor:       12,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Jump Rope Rampage",
			Speed:      14,
			StunChance: 15,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Kiss Of Death",
			Speed:      21,
			StunChance: 19,
			Damages:    49,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Whirlwind !",
			Speed:      28,
			StunChance: 31,
			Damages:    35,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Zombie Style Hug",
			Speed:      17,
			StunChance: 28,
			Damages:    47,
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
