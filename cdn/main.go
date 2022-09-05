//go:generate goversioninfo
package main

import (
	"embed"
	_ "embed"
	"os"
	"os/signal"
	"syscall"
)

//go:embed data
var webbox embed.FS

func main() {
	go hostImageService()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
