package commands_misc

import (
	"fmt"
	"runtime"
	"rwby-adventures/config"
	"rwby-adventures/main/arenas"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/dungeons"
	"rwby-adventures/main/gambles"
	"rwby-adventures/main/market"
	"rwby-adventures/main/topgg"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

var StatCommand = &discord.Command{
	Name:        "stats",
	Description: "Check the bot's stats.",
	Menu:        discord.MiscMenu,
	Call:        stats,
}

var startTime = time.Now()

func stats(ctx *discord.CmdContext) {
	//Sends stats to user
	users := 0
	for _, guild := range ctx.Session.State.Ready.Guilds {
		users += len(guild.Members)
	}
	ServerAmount := len(ctx.Session.State.Guilds)
	Uptime := time.Since(startTime)
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	GCTotal := memStats.TotalAlloc - memStats.Alloc
	GCString := ParseByte(GCTotal)
	AllocString := ParseByte(memStats.Alloc)

	var nplayers int64
	config.Database.Model(&models.Player{}).Count(&nplayers)

	var nchar int64
	config.Database.Model(&models.Character{}).Count(&nchar)

	var ngrimms int64
	config.Database.Model(&models.Grimm{}).Count(&ngrimms)

	lines := []string{
		fmt.Sprintf("DiscordGo version :   %s", discordgo.VERSION),
		fmt.Sprintf("Go version :          %s", runtime.Version()),
		fmt.Sprintf("Users cached :        %d", users),
		fmt.Sprintf("Concurrent Tasks :    %d", runtime.NumGoroutine()),
		fmt.Sprintf("Memory :              %s (%s garbage collected)", AllocString, GCString),
		fmt.Sprintf("Uptime :              %d:%d:%d\n", int(Uptime.Hours()), int(Uptime.Minutes())%60, int(Uptime.Seconds())%60),
		fmt.Sprintf("Dungeons :            %s", ParseBool(dungeons.DungeonsMicroservice.Connected())),
		fmt.Sprintf("Market :              %s", ParseBool(market.MarketMicroservice.Connected())),
		fmt.Sprintf("Arenas :              %s", ParseBool(arenas.ArenaMicroservice.Connected())),
		fmt.Sprintf("Gambles :             %s", ParseBool(gambles.GambleMicroservice.Connected())),
		fmt.Sprintf("TopGG :               %s\n", ParseBool(topgg.TopggMicroservice.Connected())),
		fmt.Sprintf("Servers :             %d", ServerAmount),
		fmt.Sprintf("Players :             %d", nplayers),
		fmt.Sprintf("Characters :          %d", nchar),
		fmt.Sprintf("Grimms :              %d", ngrimms),
	}

	var msg string
	msg = "```"
	for _, line := range lines {
		msg += line + "\n"
	}
	msg += "```"

	ctx.Reply(discord.ReplyParams{
		Content:   msg,
		Ephemeral: true,
	})
}

func ParseByte(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ParseBool(b bool) string {
	if b {
		return "ON"
	}
	return "OFF"
}
