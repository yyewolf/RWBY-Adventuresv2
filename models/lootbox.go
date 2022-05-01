package models

type PlayerBoxes struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	Boxes          int `gorm:"column:classic_char_boxes;not null;default:0"`
	RareBoxes      int `gorm:"column:rare_char_boxes;not null;default:0"`
	GrimmBoxes     int `gorm:"column:classic_grimm_boxes;not null;default:0"`
	RareGrimmBoxes int `gorm:"column:rare_grimm_boxes;not null;default:0"`
}

type PlayerLootTime struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`
	Amount    int    `gorm:"column:amount;not null"`
	Time      int64  `gorm:"column:time;not null"`
}

type PlayerGamble struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`
	Amount    int    `gorm:"column:amount;not null"`
	Time      int64  `gorm:"column:time;not null"`
}

type LimitedBoxes struct {
	ID        int    `gorm:"primary_key;column:id"`
	DiscordID string `gorm:"column:discord_id;not null"`
	For       string `gorm:"column:for;not null"`
	Type      int    `gorm:"column:type;not null"`
}

type SpecialBoxes struct {
	ID        string `gorm:"primary_key;column:id"`
	DiscordID string `gorm:"column:discord_id;not null"`
	For       string `gorm:"column:for;not null"`
}
