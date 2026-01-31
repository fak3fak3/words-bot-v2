package services

import (
	"context"
	"words/config"
	"words/repos"

	"github.com/davecgh/go-spew/spew"
	openrouter "github.com/revrost/go-openrouter"
)

type ChatComplitionService struct {
	r   *repos.Repos
	cfg *config.Config
}

func NewChatComplitionService(r *repos.Repos, cfg *config.Config) *ChatComplitionService {
	return &ChatComplitionService{
		r:   r,
		cfg: cfg,
	}
}

func (s *ChatComplitionService) GenerateResponse(prompt string) (string, error) {
	client := openrouter.NewClient(
		s.cfg.OpenRouterAPIKey,
		openrouter.WithXTitle("Words Bot"),
		openrouter.WithHTTPReferer("https://myapp.com"),
	)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openrouter.ChatCompletionRequest{
			Model: "nvidia/nemotron-3-nano-30b-a3b:free",
			Messages: []openrouter.ChatCompletionMessage{
				openrouter.UserMessage(prompt),
			},
		},
	)

	if err != nil {
		spew.Dump("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content.Text, nil
}
