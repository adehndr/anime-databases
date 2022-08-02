package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adehndr/anime-databases/app"
	"github.com/adehndr/anime-databases/repository"
)

func main() {
	dbMySql := app.OpenDatabaseConnection()
	defer dbMySql.Close()

	err := dbMySql.Ping()
	if err != nil {
		log.Fatal(err)
	}

	animeListRepository := repository.NewAnimeListRepository(dbMySql)
	data, err := animeListRepository.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
