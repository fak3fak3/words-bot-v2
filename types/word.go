package types

import "gorm.io/gorm"

type WordResponse struct {
	Word          string                   `json:"word"`
	Language      string                   `json:"language"`
	Transcription string                   `json:"transcription"`
	Definitions   []WordDefinitionResponse `json:"definitions"`
}

type WordDefinitionResponse struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Meaning      string `json:"meaning"`
	Example      string `json:"example"`
}

type Word struct {
	gorm.Model

	Word          string `gorm:"index:idx_word_language,unique"`
	Language      string `gorm:"index:idx_word_language,unique"`
	Transcription string
	Definitions   []WordDefinition `gorm:"foreignKey:WordID;constraint:OnDelete:CASCADE;"`
}

type WordDefinition struct {
	gorm.Model

	WordID       uint
	PartOfSpeech string
	Meaning      string
	Example      string
}
