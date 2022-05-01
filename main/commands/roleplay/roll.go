package commmands_roleplay

import (
	"fmt"
	"math"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var RollCommand = &discord.Command{
	Name:        "roll",
	Description: "Rolls a dice.",
	Menu:        discord.RoleplayMenu,
	Args: []discord.Arg{
		{
			Name:        "lower_bound",
			Description: "Lower bound of your dice.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
		{
			Name:        "upper_bound",
			Description: "Upper bound of your dice.",
			Size:        1,
			Type:        discordgo.ApplicationCommandOptionInteger,
		},
	},
	Call: roll,
}

func roll(ctx *discord.CmdContext) {
	lower := ctx.Arguments.GetArg("lower_bound", 0, 1)
	upper := ctx.Arguments.GetArg("upper_bound", 1, 10)

	lF, _ := strconv.ParseFloat(fmt.Sprint(lower.Value), 64)
	uF, _ := strconv.ParseFloat(fmt.Sprint(upper.Value), 64)
	l, u := int(math.Round(lF)), int(math.Round(uF))

	if u <= 0 || l <= 0 {
		ctx.Reply(discord.ReplyParams{
			Content:   "Bounds cannot be negative numbers.",
			Ephemeral: true,
		})
		return
	}

	roll := rand.Intn(u) + l + 1
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s rolled a %d ! (from %d to %d)", ctx.Author.Username, roll, l, u),
			Color: config.Botcolor,
		},
	})
}
