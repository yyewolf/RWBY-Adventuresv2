package models

import (
	"fmt"
	"rwby-adventures/config"
)

func init() {
	config.Database.AutoMigrate(
		&Player{},
		&PlayerStatus{},
		&PlayerMission{},
	)
	fmt.Println("[DATABASE] Automigrated models.")
}
