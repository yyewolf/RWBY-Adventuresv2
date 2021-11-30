package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AddCmd(cmd *Command) {
	if cmd.Type == 0 {
		cmd.Type = discordgo.ChatApplicationCommand
	}
	CommandRouter.Commands = append(CommandRouter.Commands, cmd)
	HelpMenus[string(cmd.Menu)] = append(HelpMenus[string(cmd.Menu)], cmd)
	fmt.Printf("[COMMANDS] Registered %s\n", cmd.Name)
}
