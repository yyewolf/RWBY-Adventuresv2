package market

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

var TopggMicroservice *gosf.Microservice

func StartTopGG() {
	gosf.RegisterMicroservice("topgg", "127.0.0.1", config.TopGGRPC, false)
	TopggMicroservice = gosf.GetMicroservice("topgg")
	TopggMicroservice.Listen("sendMessage", sendMessage)
	fmt.Println("[TOPGG] Initialized microservice.")
}
