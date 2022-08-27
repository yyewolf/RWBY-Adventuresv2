package config

var TestMode = true

var (
	TradePort        = ":50"
	TradeTestPort    = ":50000"
	TradeWebsocket   = 9000
	ArenaPort        = ":51"
	ArenaTestPort    = ":50001"
	ArenaWebsocket   = 9001
	DungeonPort      = ":52"
	DungeonTestPort  = ":50002"
	DungeonWebsocket = 9002
	MarketPort       = ":53"
	MarketTestPort   = ":50003"
	MarketWebsocket  = 9003
	TopGGPort        = ":54"
	TopGGTestPort    = ":50004"

	// RPC
	ArenaRPC   = 8001
	DungeonRPC = 8002
	MarketRPC  = 8003
	TopGGRPC   = 8004

	TradeHost       = "194.163.142.107"
	TradeTestHost   = "localhost"
	ArenaHost       = "194.163.142.107"
	ArenaTestHost   = "localhost"
	DungeonHost     = "194.163.142.107"
	DungeonTestHost = "localhost"
	MarketHost      = "194.163.142.107"
	MarketTestHost  = "localhost"
	TopGGHost       = "194.163.142.107"
	TopGGTestHost   = "localhost"
)

func init() {
	if TestMode {
		TradePort = TradeTestPort
		TradeHost = TradeTestHost

		ArenaPort = ArenaTestPort
		ArenaHost = ArenaTestHost

		DungeonPort = DungeonTestPort
		DungeonHost = DungeonTestHost

		MarketPort = MarketTestPort
		MarketHost = MarketTestHost

		TopGGPort = TopGGTestPort
		TopGGHost = TopGGTestHost
	}
}
