package shorten_url

import (
	"context"
	"errors"

	ports "github.com/AndresCRamos/url-shortener/internal/app/ports"
	"github.com/AndresCRamos/url-shortener/internal/domain/shorten_link/model"
)

var (
	counter      = 1000
	allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-."
)

func setCounterForTests(value int) {
	counter = value
}

type URLShortenerService struct {
	repo ports.ShortenURLRepository
}

func NewURLShortenerService(repo ports.ShortenURLRepository) *URLShortenerService {
	return &URLShortenerService{
		repo: repo,
	}
}

func (s *URLShortenerService) CreateShortenURL(ctx context.Context, originalURL string) (*model.ShortenURLModel, error) {
	if originalURL == "" {
		return &model.ShortenURLModel{}, errors.New("original URL cannot be empty")
	}
	shortURL := generateShortURL()
	urlModel := &model.ShortenURLModel{
		Original: originalURL,
		Short:    shortURL,
	}

	shortenedURL, err := s.repo.Save(ctx, urlModel)

	if err != nil {
		return &model.ShortenURLModel{}, err
	}
	return shortenedURL, nil

}

func (s *URLShortenerService) GetOriginalURL(ctx context.Context, shortURL string) (*model.ShortenURLModel, error) {
	if shortURL == "" {
		return &model.ShortenURLModel{}, errors.New("short URL cannot be empty")
	}
	urlModel, err := s.repo.FindByShort(ctx, shortURL)
	if err != nil {
		return &model.ShortenURLModel{}, err
	}
	err = s.repo.IncrementVisitCount(ctx, shortURL)
	if err != nil {
		return &model.ShortenURLModel{}, err
	}
	return urlModel, nil
}

func generateShortURL() string {
	// Dummy implementation for URL shortening
	charPositions := []int{}
	dividend := counter
	for {
		if dividend < len(allowedChars) {
			charPositions = append(charPositions, dividend)
			break
		}
		remainder := dividend % len(allowedChars)
		charPositions = append(charPositions, remainder)
		dividend = dividend / len(allowedChars)
	}
	characters := ""
	for _, charPosition := range charPositions {
		characters += allowedChars[charPosition : charPosition+1]
	}
	counter++
	return characters
}
