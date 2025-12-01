package shortener

import "context"

type ShortenURLRepository interface {
	Save(ctx context.Context, url *ShortenURLModel) (*ShortenURLModel, error)
	FindByShort(ctx context.Context, short string) (*ShortenURLModel, error)
	Exists(ctx context.Context, short string) (bool, error)
}
