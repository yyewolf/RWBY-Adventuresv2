package main

import (
	"encoding/json"
	"errors"
)

type marketBidReqOut struct {
	Action     string `json:"act"`
	SellerID   string `json:"seller_id"`
	ReceiverID string `json:"receiver_id"`
	Price      int    `json:"price"`
	CharName   string `json:"char_name"`
}

type marketAuctionBidder struct {
	ID    string `json:"user_id"`
	Price int    `json:"price"`
}

type marketAuction struct {
	ID             string                `json:"id"`
	SellerID       string                `json:"seller_id"`
	SellerName     string                `json:"seller_name"`
	StartedAt      int64                 `json:"started_at"`
	EndsAt         int64                 `json:"ends_at"`
	TimeExtensions int64                 `json:"extends"`
	Type           int                   `json:"type"`
	Char           storedCharacter       `json:"chars"`
	Grimm          storedGrimm           `json:"grimms"`
	History        []marketAuctionBidder `json:"bidders"`
}

func getAllAuctions(name string) (a []marketAuction) {
	d, _ := database.Select(`a.*, 
		(select row_to_json(allc) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = a.character_id) allc) chars, 
		(select row_to_json(allc) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = a.grimm_id) all) grimms,
		(select ('[' || array_to_string(array_agg(row_to_json(bll)), ',') || ']')::json chars from (select * from auctions_bidder bi where bi.auction_id = a.id order by bi.price desc) bll) bidders`).
		From(`auctions a
		left join characters c on c.id = a.character_id
		left join grimms g on g.id = a.grimm_id`).
		Where("c.name = $1", name).
		GroupBy("a.id").
		QueryJSON()
	json.Unmarshal(d, &a)
	return
}

func getAllAuctionsFromPlayer(id string) (a []marketAuction) {
	d, _ := database.Select(`a.*, 
		(select row_to_json(allc) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = a.character_id) allc) chars,
		(select row_to_json(allc) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = a.grimm_id) all) grimms, 
		(select ('[' || array_to_string(array_agg(row_to_json(bll)), ',') || ']')::json chars from (select * from auctions_bidder bi where bi.auction_id = a.id order by bi.price desc) bll) bidders`).
		From(`auctions a
		left join characters c on c.id = a.character_id
		left join grimms g on g.id = a.grimm_id`).
		Where("a.seller_id = $1", id).
		GroupBy("a.id").
		QueryJSON()
	json.Unmarshal(d, &a)
	return
}

func getAuction(id string) (marketAuction, error) {
	a := []marketAuction{}
	d, _ := database.Select(`a.*, 
		(select row_to_json(allc) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = a.character_id) allc) chars,
		(select row_to_json(allc) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = a.grimm_id) allc) grimms, 
		(select ('[' || array_to_string(array_agg(row_to_json(bll)), ',') || ']')::json chars from (select * from auctions_bidder bi where bi.auction_id = a.id order by bi.price desc) bll) bidders`).
		From(`auctions a
		left join characters c on c.id = a.character_id
		left join grimms g on g.id = a.grimm_id`).
		Where("a.id = $1", id).
		GroupBy("a.id").
		Limit(1).
		QueryJSON()
	e := json.Unmarshal(d, &a)
	if len(a) == 0 {
		e := errors.New("err")
		return marketAuction{}, e
	}
	return a[0], e
}
