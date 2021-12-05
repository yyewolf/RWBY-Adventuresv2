package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AddCmd(cmd *Command) {
	if cmd.Type == 0 {
		cmd.Type = discordgo.ChatApplicationCommand
	}
	if !cmd.IsSub {
		cmd.HelpName = cmd.Name
		CommandRouter.Commands = append(CommandRouter.Commands, cmd)
		fmt.Printf("[COMMANDS] Registered top %s\n", cmd.Name)
	} else {
		fmt.Printf("[COMMANDS] Registered sub %s\n", cmd.Name)
	}
	HelpMenus[string(cmd.Menu)] = append(HelpMenus[string(cmd.Menu)], cmd)
	for _, sub := range cmd.SubCommands {
		sub.HelpName = fmt.Sprintf("%s %s", cmd.Name, sub.Name)
		sub.IsSub = true
		AddCmd(sub)
	}
}
