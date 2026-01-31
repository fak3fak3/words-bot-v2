package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"words/apis"
	"words/repos"
	"words/types"
)

type WordsService struct {
	apis *apis.APIs
	r    *repos.Repos
}

func newWordsService(r *repos.Repos, apis *apis.APIs) *WordsService {
	return &WordsService{
		apis: apis,
		r:    r,
	}
}

func (s *WordsService) CreateWordDefinition(word string, language string) (*types.WordResponse, error) {
	data, err := os.ReadFile("prompts/WORD_ARTICLE.md")
	if err != nil {
		panic(err)
	}

	responseDB, err := s.r.WordsRepo.GetWordDetailsByWordAndLang(word, language)
	if err != nil {
		return nil, fmt.Errorf("failed to get word details from DB: %w", err)
	}
	if responseDB != nil {
		return responseDB, nil
	}

	str := string(data)
	str = strings.ReplaceAll(str, "{{WORD}}", word)
	str = strings.ReplaceAll(str, "{{LANGUAGE}}", language)

	chatResponseBytes, err := s.apis.OpenRouterAPI.GenerateResponse(str)
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %w", err)
	}

	response := types.WordResponse{}

	err = json.Unmarshal([]byte(chatResponseBytes), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	err = s.r.WordsRepo.SaveWordDetails(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to save word details: %w", err)
	}

	return &response, nil
}
