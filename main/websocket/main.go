package websocket

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/pmylund/go-cache"
	"github.com/yyewolf/gosf"
)

var Tokens *cache.Cache

func StartWebsocket() {
	// Start the server using a basic configuration
	gosf.Listen("tradeConnect", TradeConnect)
	gosf.Listen("tradeInfos", TradeInfos)
	gosf.Listen("tradeValidate", TradeValidate)
	go gosf.Startup(map[string]interface{}{"port": config.TradeWebsocketPort, "enableCORS": config.TradeHost})
	fmt.Println("[WS] Started.")
}

func GetString(request *gosf.Request, key string) (string, bool) {
	data, ok := request.Message.Body[key]
	if !ok {
		return "", false
	}
	str, f := data.(string)
	return str, f
}
