package config

import "os"

var (
	TradePort      = os.Getenv("TRADE_PORT")
	TradeWebsocket = getEnvInt("TRADE_WEBSOCKET")
	TradeHost      = os.Getenv("TRADE_HOST")

	ArenaPort      = os.Getenv("ARENA_PORT")
	ArenaWebsocket = getEnvInt("ARENA_WEBSOCKET")
	ArenaRPC       = getEnvInt("ARENA_RPC")
	ArenaHost      = os.Getenv("ARENA_HOST")

	DungeonPort      = os.Getenv("DUNGEON_PORT")
	DungeonWebsocket = getEnvInt("DUNGEON_WEBSOCKET")
	DungeonRPC       = getEnvInt("DUNGEON_RPC")
	DungeonHost      = os.Getenv("DUNGEON_HOST")

	MarketPort      = os.Getenv("MARKET_PORT")
	MarketWebsocket = getEnvInt("MARKET_WEBSOCKET")
	MarketRPC       = getEnvInt("MARKET_RPC")
	MarketHost      = os.Getenv("MARKET_HOST")

	TopGGPort = os.Getenv("TOPGG_PORT")
	TopGGRPC  = getEnvInt("TOPGG_RPC")
	TopGGHost = os.Getenv("TOPGG_HOST")
)
