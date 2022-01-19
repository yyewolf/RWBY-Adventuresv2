package websocket

import (
	"fmt"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
)

var Tokens *cache.Cache

func StartWebsocket() {
	// Start the server using a basic configuration
	gosf.Listen("tradeConnect", TradeConnect)
	gosf.Listen("tradeInfos", TradeInfos)
	gosf.Listen("tradeValidate", TradeValidate)
	go gosf.Startup(map[string]interface{}{"port": 9999})
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
