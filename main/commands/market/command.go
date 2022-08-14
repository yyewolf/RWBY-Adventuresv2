package commands_market

import (
	"rwby-adventures/main/discord"

	"github.com/bwmarrin/discordgo"
)

var MarketCommand = &discord.Command{
	Name:        "market",
	Description: "All market commands.",
	SubCommandsGroup: []*discord.Command{
		{
			Name:        "listings",
			Description: "All listings commands.",
			SubCommands: []*discord.Command{
				{
					Name:        "create",
					Description: "Create a listing.",
					Menu:        discord.MarketMenu,
					Call:        CreateListing,
					Args: []discord.Arg{
						{
							Name:        "price",
							Description: "Price for the listing.",
							Size:        1,
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionInteger,
						},
						{
							Name:        "id",
							Description: "Identification number of your persona.",
							Size:        1,
							Required:    false,
							Type:        discordgo.ApplicationCommandOptionString,
						},
						{
							Name:        "note",
							Description: "Note for the listing.",
							Size:        200,
							Required:    false,
							Type:        discordgo.ApplicationCommandOptionString,
						},
					},
				},
			},
		},
		{
			Name:        "auctions",
			Description: "All auctions commands.",
			SubCommands: []*discord.Command{
				{
					Name:        "create",
					Description: "Create an auction.",
					Menu:        discord.MarketMenu,
					Call:        CreateAuction,
					Args: []discord.Arg{
						{
							Name:        "id",
							Description: "Identification number of your persona.",
							Size:        1,
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionString,
						},
						{
							Name:        "duration",
							Description: "Duration for the auction.",
							Size:        1,
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionInteger,
						},
					},
				},
			},
		},
	},
}
