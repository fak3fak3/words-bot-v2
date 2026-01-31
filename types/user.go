package types

import "gorm.io/gorm"

type User struct {
	gorm.Model

	TelegramID int64  `gorm:"uniqueIndex;not null"`
	FirstName  string `gorm:"not null"`
	LastName   string
	Username   string
	Words      []Word `gorm:"many2many:user_words;constraint:OnDelete:CASCADE;"`
}

type UserResponse struct {
	TelegramID int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name,omitempty"`
	Username   string `json:"username,omitempty"`

	Words []WordResponse `json:"words,omitempty"`
}
