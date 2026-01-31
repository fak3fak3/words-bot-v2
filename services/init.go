package services

import (
	"words/apis"
	"words/config"
	"words/handlers"
	"words/repos"
)

type Services struct {
	r *repos.Repos

	WordsService *WordsService
	BotService   *BotService
}

func Init(r *repos.Repos, cfg *config.Config, apis *apis.APIs, handlers *handlers.Handlers) (*Services, error) {
	return &Services{
		r: r,

		WordsService: newWordsService(r, apis),
		BotService:   newBotService(cfg, handlers),
	}, nil
}
