package main

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/models"

	uuid "github.com/satori/go.uuid"
)

func charConverter(c storedCharacter) *models.Character {
	char := &models.Character{
		CharID:        uuid.NewV4().String(),
		UserID:        c.UserID,
		Name:          c.Name,
		Level:         c.Level,
		Rarity:        c.Rarity,
		Buffs:         c.Buffs,
		IsInFavorites: c.IsInFavorites,
		InMission:     c.InMission,
		XP:            c.XP,
		XPCap:         c.XPCap,
		Value:         c.Stats.Value,
	}
	char.CalcStats()
	char.Stats.CharID = char.CharID
	return char
}

func grimmConverter(c storedGrimm) *models.Grimm {
	grimm := &models.Grimm{
		GrimmID:       uuid.NewV4().String(),
		UserID:        c.UserID,
		Name:          c.Name,
		Level:         c.Level,
		Rarity:        c.Rarity,
		Buffs:         c.Buffs,
		IsInFavorites: c.IsInFavorites,
		InHunt:        c.InHunt,
		XP:            c.XP,
		XPCap:         c.XPCap,
		Value:         c.Stats.Value,
	}
	grimm.CalcStats()
	grimm.Stats.GrimmID = grimm.GrimmID
	return grimm
}

func playerConverter(p *player) *models.Player {
	player := &models.Player{
		DiscordID:     p.ID,
		IsNew:         p.IsNew,
		Balance:       p.Money,
		BiddedBalance: p.BiddedMoney,
		Level:         int64(p.Level),
		CP:            p.CP,
		MaxCP:         p.MaxCP,
		CharLimit:     p.CharLimit,
		Maxlootbox:    p.Maxlootbox,
		SelectedID:    p.SelectedID,
		SelectedType:  p.SelectedType,
		Disabled:      p.Disabled,
		Arms:          p.ArmsLeft,
		Minions:       p.MinionsLeft,
		Jar:           p.XPJar,
	}
	player.Save()

	player.Settings = &models.PlayerSettings{
		DiscordID: p.ID,
	}
	player.Settings.Save()

	player.Missions = &models.PlayerMission{
		DiscordID:      p.ID,
		CanGoToMission: p.CanGoToMission,
		IsInMission:    p.IsInMission,
		MissionType:    p.MissionType,
		MissionMsgLeft: p.MissionMsgLeft,

		CanGoHunt:   p.CanGoHunt,
		IsInHunt:    p.IsInHunt,
		HuntType:    p.HuntType,
		HuntMsgLeft: p.HuntMsgLeft,
	}
	player.Missions.Save()

	player.Status = &models.PlayerStatus{
		DiscordID:   p.ID,
		LastXP:      p.LastXP,
		Voted:       p.Voted,
		DailyStreak: p.DailyStreak,
		LastOpening: p.LastOpening,
		LastReport:  p.LastReport,
		LastDungeon: p.LastDungeon,
	}
	player.Status.Save()

	player.Shop = &models.PlayerShop{
		DiscordID: p.ID,

		XPBoost:       p.Shop.XPBoost,
		XPBoostTime:   p.Shop.XPBoostTime,
		LuckBoost:     p.Shop.LuckBoost,
		LuckBoostTime: p.Shop.LuckBoostTime,

		Extensions: p.Shop.Extenders / 20,
	}
	player.Shop.Save()

	player.LastBoxes = &models.PlayerLootTime{
		DiscordID: p.ID,
	}
	player.LastBoxes.Save()

	player.Gamble = &models.PlayerGamble{
		DiscordID: p.ID,
		Amount:    p.Gamble.Amount,
		Time:      p.Gamble.Time,
	}
	player.Gamble.Save()

	player.Boxes = &models.PlayerBoxes{
		DiscordID: p.ID,

		Boxes:          p.LootBoxLeft,
		RareBoxes:      p.RareLootBoxLeft,
		GrimmBoxes:     p.GrimmBoxLeft,
		RareGrimmBoxes: p.RareGrimmBoxLeft,
	}
	player.Boxes.Save()

	player.Stats = &models.PlayerStats{
		DiscordID: p.ID,

		ArenasCompleted: p.Counter.ArenasCompleted,
		BattlesWon:      p.Counter.BattlesWon,
		BattlesLost:     p.Counter.BattlesLost,
		LootboxOpened:   p.Counter.LootboxOpened,
		RoleplaySent:    p.Counter.RoleplaySent,
		DungeonDone:     p.Counter.DungeonDone,
		MarketSold:      p.Counter.MarketSold,
		MarketBought:    p.Counter.MarketBought,
	}
	player.Stats.Save()

	if p.IsInMission {
		player.CharInMission = charConverter(p.CharInMission)
		config.Database.Create(player.CharInMission)
	}

	if p.IsInHunt {
		player.GrimmInHunt = grimmConverter(p.GrimmInHunt)
		config.Database.Create(player.GrimmInHunt)
	}

	player.Market = &models.PlayerMarket{
		DiscordID: p.ID,
	}

	for _, lID := range p.MarketListings {
		listing, err := getListing(lID)
		if err != nil {
			fmt.Println(err)
		}
		newID := uuid.NewV4().String()
		l := &models.Listing{
			ID: newID,

			SellerID:   listing.SellerID,
			SellerName: listing.SellerName,
			Note:       listing.Note,
			Price:      int64(listing.Price),
			Type:       listing.Type,
		}
		if l.Type == models.CharType {
			l.Char = charConverter(listing.Char)
			l.Char.UserID = newID
			config.Database.Create(l.Char)
		} else {
			l.Grimm = grimmConverter(listing.Grimm)
			l.Grimm.UserID = newID
			config.Database.Create(l.Grimm)
		}

		l.Save()
	}

	for _, aID := range p.MarketAuctions {
		auction, err := getAuction(aID)
		if err != nil {
			fmt.Println(err)
		}
		newID := uuid.NewV4().String()
		a := &models.Auction{
			ID: newID,

			SellerID:       auction.SellerID,
			SellerName:     auction.SellerName,
			StartedAt:      auction.StartedAt,
			EndsAt:         auction.EndsAt,
			TimeExtensions: auction.TimeExtensions,

			Type: auction.Type,
		}
		if a.Type == models.CharType {
			a.Char = charConverter(auction.Char)
			a.Char.UserID = newID
			config.Database.Create(a.Char)
		} else {
			a.Grimm = grimmConverter(auction.Grimm)
			a.Grimm.UserID = newID
			config.Database.Create(a.Grimm)
		}

		a.Save()

		for _, bid := range auction.History {
			b := models.AuctionBidders{
				AuctionID: newID,
				UserID:    bid.ID,
				Bid:       int64(bid.Price),
			}
			b.Save()
		}
	}

	for _, c := range p.Characters {
		if c.CustomID == p.SelectedID {
			char := charConverter(c)
			player.SelectedID = char.CharID
			player.Save()
			config.Database.Create(char)
		} else {
			char := charConverter(c)
			config.Database.Create(char)
		}
	}

	for _, g := range p.Grimms {
		if g.CustomID == p.SelectedID {
			grimm := grimmConverter(g)
			player.SelectedID = grimm.GrimmID
			player.Save()
			config.Database.Create(grimm)
		} else {
			grimm := grimmConverter(g)
			config.Database.Create(grimm)
		}
	}

	for _, b := range badges {
		if p.hasBadge(b.value) {
			for _, badge := range models.DefaultBadges {
				if badge.Name == b.name {
					playerBadge := &models.PlayerBadges{
						DiscordID: player.DiscordID,
						BadgeID:   badge.BadgeID,
						Badge:     badge,
					}
					playerBadge.Save()
				}
			}
		}
	}

	return player
}
