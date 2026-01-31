package main

import (
	"words/apis"
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

	apis, err := apis.Init(cfg)
	if err != nil {
		panic(err)
	}

	s, err := services.Init(repos, cfg, apis)
	if err != nil {
		panic(err)
	}

	response, err := s.WordsService.CreateWordDefinition("cat", "en")
	if err != nil {
		panic(err)
	}
	spew.Dump(response)
}
