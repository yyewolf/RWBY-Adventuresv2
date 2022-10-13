package commands_settings

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var OrderCommand = &discord.Command{
	Name:        "order",
	Description: "Choose how you order your inventory.",
	Menu:        discord.ConfigurationMenu,
	Call:        order,
}

func order(ctx *discord.CmdContext) {
	menu := []string{
		"1️⃣ - **Alphabetical** (A-Z)",
		"2️⃣ - **Alphabetical** (Z-A)",
		"3️⃣ - **Rarity** (Common-Collector)",
		"4️⃣ - **Rarity** (Collector-Common)",
		"5️⃣ - **Value** (Low-High)",
		"6️⃣ - **Value** (High-Low)",
		"7️⃣ - **Level** (Low-High)",
		"8️⃣ - **Level** (High-Low)",
	}
	description := "Choose how you want to order your inventory.\n\n"
	for _, option := range menu {
		description += option + "\n"
	}

	discord.ActiveMenus.Set(ctx.ID, &discord.Menus{
		MenuID:        ctx.ID,
		SourceContext: ctx,
		Call:          orderMenu,
	}, 0)

	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       "Order",
			Description: description,
			Color:       config.Botcolor,
			Footer:      discord.DefaultFooter,
		},
		Components: orderComponent(ctx.ID),
		Ephemeral:  true,
	})
}

func orderMenu(ctx *discord.CmdContext) {
	orders := []string{
		"name ASC",
		"name DESC",
		"rarity ASC",
		"rarity DESC",
		"value ASC",
		"value DESC",
		"level ASC",
		"level DESC",
	}
	n, err := strconv.Atoi(strings.Split(ctx.ComponentData.CustomID, "-")[1])
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("Error: %s", err.Error()),
			Ephemeral: true,
		})
		return
	}
	ctx.Player.Settings.OrderBy = orders[n]
	ctx.Player.Settings.Save()
	ctx.Reply(discord.ReplyParams{
		Content:   "Your inventory order has been updated.",
		Ephemeral: true,
	})
}

func orderComponent(menuID string) []discordgo.MessageComponent {
	c := []discordgo.MessageComponent{}
	row := discordgo.ActionsRow{}
	emojis := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣"}
	for i, emoji := range emojis {
		if i == 4 {
			c = append(c, row)
			row = discordgo.ActionsRow{}
		}
		row.Components = append(row.Components, discordgo.Button{
			Style:    discordgo.PrimaryButton,
			Emoji:    discordgo.ComponentEmoji{Name: emoji},
			CustomID: fmt.Sprintf("%s-%d", menuID, i),
		})
	}
	c = append(c, row)
	return c
}
