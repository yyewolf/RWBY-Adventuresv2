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
	if len(c.SubCommands) == 0 && len(c.SubCommandsGroup) == 0 {
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
	} else if len(c.SubCommands) > 0 {
		for _, sub := range c.SubCommands {
			opt := &discordgo.ApplicationCommandOption{
				Name:        sub.Name,
				Description: sub.Description,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options:     sub.makeOption(),
			}
			opts = append(opts, opt)
		}
	} else if len(c.SubCommandsGroup) > 0 {
		for _, sub := range c.SubCommandsGroup {
			opt := &discordgo.ApplicationCommandOption{
				Name:        sub.Name,
				Description: sub.Description,
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Options:     sub.makeOption(),
			}
			opts = append(opts, opt)
		}
	}
	return
}

func (c *Command) Mention() string {
	if c.IsSub {
		return fmt.Sprintf("</%s:%s>", c.HelpName, c.ID)
	}
	return fmt.Sprintf("</%s:%s>", c.Name, c.ID)
}

func (c *Command) PropagateID() {
	for _, sub := range c.SubCommands {
		sub.ID = c.ID
		sub.PropagateID()
	}
	for _, sub := range c.SubCommandsGroup {
		sub.ID = c.ID
		sub.PropagateID()
	}
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
	for _, s := range sessions {
		dcmds, _ := s.ApplicationCommands(config.AppID, "")
		for _, dcmd := range dcmds {
			remove := true
			for _, botcmd := range r.Commands {
				if IsCommandEqual(dcmd, botcmd.make()) {
					// We found the command, we can now mention it
					botcmd.ID = dcmd.ID
					botcmd.PropagateID()
					remove = false
					break
				}
			}
			if remove {
				s.ApplicationCommandDelete(config.AppID, "", dcmd.ID)
				fmt.Printf("Replacing '%s' \n", dcmd.Name)
			}
		}
		for _, botcmd := range r.Commands {
			add := true
			made := botcmd.make()
			for _, dcmd := range dcmds {
				if IsCommandEqual(dcmd, made) {
					add = false
					break
				}
			}
			if add {
				c, err := s.ApplicationCommandCreate(config.AppID, "", made)
				if err != nil {
					fmt.Printf("Cannot create '%s' : %s\n", botcmd.Name, err.Error())
				} else {
					fmt.Printf("Created '%s' \n", botcmd.Name)
					botcmd.ID = c.ID
					botcmd.PropagateID()
				}
			}
		}
	}
	MakeEmbed()
}

func isOptionEqual(c1, c2 *discordgo.ApplicationCommandOption) (b bool) {
	if c1.Name != c2.Name {
		return
	}
	if c1.Description != c2.Description {
		return
	}
	if c1.Type != c2.Type {
		return
	}
	if len(c1.Options) != len(c2.Options) {
		return
	}
	if len(c1.Choices) != len(c2.Choices) {
		return
	}
	if c1.Autocomplete != c2.Autocomplete {
		return
	}
	if c1.Required != c2.Required {
		return
	}
	if len(c1.ChannelTypes) != len(c2.ChannelTypes) {
		return
	}
	for i, chan1 := range c1.ChannelTypes {
		chan2 := c2.ChannelTypes[i]
		if chan1 != chan2 {
			return
		}
	}
	for i, opt1 := range c1.Options {
		opt2 := c2.Options[i]
		if !isOptionEqual(opt1, opt2) {
			return
		}
	}
	return true
}

func IsCommandEqual(c1, c2 *discordgo.ApplicationCommand) (b bool) {
	if c1.Name != c2.Name {
		return
	}
	if c1.Description != c2.Description {
		return
	}
	if c1.Type != c2.Type {
		return
	}
	if len(c1.Options) != len(c2.Options) {
		return
	}
	for i, opt1 := range c1.Options {
		opt2 := c2.Options[i]
		if !isOptionEqual(opt1, opt2) {
			return
		}
	}
	return true
}
