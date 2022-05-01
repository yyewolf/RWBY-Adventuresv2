package models

type PlayerShop struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	XPBoost       bool `gorm:"column:xpboost;not null"`
	XPBoostTime   int  `gorm:"column:xpboost_time;not null"`
	LuckBoost     bool `gorm:"column:luckboost;not null"`
	LuckBoostTime int  `gorm:"column:luckboost_time;not null"`
	Extensions    int  `gorm:"column:extensions;not null"`
}
