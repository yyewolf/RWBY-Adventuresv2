package config

var TestMode = false

var (
	TradePort     = ":50"
	TradeTestPort = ":1000"

	TradeHost     = "194.163.142.107"
	TradeTestHost = "localhost"
)

func init() {
	if TestMode {
		TradePort = TradeTestPort
		TradeHost = TradeTestHost
	}
}
