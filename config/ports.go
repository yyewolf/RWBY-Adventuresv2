package config

import "os"

var (
	TradePort          = os.Getenv("TRADE_PORT")
	TradeWebsocketPort = getEnvInt("TRADE_WEBSOCKET_PORT")
	TradeWebsocketURL  = os.Getenv("TRADE_WEBSOCKET_URL")
	TradeHost          = os.Getenv("TRADE_HOST")
	TradeDomain        = os.Getenv("TRADE_DOMAIN")

	ArenaPort      = os.Getenv("ARENA_PORT")
	ArenaWebsocket = getEnvInt("ARENA_WEBSOCKET")
	ArenaRPC       = getEnvInt("ARENA_RPC")
	ArenaHost      = os.Getenv("ARENA_HOST")
	ArenaRPCHost   = os.Getenv("ARENA_RPC_HOST")
	ArenaDomain    = os.Getenv("ARENA_DOMAIN")

	DungeonPort      = os.Getenv("DUNGEON_PORT")
	DungeonWebsocket = getEnvInt("DUNGEON_WEBSOCKET")
	DungeonRPC       = getEnvInt("DUNGEON_RPC")
	DungeonHost      = os.Getenv("DUNGEON_HOST")
	DungeonRPCHost   = os.Getenv("DUNGEON_RPC_HOST")
	DungeonDomain    = os.Getenv("DUNGEON_DOMAIN")

	MarketPort      = os.Getenv("MARKET_PORT")
	MarketWebsocket = getEnvInt("MARKET_WEBSOCKET")
	MarketRPC       = getEnvInt("MARKET_RPC")
	MarketHost      = os.Getenv("MARKET_HOST")
	MarketRPCHost   = os.Getenv("MARKET_RPC_HOST")
	MarketDomain    = os.Getenv("MARKET_DOMAIN")
	MarketFront     = os.Getenv("MARKET_FRONT")

	GamblePort    = os.Getenv("GAMBLE_PORT")
	GambleRPC     = getEnvInt("GAMBLE_RPC")
	GambleRPCHost = os.Getenv("GAMBLE_RPC_HOST")
	GambleHost    = os.Getenv("GAMBLE_HOST")

	TopGGPort    = os.Getenv("TOPGG_PORT")
	TopGGRPC     = getEnvInt("TOPGG_RPC")
	TopGGRPCHost = os.Getenv("TOPGG_RPC_HOST")

	IMGPort = os.Getenv("IMG_PORT")
	IMGHost = os.Getenv("IMG_HOST")
)
