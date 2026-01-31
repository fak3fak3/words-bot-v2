package services

import (
	"words/apis"
	"words/config"
	"words/repos"
)

type Services struct {
	r *repos.Repos

	WordsService    *WordsService
	AuthService     *AuthService
	ImagesService   *ImagesService
	MessagesService *MessagesService
}

func Init(r *repos.Repos, cfg *config.Config, apis *apis.APIs) (*Services, error) {
	return &Services{
		r: r,

		WordsService:    newWordsService(r, apis),
		AuthService:     newAuthService(r),
		ImagesService:   newImagesService(),
		MessagesService: newMessagesService(),
	}, nil
}
