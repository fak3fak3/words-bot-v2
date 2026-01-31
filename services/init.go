package services

import "words/repos"

type Services struct {
	r *repos.Repos
}

func Init(r *repos.Repos) (*Services, error) {
	return &Services{
		r: r,
	}, nil
}
