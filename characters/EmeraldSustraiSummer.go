package chars

var emeraldSustraiSummer = CharacterStruct{
	Name:         "Emerald Sustrai Summer",
	Weapon:       "Thief's Respite",
	Skin:         "Summer",
	Parent:       "Emerald Sustrai",
	Rarity:       5,
	ImageAuthors: "いえすぱ",
	Category:     "Cinder's Faction (Limited Summer)",
	Limited:      true,
	NotLootable:  true,
	Stats: CharacterStatsStruct{
		Health:      180,
		Armor:       20,
		Damage:      54,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Kusarigama",
			Speed:      21,
			StunChance: 28,
			Damages:    54,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Large Sickles",
			Speed:      29,
			StunChance: 18,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Punch!",
			Speed:      32,
			StunChance: 31,
			Damages:    24,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Bullet Time",
			Speed:      26,
			StunChance: 24,
			Damages:    44,
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
