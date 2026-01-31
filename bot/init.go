package bot

import (
	"fmt"
	"os"
	"time"
	"words/config"
	"words/handlers"

	tele "gopkg.in/telebot.v4"
)

type Bot struct {
	cfg *config.Config
	b   *tele.Bot
}

func Create(cfg *config.Config) *Bot {
	return &Bot{
		cfg: cfg,
	}
}

func (s *Bot) Init(h *handlers.Handlers) error {
	pref := tele.Settings{
		Token:  os.Getenv("TG_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		fmt.Errorf("failed to create bot: %w", err)
		return err
	}
	s.b = b

	s.b.Handle("/start", h.Bot.OnStart)
	s.b.Handle(tele.OnText, h.Bot.OnText)

	return nil
}

func (s *Bot) Start() {
	s.b.Start()
}
