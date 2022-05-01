package commands_missions

import (
	"rwby-adventures/main/discord"
)

func LoadPassives() {
	discord.RegisterPassiveFunction(PassiveMissions)
}
