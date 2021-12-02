package discord

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/bwmarrin/discordgo"
)

func choiceConvert(c []*Choice) (r []*discordgo.ApplicationCommandOptionChoice) {
	for _, choice := range c {
		r = append(r, &discordgo.ApplicationCommandOptionChoice{
			Name:  choice.Name,
			Value: choice.Value,
		})
	}
	return
}

func (c *Command) makeOption() (opts []*discordgo.ApplicationCommandOption) {
	if len(c.SubCommands) == 0 {
		for _, arg := range c.Args {
			opt := &discordgo.ApplicationCommandOption{
				Name:        arg.Name,
				Description: arg.Description,
				Type:        arg.Type,
				Required:    arg.Required,
				Choices:     choiceConvert(arg.Choices),
			}
			opts = append(opts, opt)
		}
	} else {
		for _, sub := range c.SubCommands {
			opt := &discordgo.ApplicationCommandOption{
				Name:        sub.Name,
				Description: sub.Description,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options:     sub.makeOption(),
			}
			opts = append(opts, opt)
		}
	}
	return
}

func (c *Command) make() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Type:        c.Type,
		Description: c.Description,
		Options:     c.makeOption(),
	}
}

func (r *router) getSlashCommands() (out []*discordgo.ApplicationCommand) {
	for _, cmd := range r.Commands {
		out = append(out, cmd.make())
	}
	return
}

func (r *router) LoadSlashCommands(sessions []*discordgo.Session) {
	cmds := r.getSlashCommands()
	for _, s := range sessions {
		dcmds, _ := s.ApplicationCommands(config.AppID, "")
		for _, dcmd := range dcmds {
			remove := true
			for _, botcmd := range cmds {
				if botcmd.Name == dcmd.Name && botcmd.Type == dcmd.Type && len(botcmd.Options) == len(dcmd.Options) {
					remove = false
					break
				}
			}
			if remove {
				s.ApplicationCommandDelete(config.AppID, "", dcmd.ID)
				fmt.Printf("Removed '%v' \n", dcmd.Name)

			}
		}
		for _, botcmd := range cmds {
			add := true
			for _, dcmd := range dcmds {
				if botcmd.Name == dcmd.Name && botcmd.Type == dcmd.Type && len(botcmd.Options) == len(dcmd.Options) {
					add = false
					break
				}
			}
			if add {
				_, err := s.ApplicationCommandCreate(config.AppID, "", botcmd)
				if err != nil {
					fmt.Printf("Cannot create '%v' : %v\n", botcmd.Name, err.Error())
				} else {
					fmt.Printf("Created '%v' \n", botcmd.Name)
				}
			}
		}
		break
	}
}
