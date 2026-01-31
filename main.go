package main

import (
	"words/apis"
	"words/bot"
	"words/config"
	"words/db"
	"words/handlers"
	"words/repos"
	"words/services"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	db, err := db.Init(cfg)
	if err != nil {
		panic(err)
	}

	repos, err := repos.Init(db)
	if err != nil {
		panic(err)
	}

	apis, err := apis.Init(cfg)
	if err != nil {
		panic(err)
	}

	s, err := services.Init(repos, cfg, apis)
	if err != nil {
		panic(err)
	}

	handlers := handlers.NewHandlers(s)

	bot := bot.Create(cfg)

	bot.Init(handlers)
	bot.Start()
}
