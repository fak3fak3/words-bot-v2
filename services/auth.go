package services

import (
	"fmt"
	"words/repos"
)

type AuthService struct {
	r *repos.Repos
}

func newAuthService(r *repos.Repos) *AuthService {
	return &AuthService{
		r: r,
	}
}

func (s *AuthService) SignUp(telegramID int64, username string, firstName string, lastName string) (bool, error) {
	user, err := s.r.UsersRepo.GetUserByTelegramID(telegramID)
	if err != nil {
		return false, fmt.Errorf("failed to get user by telegram ID: %w", err)
	}
	if user == nil {
		err = s.r.UsersRepo.CreateUser(telegramID, username, firstName, lastName)
		if err != nil {
			return false, fmt.Errorf("failed to create user: %w", err)
		}

		return false, nil
	}
	return true, nil
}
