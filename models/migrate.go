package models

import (
	"fmt"
	"rwby-adventures/main/config"
)

func init() {
	config.Database.AutoMigrate(
		&Player{},
		&PlayerStatus{},
		&PlayerMission{},
		&PlayerShop{},
	)
	config.Database.AutoMigrate(
		&Character{},
		&CharacterStats{},
		&Grimm{},
		&CharacterStats{},
	)
	config.Database.AutoMigrate(
		&PlayerLootTime{},
		&PlayerGamble{},
		&LimitedBoxes{},
		&SpecialBoxes{},
		&PlayerBoxes{},
	)
	config.Database.AutoMigrate(
		&Guild{},
	)
	fmt.Println("[DATABASE] Automigrated models.")
}
