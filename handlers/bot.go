package handlers

import (
	"bytes"
	"words/services"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/telebot.v4"
	tele "gopkg.in/telebot.v4"
)

type BotHandlers struct {
	s *services.Services
}

func newBotHandlers(s *services.Services) *BotHandlers {
	return &BotHandlers{
		s: s,
	}
}

func (h *BotHandlers) OnStart(c tele.Context) error {
	spew.Dump(c.Sender())
	return c.Send("Welcome to the Words Bot! Use /define <word> to get the definition.")
}

func (h *BotHandlers) OnText(c tele.Context) error {
	word, err := h.s.WordsService.CreateWordDefinition(c.Text(), "en")
	if err != nil {
		return c.Send("Error creating word definition: " + err.Error())
	}

	message := h.s.MessagesService.CreateWordDefinitionMessage(word)

	imageBytes := h.s.ImagesService.CreateWordImage(word.Word, "white")
	image := &telebot.Photo{
		File:    telebot.FromReader(bytes.NewReader(imageBytes)),
		Caption: message,
	}

	return c.Send(image, &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdownV2,
	})
}
