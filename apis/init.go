package apis

import "words/config"

type APIs struct {
	cfg *config.Config

	OpenRouterAPI *OpenRouterAPI
}

func Init(cfg *config.Config) (*APIs, error) {
	openRouterAPI := NewOpenRouterAPI(cfg)

	return &APIs{
		cfg: cfg,

		OpenRouterAPI: openRouterAPI,
	}, nil
}
