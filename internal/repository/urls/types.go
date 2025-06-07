package urls

import (
	"context"
	"url-shortener/internal/domain/entites"
	"url-shortener/pkg/postgresDB"
)

type UrlRepository interface {
	Create(ctx context.Context, url *entites.ShortenUrl) error
	UpdateUsage(ctx context.Context, url *entites.ShortenUrl) error
	GetByIdentifier(ctx context.Context, identifier string) (*entites.ShortenUrl, error)
	GetByUrl(ctx context.Context, url string) (*entites.ShortenUrl, error)
	Delete(ctx context.Context, url string) error
}

type repo struct {
	pool postgresDB.Pool
}

func NewUrlRepo(pool postgresDB.Pool) UrlRepository {
	return &repo{pool: pool}
}
