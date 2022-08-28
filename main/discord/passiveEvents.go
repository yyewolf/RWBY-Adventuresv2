package discord

import (
	"fmt"
	"math/rand"
	"rwby-adventures/config"
	"rwby-adventures/main/arenas"
	"rwby-adventures/microservices"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

type PassiveFunc func(*CmdContext)

var RegisteredPassives []PassiveFunc

func RegisterPassiveFunction(f PassiveFunc) {
	RegisteredPassives = append(RegisteredPassives, f)
}

func init() {
	RegisterPassiveFunction(TrainCharacter)
	RegisterPassiveFunction(DropBoxes)
	RegisterPassiveFunction(SpawnArena)

	arenas.ArenaMicroservice.Listen("endArena", endArena)
}

func xpLimiter(id string) bool {
	id = id + "-xp"
	val, err := RateLimitCache.Get(id)
	if !err {
		rl := &rateLimit{
			lastMessage: time.Now(),
		}
		RateLimitCache.Set(id, rl, 0)
		return false
	}
	if time.Since(val.(*rateLimit).lastMessage) <= 1000*time.Millisecond {
		return true
	}
	rl := &rateLimit{
		lastMessage: time.Now(),
	}
	RateLimitCache.Set(id, rl, 0)
	return false
}

func TrainCharacter(ctx *CmdContext) {
	/*
		Rate limits :
	*/
	rateLimited := xpLimiter(ctx.Author.ID)
	if rateLimited {
		return
	}
	now, err := ctx.GetUnix()
	if err != nil {
		panic("id not parsable ?")
	}
	timing := now - ctx.Player.Status.LastXP
	if timing < 3 {
		return
	}
	if ctx.Player.SelectedChar == nil && ctx.Player.SelectedGrimm == nil {
		return
	}
	ctx.Player.Status.LastXP = now
	config.Database.Save(ctx.Player.Status)

	XP := ctx.Player.CalcSelectedXP(3, ctx.Player.Shop.XPBoost)
	ctx.GiveSelectionXP(XP, true)

	CP := ctx.Player.CalcCP(0.1)
	ctx.GiveCP(CP, true)
}

func DropBoxes(ctx *CmdContext) {
	canDrop, _ := ctx.Player.CanDropLootBox()
	if !canDrop || ctx.Player.Lootbox() >= ctx.Player.MaxChar() {
		return
	}
	randfloat := rand.Float64()
	randomdrop := randfloat < 0.015
	if !randomdrop {
		return
	}
	isClassic := randfloat < 0.5
	ctx.Player.LastBoxes.Amount++
	ctx.Player.LastBoxes.Time = time.Now().Unix()
	config.Database.Save(ctx.Player.LastBoxes)

	isRare := randfloat < 0.0007
	//Luck Booster
	if ctx.Player.Shop.LuckBoost && !isRare {
		isRare = randfloat < 0.0034
		if isRare {
			ctx.Player.Shop.LuckBoostTime--
			config.Database.Save(ctx.Player.Shop)
			ctx.SendLuckBoostNotice()
		}
	}

	resp := &discordgo.MessageEmbed{
		Title:       "Random Drop",
		Description: fmt.Sprintf("Congratulations !\n%s earned a {{TYPE}} !\nType `%sopen {{CMD}}` to open it.", ctx.Author.Mention(), "/"),
		Color:       config.Botcolor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.redd.it/t2ms587ydsq01.gif",
		},
		Footer: DefaultFooter,
	}

	if isClassic {
		if isRare {
			resp.Description = strings.ReplaceAll(resp.Description, "{{TYPE}}", "Rare Char Box")
			resp.Description = strings.ReplaceAll(resp.Description, "{{CMD}}", "rarechar")
			ctx.Player.Boxes.RareBoxes++
		} else {
			resp.Description = strings.ReplaceAll(resp.Description, "{{TYPE}}", "Normal Char Box")
			resp.Description = strings.ReplaceAll(resp.Description, "{{CMD}}", "normalchar")
			ctx.Player.Boxes.Boxes++
		}
	} else {
		if isRare {
			resp.Description = strings.ReplaceAll(resp.Description, "{{TYPE}}", "Rare Grimm Box")
			resp.Description = strings.ReplaceAll(resp.Description, "{{CMD}}", "raregrimm")
			ctx.Player.Boxes.RareGrimmBoxes++
		} else {
			resp.Description = strings.ReplaceAll(resp.Description, "{{TYPE}}", "Normal Grimm Box")
			resp.Description = strings.ReplaceAll(resp.Description, "{{CMD}}", "normalgrimm")
			ctx.Player.Boxes.GrimmBoxes++
		}
	}

	config.Database.Save(ctx.Player.Boxes)

	ctx.Reply(ReplyParams{
		Content: resp,
	})
}

func SpawnArena(ctx *CmdContext) {
	if ctx.GuildID == "" {
		return
	}
	// Every 5 hours
	if ctx.Guild.LastArena < 5*60*60 {
		return
	}
	ctx.Guild.LastArena = time.Now().Unix()
	config.Database.Save(ctx.Guild)
	//Rolls the dice now to save computing power
	r := rand.Float64()
	//0.12% chance
	willSpawn := r < 0.0012
	if !willSpawn && ctx.Author.ID != "144472011924570113" {
		return
	}

	c, err := ctx.Session.Channel(ctx.ChannelID)
	if err != nil {
		return
	}
	if c.Type == discordgo.ChannelTypeDM {
		return
	}

	ID := uuid.NewV4().String()
	in := &microservices.CreateArena{
		ID:        ID,
		ChannelID: ctx.ChannelID,
	}

	if !arenas.ArenaMicroservice.Connected() {
		return
	}

	//Prepare the embed
	msg := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "**Grimm** has appeared !",
			Description: "Click the button to join the fight !\nPrize : XP and Money !",
			Color:       config.Botcolor,
		},
		Content: ctx.Guild.PingRoles,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Join!",
						Style: discordgo.LinkButton,
						URL:   fmt.Sprintf("http://%s%s/a/%s", config.ArenaHost, config.ArenaPort, ID),
					},
				},
			},
		},
	}
	// Send the message
	ctx.Reply(ReplyParams{
		Content:   msg,
		Automated: true,
	})

	arenas.CreateArena(in)
}
