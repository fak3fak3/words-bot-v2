package services

import (
	"words/config"
	"words/repos"
)

type Services struct {
	r *repos.Repos

	ChatComplition *ChatComplitionService
}

func Init(r *repos.Repos, cfg *config.Config) (*Services, error) {
	chat := NewChatComplitionService(r, cfg)

	return &Services{
		r: r,

		ChatComplition: chat,
	}, nil
}
