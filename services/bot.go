package services

import (
	"fmt"
	"os"
	"time"
	"words/config"
	"words/handlers"

	tele "gopkg.in/telebot.v4"
)

type BotService struct {
	cfg *config.Config
	h   *handlers.Handlers
}

func newBotService(cfg *config.Config, h *handlers.Handlers) *BotService {
	return &BotService{
		cfg: cfg,
		h:   h,
	}
}

func (s *BotService) Start() {
	pref := tele.Settings{
		Token:  os.Getenv("TG_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		fmt.Errorf("failed to create bot: %w", err)
		return
	}

	b.Handle("/start", s.h.Bot.OnStart)

	b.Start()
}
