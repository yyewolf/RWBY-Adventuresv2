package models

import (
	"fmt"
	"rwby-adventures/config"
)

func init() {
	config.Database.AutoMigrate(&Player{})
	config.Database.AutoMigrate(&PlayerStatus{})
	fmt.Println("[DATABASE] Automigrated models.")
}
