package microservice

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/microservices"
	"rwby-adventures/models"

	"github.com/bwmarrin/discordgo"
)

func VerifyAuctions() {
	var keep []*models.Auction
	for _, a := range cache.Auctions {
		if a.Ended {
			var personaString string
			if a.Type == models.CharType {
				personaString = a.Char.FullString()
			} else {
				personaString = a.Grimm.FullString()
			}

			if len(a.Bidders) > 0 {
				bidder := a.Bidders[len(a.Bidders)-1]
				// Buyer's message
				go SendMessageToBot(&microservices.MarketMessage{
					UserID: bidder.UserID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction Purchase",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("You have successfully won the auction of `%s` for **%d** Liens.", personaString, bidder.Bid),
					},
				})

				// Seller's message
				go SendMessageToBot(&microservices.MarketMessage{
					UserID: a.SellerID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction Sold",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("You sold `%s` for **%d** Liens in an auction.", personaString, bidder.Bid),
					},
				})
			} else {

				// Seller's message
				go SendMessageToBot(&microservices.MarketMessage{
					UserID: a.SellerID,
					Message: &discordgo.MessageEmbed{
						Title:       "Auction End",
						Color:       config.Botcolor,
						Description: fmt.Sprintf("The auction for `%s` has ended.", personaString),
					},
				})
			}
		} else {
			keep = append(keep, a)
		}
	}
}
