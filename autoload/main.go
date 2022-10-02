package autoload

import "github.com/joho/godotenv"

func init() {
	godotenv.Load("../dev.env")
	godotenv.Load("../../../dev.env")
}
