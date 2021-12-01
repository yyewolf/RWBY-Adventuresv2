package models

type PlayerStatus struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	AuctioningID   string `gorm:"column:auctioning_id;not null"`
	AuctioningTime int64  `gorm:"column:auctioning_time;not null"`
	DeletingID     string `gorm:"column:deleting_id;not null"`

	LastXP      int64 `gorm:"column:last_xp;not null"`
	Voted       bool  `gorm:"column:has_voted;not null"`
	DailyStreak int   `gorm:"column:daily_streak;not null"`
	LastOpening int64 `gorm:"column:last_opening;not null"`

	MarketPublic  bool `gorm:"column:market_public;not null"`
	ProfilePublic bool `gorm:"column:profile_public;not null"`

	LastReport  int64 `gorm:"column:last_report;not null"`
	LastDungeon int64 `gorm:"column:last_dungeon;not null"`

	OrderBy string `gorm:"column:order_by;not null;default:''"`
}
