package models

import (
	"errors"
	"rwby-adventures/config"

	"github.com/lib/pq"
)

type TradeContent struct {
	TradeID string `gorm:"primary_key;column:trade_id"`
	Type    string `gorm:"primary_key;column:type"`

	UserID         string         `gorm:"column:user_id"`
	Characters     pq.StringArray `gorm:"column:characters;not null;type:text[]"`
	Grimms         pq.StringArray `gorm:"column:grimms;not null;type:text[]"`
	Money          int64          `gorm:"column:money;not null"`
	Boxes          int64          `gorm:"column:boxes;not null"`
	RareBoxes      int64          `gorm:"column:rare_boxes;not null"`
	GrimmBoxes     int64          `gorm:"column:grimm_boxes;not null"`
	RareGrimmBoxes int64          `gorm:"column:rare_grimm_boxes;not null"`
}

type Trade struct {
	ID         string `gorm:"primary_key;column:id"`
	SenderID   string `gorm:"column:sender_id;not null"`
	ReceiverID string `gorm:"column:receiver_id;not null"`
	StartedAt  int64  `gorm:"column:started_at;not null"`

	// Filled later
	UserSends   *TradeContent `gorm:"-"`
	TargetSends *TradeContent `gorm:"-"`
}

func GetTrade(id string) (t *Trade, err error) {
	// We get the main trade
	t = &Trade{
		ID:          id,
		UserSends:   &TradeContent{},
		TargetSends: &TradeContent{},
	}
	e := config.Database.Find(t, id)
	if e.Error != nil || e.RowsAffected == 0 {
		return t, errors.New("not found")
	}
	e = config.Database.Find(t.UserSends, id, t.SenderID)
	if e.Error != nil || e.RowsAffected == 0 {
		return t, errors.New("not found")
	}
	e = config.Database.Find(t.TargetSends, id, t.ReceiverID)
	if e.Error != nil || e.RowsAffected == 0 {
		return t, errors.New("not found")
	}
	return
}

func (t *Trade) Delete() {
	config.Database.Delete(&TradeContent{}, t.ID)
	config.Database.Delete(t)
}
