package apis

import (
	"context"
	"words/config"

	"github.com/davecgh/go-spew/spew"
	openrouter "github.com/revrost/go-openrouter"
)

type OpenRouterAPI struct {
	cfg *config.Config
}

func NewOpenRouterAPI(cfg *config.Config) *OpenRouterAPI {
	return &OpenRouterAPI{
		cfg: cfg,
	}
}

func (s *OpenRouterAPI) GenerateResponse(prompt string) (string, error) {
	client := openrouter.NewClient(
		s.cfg.OpenRouterAPIKey,
		openrouter.WithXTitle("Words Bot"),
		openrouter.WithHTTPReferer("https://myapp.com"),
	)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openrouter.ChatCompletionRequest{
			Model: "arcee-ai/trinity-large-preview:free",
			Messages: []openrouter.ChatCompletionMessage{
				openrouter.UserMessage(prompt),
			},
		},
	)

	if err != nil {
		spew.Dump("ChatCompletion error: %v\n", err)
		return "", err
	}

	respStr := resp.Choices[0].Message.Content.Text

	return respStr, nil
}
