package handlers

type Handlers struct {
	Bot *BotHandlers
}

func NewHandlers() *Handlers {
	return &Handlers{
		Bot: newBotHandlers(),
	}
}
