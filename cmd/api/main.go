package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/AndresCRamos/url-shortener/internal/adapter/sqlite"
	service "github.com/AndresCRamos/url-shortener/internal/app/service/shorten_url"
	"github.com/AndresCRamos/url-shortener/internal/config"
	sqlc "github.com/AndresCRamos/url-shortener/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()
	db, err := sql.Open("sqlite3", cfg.DatabaseURL)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	repo := sqlite.NewShortenerSQLiteRepo(queries)
	service := service.NewURLShortenerService(repo)

	ctx := context.Background()

	shortenModel, err := service.CreateShortenURL(ctx, "https://example.com/some/very/long/url4")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Shortened URL: ", shortenModel)
	}

	retrievedModel, err := service.GetOriginalURL(ctx, "Op")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Retrieved Original URL: ", retrievedModel)
	}

}
