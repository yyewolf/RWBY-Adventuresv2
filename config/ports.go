package config

var TestMode = true

var (
	TradePort        = ":50"
	TradeTestPort    = ":50000"
	TradeWebsocket   = 9001
	ArenaPort        = ":51"
	ArenaTestPort    = ":50001"
	ArenaWebsocket   = 9002
	DungeonPort      = ":52"
	DungeonTestPort  = ":50002"
	DungeonWebsocket = 9003

	// RPC
	ArenaRPC   = "8000"
	DungeonRPC = 8001

	TradeHost       = "194.163.142.107"
	TradeTestHost   = "localhost"
	ArenaHost       = "194.163.142.107"
	ArenaTestHost   = "localhost"
	DungeonHost     = "194.163.142.107"
	DungeonTestHost = "localhost"
)

func init() {
	if TestMode {
		TradePort = TradeTestPort
		TradeHost = TradeTestHost

		ArenaPort = ArenaTestPort
		ArenaHost = ArenaTestHost

		DungeonPort = DungeonTestPort
		DungeonHost = DungeonTestHost
	}
}
