package market

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/yyewolf/gosf"
)

var MarketMicroservice *gosf.Microservice

func init() {
	gosf.RegisterMicroservice("market", "127.0.0.1", config.MarketRPC, false)
	MarketMicroservice = gosf.GetMicroservice("market")
	MarketMicroservice.Listen("sendMessage", sendMessage)
	fmt.Println("[MARKET] Initialized market microservice.")
}
