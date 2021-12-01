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
		&PlayerShop{},
	)
	config.Database.AutoMigrate(
		&PlayerCharacter{},
		&PlayerCharacterStats{},
		&PlayerGrimm{},
		&PlayerCharacterStats{},
	)
	config.Database.AutoMigrate(
		&PlayerLootTime{},
		&PlayerGamble{},
		&LimitedBoxes{},
		&SpecialBoxes{},
	)
	fmt.Println("[DATABASE] Automigrated models.")
}
