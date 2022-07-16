package websocket

import (
	"fmt"
	"rwby-adventures/config"
	dungeonpc "rwby-adventures/dungeons_rpc"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
)

var Tokens *cache.Cache

func StartWebsocket() {
	// Start the server using a basic configuration
	gosf.Listen("dungeonConnect", DungeonConnect)
	gosf.Listen("dungeonMove", DungeonMove)

	go gosf.Startup(map[string]interface{}{"port": config.DungeonWebsocket, "enableCORS": "http://localhost:8080"})

	fmt.Println("[WS] Started.")

	go CreateDungeon(&dungeonpc.CreateDungeonReq{
		Id: "test",
	})
}

func GetString(request *gosf.Request, key string) (string, bool) {
	data, ok := request.Message.Body[key]
	if !ok {
		return "", false
	}
	str, f := data.(string)
	return str, f
}

func GetInteger(request *gosf.Request, key string) (int, bool) {
	data, ok := request.Message.Body[key]
	if !ok {
		return 0, false
	}
	switch data := data.(type) {
	case int:
		return data, true
	case float64:
		return int(data), true
	default:
		return 0, false
	}
}
