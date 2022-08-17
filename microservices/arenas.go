package microservices

type CreateArena struct {
	ID        string
	ChannelID string
}

type EndArena struct {
	ChannelID string
	Message   string
}
