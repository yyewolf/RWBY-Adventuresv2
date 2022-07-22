package discord

import (
	"rwby-adventures/models"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CmdContext struct {
	Session *discordgo.Session

	ID        string
	ChannelID string
	GuildID   string

	Arguments Args

	Author        *discordgo.User
	Player        *models.Player
	Guild         *models.Guild
	IsInteraction bool
	IsComponent   bool
	Menu          *Menus

	ComponentData discordgo.MessageComponentInteractionData
	ModalData     discordgo.ModalSubmitInteractionData
	Interaction   *discordgo.Interaction
	Message       *discordgo.Message
}

type Choice struct {
	Name  string
	Value interface{}
}
type CmdAlias []string

func (a CmdAlias) Has(alias string) bool {
	for _, str := range a {
		if str == alias {
			return true
		}
	}
	return false
}

type Command struct {
	Name        string
	Description string
	Type        discordgo.ApplicationCommandType
	Aliases     CmdAlias
	Menu        menuName

	Args        []Arg
	SubCommands []*Command

	// Not needed when registering a command
	IsSub    bool
	HelpName string

	Call func(*CmdContext)
}

type router struct {
	Prefix         string
	ListenerPrefix string
	//RateLimit in milliseconds
	RateLimit int

	Commands []*Command
}

type Menus struct {
	MenuID        string
	Modal         bool
	Source        *discordgo.MessageEmbed
	SourceContext *CmdContext
	Data          interface{}

	Call func(*CmdContext)
}

func (r *router) findTopCommand(name string) *Command {
	for _, cmd := range r.Commands {
		if cmd.Name == name || cmd.Aliases.Has(name) {
			return cmd
		}
	}
	return nil
}

func (c *Command) findDeepestLink(args []string) (*Command, []string) {
	if len(c.SubCommands) == 0 {
		return c, args
	} else {
		if len(args) == 0 {
			return c, args
		}
		for _, sub := range c.SubCommands {
			if args[0] == sub.Name {
				test, args := sub.findDeepestLink(args[1:])
				if test != nil {
					return test, args
				}
			}
		}
		return nil, args
	}
}

func slicer(data *discordgo.ApplicationCommandInteractionDataOption, args []string) ([]string, []*discordgo.ApplicationCommandInteractionDataOption) {
	args = append(args, data.Name)
	if len(data.Options) == 0 {
		return args, []*discordgo.ApplicationCommandInteractionDataOption{}
	}
	if len(data.Options) > 1 {
		return args, data.Options
	}
	if data.Options[0].Type != discordgo.ApplicationCommandOptionSubCommand {
		return args, data.Options
	}
	return slicer(data.Options[0], args)
}

func interactionToSlice(data *discordgo.ApplicationCommandInteractionData) ([]string, []*discordgo.ApplicationCommandInteractionDataOption) {
	args := []string{data.Name}
	if len(data.Options) == 0 {
		return args, []*discordgo.ApplicationCommandInteractionDataOption{}
	}
	if len(data.Options) > 1 {
		return args, data.Options
	}
	if data.Options[0].Type != discordgo.ApplicationCommandOptionSubCommand {
		return args, data.Options
	}
	return slicer(data.Options[0], args)
}

func routeMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	/*
		Create context :
	*/
	ctx := &CmdContext{
		Session:   s,
		ID:        m.ID,
		GuildID:   m.GuildID,
		ChannelID: m.ChannelID,
		Author:    m.Author,
		Message:   m.Message,
	}
	ctx.Player = models.GetPlayer(ctx.Author.ID)
	ctx.Guild = models.GetGuild(ctx.GuildID)
	/*
		Is Command or for Listener :
	*/
	if !strings.HasPrefix(m.Content, CommandRouter.Prefix) {
		if !strings.HasPrefix(m.Content, CommandRouter.ListenerPrefix) {
			for _, callPassive := range RegisteredPassives {
				callPassive(ctx)
			}
			return
		}
		/*
			Is for Listener :
		*/
		data, _, found, callback := getDataFromCache(m.ChannelID)
		if !found {
			return
		}
		/*
			Create context :
		*/
		ctx := &ListenerContext{
			s:         s,
			ID:        m.ID,
			GuildID:   m.GuildID,
			ChannelID: m.ChannelID,
			Author:    m.Author,
			Message:   m.Message,
			Data:      data,
		}
		ctx.Player = models.GetPlayer(ctx.Author.ID)
		ctx.Guild = models.GetGuild(ctx.GuildID)
		callback(ctx)
		return
	}

	/*
		Rate limits :
	*/
	rateLimited := checkUser(ctx.Author.ID)
	if rateLimited {
		s.MessageReactionAdd(ctx.ChannelID, ctx.ID, "⌛")
		return
	}
	/*
		Find command & args :
	*/
	m.Content = strings.TrimSpace(m.Content)
	m.Content = strings.TrimPrefix(m.Content, CommandRouter.Prefix)
	splt := strings.Split(m.Content, " ")
	if len(splt) == 0 {
		return
	}
	topCmd := splt[0]
	cmd := CommandRouter.findTopCommand(topCmd)
	if cmd == nil {
		return
	}
	deepestLink, argsLeft := cmd.findDeepestLink(splt[1:])

	var realArgs []*CommandArg
	for _, cmdArg := range deepestLink.Args {
		i := 0
		if i >= len(argsLeft) {
			break
		}
		if cmdArg.Size > 1 {
			current := &CommandArg{
				Name:  cmdArg.Name,
				Value: "",
			}
			for j := i; j < i+cmdArg.Size; j++ {
				if j != i+cmdArg.Size-1 {
					current.Value = current.Value.(string) + argsLeft[j] + " "
				} else {
					current.Value = current.Value.(string) + argsLeft[j]
				}
			}
			realArgs = append(realArgs, current)
		} else {
			current := &CommandArg{
				Name:  cmdArg.Name,
				Value: argsLeft[i],
			}
			realArgs = append(realArgs, current)
			i++
		}
	}

	ctx.Arguments = realArgs

	if ctx.Player.IsNew {
		HandleNewPlayer(ctx)
		return
	}

	if deepestLink.Call == nil {
		return
	}
	deepestLink.Call(ctx)
}

func routeInteraction(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	/*
		Verify type :
	*/
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}
	/*
		Create context :
	*/
	ctx := &CmdContext{
		Session:   s,
		ID:        interaction.ID,
		GuildID:   interaction.GuildID,
		ChannelID: interaction.ChannelID,

		Interaction:   interaction.Interaction,
		IsInteraction: true,
	}

	if interaction.Member != nil {
		ctx.Author = interaction.Member.User
	} else {
		ctx.Author = interaction.User
	}

	ctx.Player = models.GetPlayer(ctx.Author.ID)
	ctx.Guild = models.GetGuild(ctx.GuildID)

	/*
		Find command & args :
	*/
	data := interaction.ApplicationCommandData()

	splt, parsedArgs := interactionToSlice(&data)
	if len(splt) == 0 {
		return
	}
	topCmd := splt[0]
	cmd := CommandRouter.findTopCommand(topCmd)
	if cmd == nil {
		return
	}
	deepestLink, _ := cmd.findDeepestLink(splt[1:])
	var realArgs []*CommandArg
	for _, arg := range parsedArgs {
		for _, cmdArg := range deepestLink.Args {
			if arg.Name == cmdArg.Name {
				realArgs = append(realArgs, &CommandArg{
					Name:  arg.Name,
					Value: arg.Value,
				})
			}
		}
	}
	ctx.Arguments = realArgs

	/*
		Not working because you need to know whether you want to set the flag or not and you can't know that before you know you need it (aka in the command func). Thanks Discord.

		s.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: 1 << 6,
			},
		})
	*/

	if ctx.Player.IsNew {
		HandleNewPlayer(ctx)
		return
	}
	if deepestLink.Call == nil {
		return
	}
	deepestLink.Call(ctx)
}

func (c *CmdContext) HandleError() {
	c.Reply(ReplyParams{
		Content:   "There has been an error, please contact the support if this happens again.",
		Ephemeral: true,
	})
}

func (l *ListenerContext) HandleError() {
	l.reply(ReplyParams{
		Content:   "There has been an error, please contact the support if this happens again.",
		Ephemeral: true,
	})
}
