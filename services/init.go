package services

import (
	"words/apis"
	"words/config"
	"words/repos"
)

type Services struct {
	r *repos.Repos

	WordsService *WordsService
}

func Init(r *repos.Repos, cfg *config.Config, apis *apis.APIs) (*Services, error) {
	wordsService := NewWordsService(r, apis)

	return &Services{
		r: r,

		WordsService: wordsService,
	}, nil
}
