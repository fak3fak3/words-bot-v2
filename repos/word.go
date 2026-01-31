package repos

import (
	"words/db"
	"words/types"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

type WordsRepo struct {
	db *db.DB
}

func NewWordsRepo(db *db.DB) *WordsRepo {
	return &WordsRepo{db: db}
}

func (r *WordsRepo) GetWordDetailsByWordAndLang(wordStr string, language string) (*types.WordResponse, error) {
	var word types.Word
	if err := r.db.Postgres.Where("word = ? AND language = ?", wordStr, language).Preload("Definitions").First(&word).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			spew.Dump("Word not found in DB: %s (%s)", wordStr, language)
			return nil, nil
		}
	}

	definitions := make([]types.WordDefinitionResponse, len(word.Definitions))
	for i, defModel := range word.Definitions {
		definitions[i] = types.WordDefinitionResponse{
			PartOfSpeech: defModel.PartOfSpeech,
			Meaning:      defModel.Meaning,
			Example:      defModel.Example,
		}
	}

	wordDetails := &types.WordResponse{
		Word:          word.Word,
		Language:      word.Language,
		Transcription: word.Transcription,
		Definitions:   definitions,
	}

	return wordDetails, nil
}

func (r *WordsRepo) SaveWordDetails(wordResp *types.WordResponse) error {
	wordModel := types.Word{
		Word:          wordResp.Word,
		Language:      wordResp.Language,
		Transcription: wordResp.Transcription,
	}

	if err := r.db.Postgres.Create(&wordModel).Error; err != nil {
		return err
	}

	for _, defResp := range wordResp.Definitions {
		defModel := types.WordDefinition{
			WordID:       wordModel.ID,
			PartOfSpeech: defResp.PartOfSpeech,
			Meaning:      defResp.Meaning,
			Example:      defResp.Example,
		}
		if err := r.db.Postgres.Create(&defModel).Error; err != nil {
			return err
		}
	}

	return nil
}
