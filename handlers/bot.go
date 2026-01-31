package handlers

import (
	tele "gopkg.in/telebot.v4"
)

type BotHandlers struct{}

func newBotHandlers() *BotHandlers {
	return &BotHandlers{}
}

func (h *BotHandlers) OnStart(c tele.Context) error {
	return c.Send("Welcome to the Words Bot! Use /define <word> to get the definition.")
}
