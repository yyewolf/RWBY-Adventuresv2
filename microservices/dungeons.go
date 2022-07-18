package microservices

type DungeonReward struct {
	// Amount of money to be rewarded
	Lien  int `json:"liens"`
	CCBox int `json:"ccBox"`
	Arms  int `json:"arms"`
}

type DungeonCreateRequest struct {
	ID string
}

type DungeonEndResponse struct {
	Rewards *DungeonReward
}
