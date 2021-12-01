package models

type PlayerBoxes struct {
	DiscordID string `gorm:"primary_key;column:discord_id"`

	Boxes          int `json:"lootboxleft" db:"lootboxleft"`
	RareBoxes      int `json:"rarelootboxleft" db:"rarelootboxleft"`
	GrimmBoxes     int `json:"grimmboxleft" db:"grimmboxleft"`
	RareGrimmBoxes int `json:"raregrimmboxleft" db:"raregrimmboxleft"`
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
