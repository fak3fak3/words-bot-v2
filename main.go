package main

import (
	"words/config"
	"words/db"
	"words/repos"
	"words/services"

	"github.com/davecgh/go-spew/spew"
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

	s, err := services.Init(repos, cfg)
	if err != nil {
		panic(err)
	}

	response, err := s.ChatComplition.GenerateResponse("hello mr model")
	if err != nil {
		panic(err)
	}
	spew.Dump(response)
}
