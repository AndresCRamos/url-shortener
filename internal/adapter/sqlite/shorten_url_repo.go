package sqlite

import (
	"context"

	"github.com/AndresCRamos/url-shortener/internal/db"
	model "github.com/AndresCRamos/url-shortener/internal/domain/shorten_link/model"
)

type ShortenerSQLiteRepo struct {
	queries *db.Queries
}

// NewShortenerSQLiteRepo creates a new instance of ShortenerSQLiteRepo
func NewShortenerSQLiteRepo(q *db.Queries) *ShortenerSQLiteRepo {
	return &ShortenerSQLiteRepo{
		queries: q,
	}
}

func (sqr *ShortenerSQLiteRepo) Save(ctx context.Context, url *model.ShortenURLModel) (*model.ShortenURLModel, error) {
	res, err := sqr.queries.CreateShortenLink(ctx, db.CreateShortenLinkParams{
		ShortCode:   url.Short,
		OriginalUrl: url.Original,
	})
	if err != nil {
		return nil, err
	}
	return &model.ShortenURLModel{
		ID:        int(res.ID.(int64)),
		Short:     res.ShortCode,
		Original:  res.OriginalUrl,
		CreatedAt: res.CreatedAt.Time,
		Views:     int(res.Views.Int64),
	}, nil
}

func (sqr *ShortenerSQLiteRepo) FindByShort(ctx context.Context, short string) (*model.ShortenURLModel, error) {
	res, err := sqr.queries.GetShortenLink(ctx, short)
	if err != nil {
		return nil, err
	}
	return &model.ShortenURLModel{
		ID:        int(res.ID.(int64)),
		Short:     short,
		Original:  res.OriginalUrl,
		CreatedAt: res.CreatedAt.Time,
		Views:     int(res.Views.Int64),
	}, nil
}

func (sqr *ShortenerSQLiteRepo) IncrementVisitCount(ctx context.Context, short string) error {
	return sqr.queries.IncrementViews(ctx, short)
}
