package commands_gambles

import "math/rand"

type gambleLoot struct {
	name    string
	display string
}

func pickLoots(canLootChar bool, lucky bool) (loots []gambleLoot) {
	var last string
	for i := 0; i < 3; i++ {
		var current gambleLoot
		r := rand.Float64() * 100
		if last == "arm" {
			r += 3
		}
		if r > 88 || (r > 82 && lucky) && current.name == "" {
			current.name = "arm"
			current.display = "Arm"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleArm++
			// } else {
			// 	oGambleArm++
			// }
		}
		r = rand.Float64() * 100
		if last == "char" {
			r += 3
		}
		if (r > 90 || (r > 84 && lucky)) && canLootChar && current.name == "" {
			current.name = "char"
			current.display = "Character"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleChar++
			// } else {
			// 	oGambleChar++
			// }
		}
		r = rand.Float64() * 100
		if last == "rarelootbox" {
			r += 3
		}
		if r > 86 || (r > 78 && lucky) && current.name == "" {
			current.name = "rarelootbox"
			current.display = "Rare Loot Box"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleRareLootbox++
			// } else {
			// 	oGambleRareLootbox++
			// }
		}
		r = rand.Float64() * 100
		if last == "lootbox" {
			r += 3
		}
		if r > 70 || (r > 62 && lucky) && current.name == "" {
			current.name = "lootbox"
			current.display = "Loot Box"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleLootbox++
			// } else {
			// 	oGambleLootbox++
			// }
		}
		r = rand.Float64() * 100
		if last == "money" {
			r += 3
		}
		if r > 60 || (r > 56 && lucky) && current.name == "" {
			current.name = "money"
			current.display = "Money"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleMoney++
			// } else {
			// 	oGambleMoney++
			// }
		}
		r = rand.Float64() * 100
		if r > 50 || (r > 40 && lucky) && current.name == "" {
			current.name = "xp"
			current.display = "Experience"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleXP++
			// } else {
			// 	oGambleXP++
			// }
		}
		last = current.name
		if current.name == "" {
			current.name = "lose"
			current.display = "Nothing"
			// /*
			// 	STATS
			// */
			// if lucky {
			// 	iGambleNothing++
			// } else {
			// 	oGambleNothing++
			// }
		}
		loots = append(loots, current)
	}
	return
}
