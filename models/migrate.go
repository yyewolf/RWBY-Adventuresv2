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
	config.Database.AutoMigrate(
		&PlayerStats{},
		&Badges{},
		&PlayerBadges{},
	)
	config.Database.AutoMigrate(
		&Submission{},
		&SubmissionFile{},
	)
	for i, b := range DefaultBadges {
		b.BadgeID = i + 1
		b.Save()
	}

	fmt.Println("[DATABASE] Automigrated models.")
}
