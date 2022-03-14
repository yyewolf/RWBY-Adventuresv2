package main

import (
	"os"
	"os/signal"
	"rwby-adventures/config"
	commands_missions "rwby-adventures/main/commands/missions"
	"rwby-adventures/main/discord"
	"rwby-adventures/main/web"
	"rwby-adventures/main/websocket"
	"syscall"
)

func main() {
	websocket.StartWebsocket()
	web.StartWeb()
	loadPassives()
	StartDiscord()
	config.LoadCharacters()
	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Session.Close()
}

func loadPassives() {
	commands_missions.LoadPassives()
}
