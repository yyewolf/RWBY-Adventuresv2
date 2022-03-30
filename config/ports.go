package config

var TestMode = true

var (
	TradePort      = ":50"
	TradeTestPort  = ":50000"
	TradeWebsocket = 9001
	ArenaPort      = ":51"
	ArenaTestPort  = ":50001"
	ArenaWebsocket = 9000

	TradeHost     = "194.163.142.107"
	TradeTestHost = "localhost"
	ArenaHost     = "194.163.142.107"
	ArenaTestHost = "localhost"
)

func init() {
	if TestMode {
		TradePort = TradeTestPort
		TradeHost = TradeTestHost

		ArenaPort = ArenaTestPort
		ArenaHost = ArenaTestHost
	}
}
