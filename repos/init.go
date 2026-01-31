package repos

import (
	"words/db"
	"words/types"
)

type WordsRepoInterface interface {
	GetWordDetailsByWordAndLang(word string, language string) (*types.WordResponse, error)
	SaveWordDetails(wordResp *types.WordResponse) error
}

type UsersRepoInterface interface {
	GetUserByTelegramID(telegramID int64) (*types.UserResponse, error)
	CreateUser(telegramID int64, username string, firstName string, lastName string) error
}

type Repos struct {
	db        *db.DB
	WordsRepo WordsRepoInterface
	UsersRepo UsersRepoInterface
}

func Init(db *db.DB) (*Repos, error) {
	wordsRepo := newWordsRepo(db)
	usersRepo := newUsersRepo(db)

	return &Repos{
		db:        db,
		WordsRepo: wordsRepo,
		UsersRepo: usersRepo,
	}, nil
}
