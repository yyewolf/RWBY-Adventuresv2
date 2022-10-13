package admin

import "rwby-adventures/main/discord"

func LoadPassives() {
	discord.RegisterPassiveFunction(PassiveUpdate)
}
