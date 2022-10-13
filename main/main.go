package main

import (
	"os"
	"os/signal"
	commands_admin "rwby-adventures/main/commands/admin"
	commands_missions "rwby-adventures/main/commands/missions"
	"rwby-adventures/main/discord"
	topgg "rwby-adventures/main/topgg"
	"rwby-adventures/main/web"
	"rwby-adventures/main/websocket"
	"syscall"
)

func main() {
	topgg.StartTopGG()
	websocket.StartWebsocket()
	web.StartWeb()
	StartDiscord()
	loadPassives()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Session.Close()
}

func loadPassives() {
	commands_missions.LoadPassives()
	commands_admin.LoadPassives()
}
