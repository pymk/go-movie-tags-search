package main

import (
	"fmt"
	"log"
	"os"

	"go-movie-tags-search/internal/config"
	"go-movie-tags-search/internal/db"
	"go-movie-tags-search/internal/display"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	database, err := db.Connect("./data/movies.duckdb")
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}
	defer database.Close()

	if cfg.ShowTags {
		tags, err := db.GetTopTags(database, cfg.Limit)
		if err != nil {
			log.Fatalf("could not get display tags: %v", err)
		}
		display.Tags(tags)
	} else {
		movies, err := db.SearchMovies(database, cfg.Tags)
		if err != nil {
			log.Fatalf("could not search movies: %v", err)
		}
		display.Movies(movies, cfg.Limit)
	}
}
