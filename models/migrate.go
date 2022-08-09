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
		&PlayerSettings{},
		&PlayerShop{},
	)
	config.Database.AutoMigrate(
		&Character{},
		&CharacterStats{},
		&Grimm{},
		&GrimmStat{},
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
	config.Database.AutoMigrate(
		&Listing{},
		&Auction{},
		&AuctionBidders{},
	)
	config.Database.AutoMigrate(
		&Trade{},
		&TradeContent{},
	)
	fmt.Println("[DATABASE] Automigrated models.")
}
