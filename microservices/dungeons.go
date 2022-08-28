package microservices

import "github.com/bwmarrin/discordgo"

type DungeonReward struct {
	// Amount of money to be rewarded
	Lien    int `json:"liens"`
	CCBox   int `json:"ccBox"`
	Arms    int `json:"arms"`
	Minions int `json:"minions"`
}

type DungeonCreateRequest struct {
	ID     string
	UserID string
}

type DungeonEndResponse struct {
	Win     bool
	Rewards *DungeonReward
}

type DungeonsMessage struct {
	UserID  string
	Message *discordgo.MessageEmbed
}
