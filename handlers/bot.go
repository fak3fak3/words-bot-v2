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
	ok, err := h.s.AuthService.SignUp(c.Sender().ID, c.Sender().Username, c.Sender().FirstName, c.Sender().LastName)

	spew.Dump(ok)

	if err != nil {
		spew.Dump("SignUp error: %v\n", err)
		return c.Send("Error during sign up: " + err.Error())
	}
	if !ok {
		return c.Send("Welcome to the Words Bot! Send me a word and I'll provide you with its definition and an image.")
	}
	return c.Send("Welcome back to the Words Bot!")
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
