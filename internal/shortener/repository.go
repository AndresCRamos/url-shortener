package shortener

import (
	"context"

	model "github.com/AndresCRamos/url-shortener/internal/model"
)

type ShortenURLRepository interface {
	Save(ctx context.Context, url *model.ShortenURLModel) (*model.ShortenURLModel, error)
	FindByShort(ctx context.Context, short string) (*model.ShortenURLModel, error)
}
