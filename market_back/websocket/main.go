package websocket

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/ambelovsky/gosf"
	"github.com/pmylund/go-cache"
)

var Tokens *cache.Cache

func StartWebsocket() {
	// Start the server using a basic configuration
	gosf.Listen("listings/latest", getLatestListings)
	gosf.Listen("listings/buy", listingPurchase)
	gosf.Listen("auctions/latest", getLatestAuctions)
	gosf.Listen("auctions/bid", auctionBid)
	gosf.Listen("randomPersonas", getRandomPersonas)

	gosf.OnConnect(func(client *gosf.Client, request *gosf.Request) {
		client.Join("*")
	})

	fmt.Println("[MARKET] WS server started on port", config.MarketWebsocket)

	go gosf.Startup(map[string]interface{}{"port": config.MarketWebsocket, "enableCORS": "http://localhost:8080"})
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
