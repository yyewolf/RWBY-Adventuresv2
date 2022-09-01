package main

import (
	"encoding/json"
	"math"
	"math/rand"
	"time"
)

// CharacterStatsStruct represents the stats of a character
type CharacterStatsStruct struct {
	CharID      string  `json:"char_id" db:"char_id"`
	Value       float64 `json:"value" db:"value"`
	Health      int     `json:"health"  db:"health"`
	Armor       int     `json:"armor" db:"armor"`
	Damage      int     `json:"damage" db:"damage"`
	Healing     int     `json:"healing" db:"healing"`
	DodgeChance int     `json:"dodge_chance" db:"dodge_chance"`
}

type GrimmStatsStruct struct {
	CharID      string  `json:"grimm_id" db:"grimm_id"`
	Value       float64 `json:"value" db:"value"`
	Health      int     `json:"health"  db:"health"`
	Armor       int     `json:"armor" db:"armor"`
	Damage      int     `json:"damage" db:"damage"`
	Healing     int     `json:"healing" db:"healing"`
	DodgeChance int     `json:"dodge_chance" db:"dodge_chance"`
}

type storedCharacter struct {
	Name          string `json:"name" db:"name"`
	CustomID      string `json:"id" db:"id"`
	UserID        string `json:"user_id" db:"user_id"`
	Level         int    `json:"lvl" db:"lvl"`
	XP            int64  `json:"xp" db:"xp"`
	XPCap         int64  `json:"xp_max" db:"xp_max"`
	Rarity        int    `json:"rarity" db:"rarity"`
	InMission     bool   `json:"in_mission" db:"in_mission"`
	IsInFavorites bool   `json:"is_in_favorites" db:"is_in_favorites"`
	Buffs         int    `json:"buffs" db:"buffs"`
	OwnedAt       int64  `json:"owned_at" db:"owned_at"`
	//Populated somewhere else
	Stats CharacterStatsStruct `json:"stats" db:"-"`
}

type storedGrimm struct {
	Name          string `json:"name" db:"name"`
	CustomID      string `json:"id" db:"id"`
	UserID        string `json:"user_id" db:"user_id"`
	Level         int    `json:"lvl" db:"lvl"`
	XP            int64  `json:"xp" db:"xp"`
	XPCap         int64  `json:"xp_max" db:"xp_max"`
	Rarity        int    `json:"rarity" db:"rarity"`
	InHunt        bool   `json:"in_hunt" db:"in_hunt"`
	IsInFavorites bool   `json:"is_in_favorites" db:"is_in_favorites"`
	Buffs         int    `json:"buffs" db:"buffs"`
	OwnedAt       int64  `json:"owned_at" db:"owned_at"`
	//Populated somewhere else
	Stats GrimmStatsStruct `json:"stats" db:"-"`
}

type shop struct {
	UserID        string `json:"user_id" db:"user_id"`
	XPBoost       bool   `json:"xpboost" db:"xpboost"`
	XPBoostTime   int    `json:"xpboost_time" db:"xpboost_time"`
	LuckBoost     bool   `json:"luckboost" db:"luckboost"`
	LuckBoostTime int    `json:"luckboost_time" db:"luckboost_time"`
	Extenders     int    `json:"extensions" db:"extensions"`
}

type lootBoxTime struct {
	UserID string `json:"user_id" db:"user_id"`
	Amount int    `json:"amount" db:"amount"`
	Time   int64  `json:"time" db:"time"`
}

type gambleTime struct {
	UserID string `json:"user_id" db:"user_id"`
	Amount int    `json:"amount" db:"amount"`
	Time   int64  `json:"time" db:"time"`
}

type playerSort struct {
	UserID string `json:"user_id" db:"user_id"`
	Sort   string `json:"sort" db:"sort"`
}

type limitedEditionLootBox struct {
	UserID string `json:"user_id" db:"user_id"`
	For    string `json:"for" db:"for"`
}

type specialEditionLootBox struct {
	UserID string `json:"user_id" db:"user_id"`
	For    string `json:"for" db:"for"`
}

type playerBadgeCounter struct {
	UserID          string `db:"id"`
	ArenasCompleted int    `db:"arenas"`
	BattlesWon      int    `db:"battles_won"`
	BattlesLost     int    `db:"battles_lost"`
	LootboxOpened   int    `db:"lootbox_opened"`
	RoleplaySent    int    `db:"roleplay_sent"`
	DungeonDone     int    `db:"dungeon_done"`
	MarketSold      int    `db:"market_sold"`
	MarketBought    int    `db:"market_bought"`
}

type player struct {
	ID             string `json:"id" db:"id"`
	Level          int    `json:"lvl" db:"lvl"`
	CP             int64  `json:"cp" db:"cp"`
	MaxCP          int64  `json:"cp_max" db:"cp_max"`
	CharLimit      int    `json:"char_limit" db:"char_limit"`
	IsNew          bool   `json:"is_new" db:"is_new"`
	Maxlootbox     int    `json:"max_lootbox" db:"max_lootbox"`
	IsAuctioning   string `json:"auctioning_id" db:"auctioning_id"`
	AuctioningTime int64  `json:"auctioning_time" db:"auctioning_time"`
	IsDeleting     string `json:"deleting_id" db:"deleting_id"`
	LastXP         int64  `json:"last_xp" db:"last_xp"`
	Voted          bool   `json:"has_voted" db:"has_voted"`
	DailyStreak    int    `json:"daily_streak" db:"daily_streak"`
	LastOpening    int64  `json:"last_opening" db:"last_opening"`
	SelectedID     string `json:"selected_id" db:"selected_id"`
	SelectedType   int    `json:"selected_type" db:"selected_type"`
	MarketPublic   bool   `json:"market_public" db:"market_public"`
	ProfilePublic  bool   `json:"profile_public" db:"profile_public"`
	Patreon        int    `json:"patreon" db:"patreon"`
	PatreonClaim   bool   `json:"patreonclaim" db:"patreonclaim"`
	LastReport     int64  `json:"last_report" db:"last_report"`
	Badges         int64  `json:"badges" db:"badges"`
	Settings       int64  `json:"settings" db:"settings"`
	Disabled       bool   `json:"disabled" db:"disabled"`
	//STRQ Event
	PiecesOfPhoto int `json:"piecesofphoto" db:"piecesofphoto"`
	//"Inventory"
	LootBoxLeft      int   `json:"lootboxleft" db:"lootboxleft"`
	RareLootBoxLeft  int   `json:"rarelootboxleft" db:"rarelootboxleft"`
	GrimmBoxLeft     int   `json:"grimmboxleft" db:"grimmboxleft"`
	RareGrimmBoxLeft int   `json:"raregrimmboxleft" db:"raregrimmboxleft"`
	ArmsLeft         int   `json:"armsleft" db:"armsleft"`
	MinionsLeft      int   `json:"minionsleft" db:"minionsleft"`
	Money            int64 `json:"money" db:"money"`
	XPJar            int64 `json:"xp_jar" db:"xp_jar"`
	//Mission Related
	CanGoToMission bool `json:"mission_can" db:"mission_can"`
	IsInMission    bool `json:"mission_in" db:"mission_in"`
	MissionType    int  `json:"mission_type" db:"mission_type"`
	MissionMsgLeft int  `json:"mission_msgleft" db:"mission_msgleft"`
	//Hunt Related
	CanGoHunt   bool `json:"hunt_can" db:"hunt_can"`
	IsInHunt    bool `json:"hunt_in" db:"hunt_in"`
	HuntType    int  `json:"hunt_type" db:"hunt_type"`
	HuntMsgLeft int  `json:"hunt_msgleft" db:"hunt_msgleft"`
	//Dungeon Related
	LastDungeon int64 `json:"last_dungeon" db:"last_dungeon"`
	//These will be populated by default :
	Characters     []storedCharacter `json:"chars" db:"-"`
	Grimms         []storedGrimm     `json:"grimms" db:"-"`
	Favorites      []storedCharacter `json:"favorites" db:"-"`
	FavoriteGrimms []storedGrimm     `json:"favorites_grimms" db:"-"`
	OutgoingTrade  int               `json:"outgoingtrade" db:"-"`
	IncomingTrade  int               `json:"incomingtrade" db:"-"`
	MarketListings []string          `json:"marketlistings" db:"-"`
	MarketAuctions []string          `json:"marketauctions" db:"-"`
	LastLootbox    lootBoxTime       `json:"lastlootbox" db:"-"`
	Gamble         gambleTime        `json:"gamble" db:"-"`
	Sort           playerSort        `json:"sort" db:"-"`
	Shop           shop              `json:"shop" db:"-"`
	CharInMission  storedCharacter   `json:"in_mission" db:"-"`
	GrimmInHunt    storedGrimm       `json:"in_hunt" db:"-"`
	SelectedChar   storedCharacter   `json:"selected_char" db:"-"`
	SelectedGrimm  storedGrimm       `json:"selected_grimm" db:"-"`
	BiddedMoney    int64             `json:"bidded_money"`
	//Limited edition lootbox
	LimitedLootBox  []limitedEditionLootBox `json:"limlootbox,omitnil" db:"-"`
	LimitedGrimmBox []limitedEditionLootBox `json:"limgrimmbox,omitnil" db:"-"`
	//Special edition lootbox
	SpecialLootBox []specialEditionLootBox `json:"spelootbox,omitnil" db:"-"`
	//ANTI SPAM PREVENTION
	LastCmd int64 `json:"last_cmd" db:"last_cmd"`
	LastMsg int64 `json:"last_msg" db:"last_msg"`
	//Badges counter (not sent when not wanted)
	Counter playerBadgeCounter `json:"-" db:"-"`

	//EVENT
	ClaimedYang bool `json:"claimedyang" db:"claimedyang"`
}

func getPlayerInv(id string) (p *player) {
	tx, _ := database.Begin()
	defer tx.AutoRollback()
	/*
		Best way to reorder :
		 - have a separate table with user preferences
		 - query it before doing the whole player infos
		 - use it's info to order user
	*/
	var order playerSort
	err := database.Select(`*`).
		From(`players_orders p`).
		Where("p.user_id = $1", id).
		Limit(1).
		QueryStruct(&order)
	if err != nil || order.Sort == "" {
		order.Sort = "c.id asc"
	}
	/*
		Best way to get the right selected :
		 - ask selected type before
		 - choose the right one
	*/
	var selectedType int
	database.Select(`selected_type`).
		From(`players p`).
		Where("p.id = $1", id).
		Limit(1).
		QueryScalar(&selectedType)
	selectedString := "(select row_to_json(a) chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.id = p.selected_id and c.user_id = p.id) a) selected_char,"
	if selectedType == 1 {
		selectedString = "(select row_to_json(a) grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.id = p.selected_id and c.user_id = p.id) a) selected_grimm,"
	}
	d, _ := tx.Select(`p.*, 
			row_to_json(s) as shop, 
			row_to_json(l) as lastlootbox, 
			row_to_json(g) as gamble, 
			row_to_json(o) as sort, 
			(select ('[' || array_to_string(array_agg(row_to_json(a)), ',') || ']')::json chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.in_mission = false and c.user_id = p.id order by `+order.Sort+`, c.id) a) chars,
			(select ('[' || array_to_string(array_agg(row_to_json(a)), ',') || ']')::json grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.in_hunt = false and c.user_id = p.id order by `+order.Sort+`, c.id) a) grimms,  
			(select ('[' || array_to_string(array_agg(row_to_json(a)), ',') || ']')::json chars from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.is_in_favorites and c.user_id = p.id order by `+order.Sort+`, c.id) a) favorites,  
			(select ('[' || array_to_string(array_agg(row_to_json(a)), ',') || ']')::json grimms from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.is_in_favorites and c.user_id = p.id order by `+order.Sort+`, c.id) a) favorites_grimms,
			(select row_to_json(a) in_mission from (select c.*,row_to_json(st) stats from characters c join characters_stats st on st.char_id = c.id where c.in_mission and c.user_id = p.id order by `+order.Sort+` limit 1) a) in_mission,
			(select row_to_json(a) in_hunt from (select c.*,row_to_json(st) stats from grimms c join grimms_stats st on st.grimm_id = c.id where c.in_hunt and c.user_id = p.id order by `+order.Sort+` limit 1) a) in_hunt,
			`+selectedString+`
			('[' || array_to_string(array_agg(row_to_json(li)), ',') || ']')::json limlootbox,
			('[' || array_to_string(array_agg(row_to_json(gri)), ',') || ']')::json limgrimmbox,
			('[' || array_to_string(array_agg(row_to_json(spe)), ',') || ']')::json spelootbox,
			(select array_to_json(array_agg(m.id)) from market m where m.seller_id = p.id) marketlistings,
			(select array_to_json(array_agg(a.id)) from auctions a where a.seller_id = p.id or a.id in (select b.auction_id from auctions_bidder b where b.user_id = p.id)) marketauctions,
			(select sum(b.price) from auctions_bidder b where b.user_id = p.id) bidded_money`).
		From(`players p 
			join player_shops s on s.user_id = p.id
			join player_lastlootboxes l on l.user_id = p.id
			join players_casinos g on g.user_id = p.id
			join players_orders o on o.user_id = p.id
			left join players_limitedlootboxes li on li.user_id = p.id
			left join players_limitedgrimmboxes gri on gri.user_id = p.id
			left join players_speciallootboxes spe on spe.user_id = p.id`).
		Where("p.id = $1", id).
		GroupBy("p.id, s.*, l.*, g.*, o.*").
		Limit(1).
		QueryJSON()
	tx.Commit()
	var pl []player
	e := json.Unmarshal(d, &pl)
	if e != nil {
		defaultInventory := player{
			CharLimit:   30,
			Maxlootbox:  3,
			LootBoxLeft: 1,
			Money:       100,
			IsNew:       true,
			ID:          id,
		}
		return &defaultInventory
	}
	p = &pl[0]
	p.ID = id

	return p
}

type inventoryCacheStruct struct {
	p *player
	t int64
}

var inventoryCache map[string]*inventoryCacheStruct

// Cornelian Points
func cpCalc(Level int) int64 {
	return int64(10*int(math.Pow(float64(Level), 1.8)) + 20)
}

func (p *player) charAmount() int {
	var r int
	var d []string
	d = append(d, p.ID)
	d = append(d, p.MarketListings...)
	d = append(d, p.MarketAuctions...)
	database.Select(`count(*)`).
		From(`characters c`).
		Where("c.user_id in $1", d).
		Limit(1).
		QueryScalar(&r)
	return r
}

func (p *player) grimmAmount() int {
	var r int
	var d []string
	d = append(d, p.ID)
	d = append(d, p.MarketListings...)
	d = append(d, p.MarketAuctions...)
	database.Select(`count(*)`).
		From(`grimms c`).
		Where("c.user_id in $1", d).
		Limit(1).
		QueryScalar(&r)
	return r
}

func (p *player) maxChar() int {
	return p.CharLimit + p.Shop.Extenders
}

func (p *player) money() int64 {
	return p.Money - p.BiddedMoney
}

func (p *player) lootbox() (s int) {
	s += p.LootBoxLeft
	s += p.RareLootBoxLeft
	s += p.GrimmBoxLeft
	s += p.RareGrimmBoxLeft
	s += len(p.LimitedLootBox)
	s += len(p.LimitedGrimmBox)
	s += len(p.SpecialLootBox)
	return
}

func (p *player) giveXPCalc(Level, multiplier int) int64 {
	if Level >= 500 {
		return 0
	}
	rand.Seed(time.Now().UTC().UnixNano())
	rint := multiplier * (Level)
	add := int(float64((rand.Intn(26+rint))+15) * (math.Pow(float64(Level), 0.72) + 1))
	if p.Shop.XPBoost {
		rint = ((3 / 2) * multiplier) * (Level)
		add = int(float64((rand.Intn(33+rint))+25) * (math.Pow(float64(Level), 0.84) + 1))
	}
	return int64(add)
}

func (p *player) giveCPCalc(Level int, difficulty float64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	rint := int(5*difficulty*math.Pow(float64(Level), 1.48)) + 10
	add := difficulty*float64(rand.Intn(rint)) + 5 + math.Pow(float64(Level), 1.45)
	/*
		if p.Shop.XPBoost {
			rint = ((3 / 2) * multiplier) * (Level)
			add = int(float64((rand.Intn(33+rint))+25) * (math.Pow(float64(Level), 0.84) + 1))
			p.Shop.XPBoostTime--
		}
	*/
	return int64(add)
}

func (p *player) delete() {
	tx, _ := database.Begin()
	defer tx.AutoRollback()
	tx.DeleteFrom("market").
		Where("seller_id = $1", p.ID).
		Exec()
	for _, m := range p.MarketListings {
		tx.DeleteFrom("characters_stats s").
			Where("(select c.user_id from characters c where c.id = s.char_id) = $1", m).
			Exec()
		tx.DeleteFrom("characters").
			Where("user_id = $1", m).
			Exec()
		tx.DeleteFrom("grimms_stats s").
			Where("(select c.user_id from grimms c where c.id = s.grimm_id) = $1", m).
			Exec()
		tx.DeleteFrom("grimms").
			Where("user_id = $1", m).
			Exec()
	}
	tx.DeleteFrom("auctions").
		Where("seller_id = $1", p.ID).
		Exec()
	for _, a := range p.MarketAuctions {
		tx.DeleteFrom("characters_stats s").
			Where("(select c.user_id from characters c where c.id = s.char_id) = $1", a).
			Exec()
		tx.DeleteFrom("characters").
			Where("user_id = $1", a).
			Exec()
		tx.DeleteFrom("grimms_stats s").
			Where("(select c.user_id from grimms c where c.id = s.grimm_id) = $1", a).
			Exec()
		tx.DeleteFrom("grimms").
			Where("user_id = $1", a).
			Exec()
		tx.DeleteFrom("auctions_bidder").
			Where("auction_id = $1", a).
			Exec()
	}
	tx.DeleteFrom("characters_stats s").
		Where("(select c.user_id from characters c where c.id = s.char_id) = $1", p.ID).
		Exec()
	tx.DeleteFrom("characters").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("grimms_stats s").
		Where("(select c.user_id from grimms c where c.id = s.grimm_id) = $1", p.ID).
		Exec()
	tx.DeleteFrom("grimms").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("trades").
		Where("user_id = $1 or target_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("player_lastlootboxes").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("player_shops").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players_casinos").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players_orders").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players_limitedlootboxes").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players_limitedgrimmboxes").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players_speciallootboxes").
		Where("user_id = $1", p.ID).
		Exec()
	tx.DeleteFrom("players").
		Where("id = $1", p.ID).
		Exec()
	tx.DeleteFrom("event_signed").
		Where("id = $1", p.ID).
		Exec()
	tx.DeleteFrom("badges").
		Where("id = $1", p.ID).
		Exec()
}

func (p *player) getLatestChar() (char storedCharacter, index int) {
	c := []storedCharacter{}
	d, _ := database.Select(`c.*,row_to_json(st) stats`).
		From(`characters c 
		join characters_stats st on st.char_id = c.id `).
		Where("c.in_mission = false and c.user_id = $1", p.ID).
		OrderBy("owned_at desc").
		Limit(1).
		QueryJSON()
	json.Unmarshal(d, &c)
	char = c[0]
	for i, char := range p.Characters {
		if char.CustomID == c[0].CustomID {
			index = i
			break
		}
	}
	return
}

func (p *player) getLatestGrimm() (grimm storedGrimm, index int) {
	c := []storedGrimm{}
	d, _ := database.Select(`c.*,row_to_json(st) stats`).
		From(`grimms c 
		join grimms_stats st on st.grimm_id = c.id `).
		Where("c.in_hunt = false and c.user_id = $1", p.ID).
		OrderBy("owned_at desc").
		Limit(1).
		QueryJSON()
	json.Unmarshal(d, &c)
	grimm = c[0]
	for i, grimm := range p.Grimms {
		if grimm.CustomID == c[0].CustomID {
			index = i
			break
		}
	}
	return
}

// canDropLootBox will send two bool : canHe and reset
func (p *player) canDropLootBox() (canHe bool, reset bool) {
	lastTime := time.Unix(p.LastLootbox.Time, 0)
	if p.LastLootbox.Amount < p.Maxlootbox {
		canHe = true
		reset = false
		return
	} else if time.Now().Sub(lastTime).Hours() > 24 && p.LastLootbox.Amount == p.Maxlootbox {
		canHe = true
		reset = true
		return
	}
	return
}

// canGamble will send two bool : canHe and reset
func (p *player) canGamble() (canHe, reset bool) {
	lastTime := time.Unix(p.Gamble.Time, 0)
	if p.Gamble.Amount < 3 {
		canHe = true
		reset = false
		return
	} else if time.Now().Sub(lastTime).Hours() > 24 && p.Gamble.Amount >= 3 {
		canHe = true
		reset = true
		return
	}
	return
}

// canDungeon will check if the player can do a dungeon
func (p *player) canDungeon() bool {
	if time.Now().Unix()-p.LastDungeon > 18000 {
		return true
	}
	return false
}

// selectedLevel will return the level of the selected character/grimm
func (p *player) selectedLevel() int {
	if p.SelectedType == 0 {
		return p.SelectedChar.Level
	}
	return p.SelectedGrimm.Level
}
