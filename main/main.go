package main

import (
	"os"
	"os/signal"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"syscall"
)

func main() {
	StartDiscord()
	config.LoadCharacters()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Session.Close()
}
