package repos

import (
	"words/db"
	"words/types"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *db.DB
}

func newUsersRepo(db *db.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) GetUserByTelegramID(telegramID int64) (*types.UserResponse, error) {
	var user types.User
	err := r.db.Postgres.Where("telegram_id = ?", telegramID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	userResponse := types.UserResponse{
		TelegramID: user.TelegramID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Username:   user.Username,
	}

	return &userResponse, nil
}

func (r *UsersRepo) CreateUser(telegramID int64, username string, firstName string, lastName string) error {
	user := types.User{
		TelegramID: telegramID,
		FirstName:  firstName,
		LastName:   lastName,
		Username:   username,
	}
	return r.db.Postgres.Create(&user).Error
}
