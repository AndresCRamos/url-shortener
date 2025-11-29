package main

import (
	"fmt"

	config "github.com/AndresCRamos/url-shortener/internal/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
}
