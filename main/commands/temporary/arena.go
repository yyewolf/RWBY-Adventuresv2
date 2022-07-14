package commands_temporary

import (
	"context"
	"fmt"
	arenapc "rwby-adventures/arenas_rpc"
	"rwby-adventures/main/discord"
	rwby_grpc "rwby-adventures/main/grpc"

	uuid "github.com/satori/go.uuid"
)

var ArenaCommand = &discord.Command{
	Name:        "arena",
	Description: "Create arenas.",
	Menu:        discord.GeneralMenu,
	Call:        createArena,
}

func createArena(ctx *discord.CmdContext) {
	ID := uuid.NewV4().String()
	in := &arenapc.CreateArenaReq{
		Id: ID,
	}

	ctx.Reply(discord.ReplyParams{
		Content: fmt.Sprintf("Creating arena with ID %s", ID),
	})
	rep, err := rwby_grpc.ArenaServer.CreateArena(context.Background(), in)
	if err != nil {
		ctx.Reply(discord.ReplyParams{
			Content: fmt.Sprint(err),
		})
		return
	}
	fmt.Println(rep)
	if rep.Status == 1 {
		ctx.Reply(discord.ReplyParams{
			Content: "Arena creation failed.",
		})
		return
	}
	ctx.Reply(discord.ReplyParams{
		Content: rep.GetLoots(),
	})
}
