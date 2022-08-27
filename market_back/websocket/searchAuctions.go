package websocket

import (
	"encoding/json"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/models"

	"github.com/yyewolf/gosf"
)

func searchAuctions(client *gosf.Client, request *gosf.Request) *gosf.Message {
	filters := SearchFilters{}
	data, err := json.Marshal(request.Message.Body)
	if err != nil {
		return gosf.NewFailureMessage(err)
	}
	err = json.Unmarshal(data, &filters)
	if err != nil {
		return gosf.NewFailureMessage(err)
	}

	var o = "auction_id"
	var t = "desc"
	var c bool
	switch filters.OrderBy {
	case "price":
		o = "price"
	case "value":
		o = "value"
		c = true
	case "level":
		o = "level"
		c = true
	case "rarity":
		o = "rarity"
		c = true
	case "id":
		o = "auction_id"
	case "type":
		o = "type"
	}

	switch filters.OrderType {
	case "asc":
		t = "asc"
	case "desc":
		t = "desc"
	}

	if c {
		o = fmt.Sprintf("\"Char\".\"%s\" %s, \"Grimm\".\"%s\" %s", o, t, o, t)
	}

	var m []*models.Auction
	var icons []string
	config.Database.
		Order(o).
		Joins("Char").
		Joins("Grimm").
		Where("(type = 0 and \"Char\".\"name\" like ? and \"Char\".\"value\" >= ? and \"Char\".\"value\" <= ? and \"Char\".\"level\" >= ? and \"Char\".\"level\" <= ? and \"Char\".\"buffs\" >= ? and \"Char\".\"buffs\" <= ? and \"Char\".\"rarity\" >= ? and \"Char\".\"rarity\" <= ?) or (type = 1 and \"Grimm\".\"name\" like ? and \"Grimm\".\"value\" >= ? and \"Grimm\".\"value\" <= ? and \"Grimm\".\"level\" >= ? and \"Grimm\".\"level\" <= ? and \"Grimm\".\"buffs\" >= ? and \"Grimm\".\"buffs\" <= ? and \"Grimm\".\"rarity\" >= ? and \"Grimm\".\"rarity\" <= ?)", "%"+filters.NameHas+"%", filters.ValueAbove, filters.ValueBelow, filters.LevelAbove, filters.LevelBelow, filters.BuffsAbove, filters.BuffsBelow, filters.RarityAbove, filters.RarityBelow, "%"+filters.NameHas+"%", filters.ValueAbove, filters.ValueBelow, filters.LevelAbove, filters.LevelBelow, filters.BuffsAbove, filters.BuffsBelow, filters.RarityAbove, filters.RarityBelow).
		Find(&m)

	for _, a := range m {
		config.Database.Order("bid desc").Find(&a.Bidders, "auction_id = ?", a.ID)

		if a.Type == models.CharType {
			icons = append(icons, a.Char.ToRealChar().IconURL)
		} else {
			icons = append(icons, a.Grimm.ToRealGrimm().IconURL)
		}
	}

	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["auctions"] = m
	msg.Body["icons"] = icons
	return msg
}
