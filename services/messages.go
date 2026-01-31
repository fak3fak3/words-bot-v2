package services

import (
	"fmt"
	"strings"
	"words/types"
)

type MessagesService struct{}

func newMessagesService() *MessagesService {
	return &MessagesService{}
}

func (s *MessagesService) CreateWordDefinitionMessage(w *types.WordResponse) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("_%s_\n\n", escapeMarkdownV2(w.Transcription)))

	for i, d := range w.Definitions {
		num := escapeMarkdownV2(fmt.Sprintf("%d.", i+1))
		line := fmt.Sprintf("%s *%s* %s\n",
			num,
			escapeMarkdownV2(d.PartOfSpeech),
			escapeMarkdownV2(d.Meaning),
		)
		b.WriteString(line)

		if d.Example != "" {
			b.WriteString(fmt.Sprintf("_%s_\n", escapeMarkdownV2(d.Example)))
		}

		b.WriteString("\n")
	}

	return b.String()
}

func escapeMarkdownV2(s string) string {
	special := "_*[]()~`>#+-=|{}.!" // добавляем точку и тире
	for _, c := range special {
		s = strings.ReplaceAll(s, string(c), `\`+string(c))
	}
	return s
}
