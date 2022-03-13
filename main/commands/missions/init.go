package commands_missions

import "rwby-adventures/main/discord"

func init() {
	discord.RegisteredPassives = append(discord.RegisteredPassives, PassiveMissions)
}
