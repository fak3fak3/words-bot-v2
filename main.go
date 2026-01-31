package main

import (
	"words/config"
	"words/db"
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

	_, err = services.Init(repos)
	if err != nil {
		panic(err)
	}

}
