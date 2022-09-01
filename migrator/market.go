package main

import (
	"encoding/json"
	"errors"
)

type marketListing struct {
	ID         string          `json:"id"`
	SellerID   string          `json:"seller_id"`
	SellerName string          `json:"seller_name"`
	Price      int             `json:"price"`
	Note       string          `json:"note"`
	Type       int             `json:"type"`
	Char       storedCharacter `json:"chars"`
	Grimm      storedGrimm     `json:"grimms"`
}

func getAllListingFromPlayer(id string) (f []marketListing) {
	f = []marketListing{}
	d, _ := database.Select(`m.*, 
		(select row_to_json(a) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = m.char_id) a) chars,
		(select row_to_json(a) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = m.grimm_id) a) grimms`).
		From(`market m`).
		Where("m.seller_id = $1", id).
		GroupBy("m.id").
		QueryJSON()
	json.Unmarshal(d, &f)
	return
}

func getListing(id string) (marketListing, error) {
	m := []marketListing{}
	d, _ := database.Select(`m.*, 
		(select row_to_json(a) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = m.char_id) a) chars,
		(select row_to_json(a) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = m.grimm_id) a) grimms`).
		From(`market m`).
		Where("m.id = $1", id).
		GroupBy("m.id").
		Limit(1).
		QueryJSON()
	e := json.Unmarshal(d, &m)
	if len(m) == 0 {
		e := errors.New("err")
		return marketListing{}, e
	}
	return m[0], e
}
