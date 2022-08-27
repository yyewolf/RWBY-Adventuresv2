package config

import (
	"os"
	"strconv"

	_ "rwby-adventures/autoload"
)

func getEnvInt(key string) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0
	}
	return v
}
