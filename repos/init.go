package repos

import (
	"words/db"
	"words/types"
)

type WordsRepoInterface interface {
	GetWordDetailsByWordAndLang(word string, language string) (*types.WordResponse, error)
	SaveWordDetails(wordResp *types.WordResponse) error
}

type Repos struct {
	db        *db.DB
	WordsRepo WordsRepoInterface
}

func Init(db *db.DB) (*Repos, error) {
	wordsRepo := newWordsRepo(db)

	return &Repos{
		db:        db,
		WordsRepo: wordsRepo,
	}, nil
}
