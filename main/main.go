package main

import (
	"os"
	"os/signal"
	"rwby-adventures/discord"
	"syscall"
)

func main() {
	StartDiscord()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Session.Close()
}
