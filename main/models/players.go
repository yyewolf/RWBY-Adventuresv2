package models

import "gorm.io/gorm"

type PlayerStatus struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`

	AuctioningID   string `gorm:"column:auctioning_id"`
	AuctioningTime int64  `gorm:"column:auctioning_time"`
	DeletingID     string `gorm:"column:deleting_id"`

	LastXP      int64 `gorm:"column:last_xp"`
	Voted       bool  `gorm:"column:has_voted"`
	DailyStreak int   `gorm:"column:daily_streak"`
	LastOpening int64 `gorm:"column:last_opening"`

	MarketPublic  bool `gorm:"column:market_public"`
	ProfilePublic bool `gorm:"column:profile_public"`

	LastReport  int64 `gorm:"column:last_report"`
	LastDungeon int64 `gorm:"column:last_dungeon"`
}

type PlayerMission struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`

	//Mission Related
	CanGoToMission bool `gorm:"column:mission_can"`
	IsInMission    bool `gorm:"column:mission_in"`
	MissionType    int  `gorm:"column:mission_type"`
	MissionMsgLeft int  `gorm:"column:mission_msgleft"`

	//Hunt Related
	CanGoHunt   bool `gorm:"column:hunt_can"`
	IsInHunt    bool `gorm:"column:hunt_in"`
	HuntType    int  `gorm:"column:hunt_type"`
	HuntMsgLeft int  `gorm:"column:hunt_msgleft"`
}

type Player struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`
	IsNew     bool   `gorm:"column:is_new"`

	Balance int64 `gorm:"column:balance"`

	Level int64 `gorm:"column:level"`
	CP    int64 `gorm:"column:cp"`
	MaxCP int64 `gorm:"column:max_cp"`

	CharLimit int `gorm:"column:max_char"`

	Maxlootbox int `gorm:"column:max_lootbox"`

	SelectedID   string `gorm:"column:selected_id"`
	SelectedType int    `gorm:"column:selected_type"`

	Badges   int64 `gorm:"column:badges"`
	Settings int64 `gorm:"column:settings"`
	Disabled bool  `gorm:"column:disabled"`
	// Foreign keys
	Missions PlayerMission `gorm:"foreignkey:DiscordID"`
	Status   PlayerStatus  `gorm:"foreignkey:DiscordID"`
}
