package commands_boxes

import "rwby-adventures/main/discord"

var BoxesCommand = &discord.Command{
	Name:        "open",
	Description: "To open boxes.",
	SubCommandsGroup: []*discord.Command{
		{
			Name:        "char",
			Description: "To open a character box.",
			SubCommands: []*discord.Command{
				//{
				//Name:        "character",
				//Description: "Open boxes for characters",
				//SubCommands: []*discord.Command{
				{
					Name:        "normal",
					Description: "Open a normal character box",
					Menu:        discord.BoxMenu,
					Call:        OpenClassicChar,
				},
				{
					Name:        "rare",
					Description: "Open a rare character box",
					Menu:        discord.BoxMenu,
					Call:        OpenRareChar,
				},
			},
		},
		{
			Name:        "grimm",
			Description: "To open a grimm box.",
			SubCommands: []*discord.Command{
				//	},
				//},
				//{
				//	Name:        "grimm",
				//	Description: "Open boxes for grimms",
				//	SubCommands: []*discord.Command{
				{
					Name:        "normal",
					Description: "Open a normal grimm box",
					Menu:        discord.BoxMenu,
					Call:        OpenClassicGrimm,
				},
				{
					Name:        "rare",
					Description: "Open a rare grimm box",
					Menu:        discord.BoxMenu,
					Call:        OpenRareGrimm,
				},
			},
		},
		{
			Name:        "other",
			Description: "To open other boxes.",
			SubCommands: []*discord.Command{
				//	},
				//},
				{
					Name:        "special",
					Description: "Open a special box",
					Menu:        discord.BoxMenu,
					Call:        OpenSpecialCmd,
				},
				{
					Name:        "limited",
					Description: "Open a limited box",
					Menu:        discord.BoxMenu,
					Call:        OpenLimited,
				},
			},
		},
	},
}
