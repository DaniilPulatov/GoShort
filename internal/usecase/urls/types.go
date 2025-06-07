package urls

import (
	"context"
	"url-shortener/internal/domain/entites"
	"url-shortener/internal/repository/urls"
)

type UrlService interface {
	Shorten(ctx context.Context, url *entites.InputUrl) (string, error)
	Redirect(ctx context.Context, token string) (*entites.ShortenUrl, error)
}

type service struct {
	urlRepo urls.UrlRepository
}

func NewUrlService(urlRepo urls.UrlRepository) UrlService {
	return &service{urlRepo: urlRepo}
}
