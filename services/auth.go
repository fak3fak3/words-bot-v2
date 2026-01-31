package services

import "words/repos"

type AuthService struct {
	r *repos.Repos
}

func newAuthService(r *repos.Repos) *AuthService {
	return &AuthService{
		r: r,
	}
}
