package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndresCRamos/url-shortener/internal/adapter/sqlite"
	"github.com/AndresCRamos/url-shortener/internal/config"
	sqlc "github.com/AndresCRamos/url-shortener/internal/db"
	"github.com/AndresCRamos/url-shortener/internal/model"
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

	ctx := context.Background()
	code := "abc123"
	created, err := repo.Save(ctx, &model.ShortenURLModel{
		Short:    code,
		Original: "https://example.com",
	})
	fmt.Println(created, err)
	err = repo.IncrementVisitCount(ctx, code)
	fmt.Println("Incremented views:", err)

	search, err := repo.FindByShort(ctx, code)
	fmt.Println(search, err)

}
