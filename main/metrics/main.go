package metrics

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func StartMetrics() {
	go func() {
		log.Println(http.ListenAndServe(":2112", nil))
	}()
}
