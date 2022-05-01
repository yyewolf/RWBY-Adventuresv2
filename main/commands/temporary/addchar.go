package commands_temporary

import (
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

func Addchar(ctx *discord.CmdContext) {
	ID := uuid.NewV4().String()
	config.Database.Create(&models.Character{
		CharID: ID,
		Level:  rand.Intn(20) + 5,
		Rarity: rand.Intn(6),
		XP:     400,
		XPCap:  5000,
		Name:   config.BaseCharacters[rand.Intn(len(config.BaseCharacters))].Name,
		UserID: ctx.Author.ID,
		Stats: models.CharacterStats{
			Value: rand.Float64() * 100,
		},
		OwnedAt: time.Now(),
	})
	ctx.Reply(discord.ReplyParams{
		Content: "Ok boomer.",
	})
}
