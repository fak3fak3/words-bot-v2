package handlers

import "words/services"

type Handlers struct {
	Bot *BotHandlers
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		Bot: newBotHandlers(s),
	}
}
