package chars

var velvetScarlatina = CharacterStruct{
	Name:         "Velvet Scarlatina",
	Weapon:       "Anesidora",
	Rarity:       5,
	ImageAuthors: "Rooster Teeth Anim.",
	Category:     "Team CFVY",
	Stats: CharacterStatsStruct{
		Health:      160,
		Armor:       12,
		Damage:      51,
		Healing:     0,
		DodgeChance: 5,
	},
	Attacks: []CharacterAttacksStruct{
		{
			Name:       "Mini-Gun",
			Speed:      32,
			StunChance: 17,
			Damages:    52,
			Heal:       0,
			Every:      4,
			LastUsed:   -5,
		},
		{
			Name:       "Scythe Slash",
			Speed:      21,
			StunChance: 23,
			Damages:    42,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
		{
			Name:       "Wink Wink",
			Speed:      19,
			StunChance: 35,
			Damages:    26,
			Heal:       0,
			Every:      2,
			LastUsed:   -5,
		},
		{
			Name:       "Stomp!",
			Speed:      21,
			StunChance: 29,
			Damages:    36,
			Heal:       0,
			Every:      3,
			LastUsed:   -5,
		},
	},
	SemblancePriority: 1,
	Semblance: CharacterSemblance{
		Every: 5,
		Main: func(stats *CharacterStatsStruct, semblance *CharacterSemblance) (d CharacterSemblanceUsed) {
			d.Damage = semblance.CustomData["taken"].(int)
			semblance.CustomData["taken"] = 0
			return
		},
		GotAttacked: func(stats *CharacterStatsStruct, dealt int, semblance *CharacterSemblance) {
			semblance.CustomData["taken"] = dealt
		},
		CustomData: make(map[string]interface{}),
		Desc:       "Deals the same amount of damage than the last attack you took.",
	},
}
