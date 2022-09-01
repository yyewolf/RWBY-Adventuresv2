package websocket

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/market_back/microservice"
	"rwby-adventures/microservices"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/yyewolf/gosf"
)

func getLatestAuctions(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Get 10 auctions from cache :
	var auctions []*models.Auction
	var icons []string
	for i, a := range cache.Auctions {
		if i >= 10 {
			break
		}
		a.Ended = a.EndsAt < time.Now().Unix()
		if a.Ended {
			a.End()
			if len(a.Bidders) > 0 {
				var personaString string
				if a.Type == models.CharType {
					personaString = a.Char.FullString()
				} else {
					personaString = a.Grimm.FullString()
				}

				go microservice.SendMessageToBot(&microservices.MarketMessage{
					UserID: a.SellerID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("Your auction on `%s` has ended, you earned : **%d** Liens.", personaString, a.Bidders[0].Bid),
					},
				})

				go microservice.SendMessageToBot(&microservices.MarketMessage{
					UserID: a.Bidders[0].UserID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("You purchased `%s` for **%d** Liens on an auction.", personaString, a.Bidders[0].Bid),
					},
				})
			} else {
				var personaString string
				if a.Type == models.CharType {
					personaString = a.Char.FullString()
				} else {
					personaString = a.Grimm.FullString()
				}

				go microservice.SendMessageToBot(&microservices.MarketMessage{
					UserID: a.SellerID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction Ended",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("Your auction on `%s` has ended, nobody placed a bid.", personaString),
					},
				})
			}
			continue
		}
		auctions = append(auctions, a)
		if a.Type == models.CharType {
			icons = append(icons, a.Char.ToRealChar().IconURL)
		} else {
			icons = append(icons, a.Grimm.ToRealGrimm().IconURL)
		}
	}

	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["auctions"] = auctions
	msg.Body["icons"] = icons

	return msg
}
