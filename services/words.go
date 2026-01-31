package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"words/apis"
	"words/repos"
)

type WordsService struct {
	apis *apis.APIs
	r    *repos.Repos
}

func NewWordsService(r *repos.Repos, apis *apis.APIs) *WordsService {
	return &WordsService{
		apis: apis,
		r:    r,
	}
}

func (s *WordsService) CreateWordDefinition(word string, language string) (*repos.Word, error) {
	data, err := os.ReadFile("prompts/WORD_ARTICLE.md")
	if err != nil {
		panic(err)
	}

	str := string(data)
	str = strings.ReplaceAll(str, "{{WORD}}", word)
	str = strings.ReplaceAll(str, "{{LANGUAGE}}", language)

	chatResponseBytes, err := s.apis.OpenRouterAPI.GenerateResponse(str)
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %w", err)
	}

	response := repos.Word{}

	err = json.Unmarshal([]byte(chatResponseBytes), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
