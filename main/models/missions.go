package models

import "github.com/jinzhu/gorm"

type PlayerMission struct {
	gorm.Model
	DiscordID string `gorm:"primary_key;column:discord_id"`

	//Mission Related
	CanGoToMission bool `gorm:"column:mission_can;not null"`
	IsInMission    bool `gorm:"column:mission_in;not null"`
	MissionType    int  `gorm:"column:mission_type;not null"`
	MissionMsgLeft int  `gorm:"column:mission_msgleft;not null"`

	//Hunt Related
	CanGoHunt   bool `gorm:"column:hunt_can;not null"`
	IsInHunt    bool `gorm:"column:hunt_in;not null"`
	HuntType    int  `gorm:"column:hunt_type;not null"`
	HuntMsgLeft int  `gorm:"column:hunt_msgleft;not null"`
}
