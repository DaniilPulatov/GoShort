package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"url-shortener/internal/domain/entites"
)

type UrlsServiceMock struct {
	mock.Mock
}

func (m *UrlsServiceMock) Shorten(ctx context.Context, url *entites.InputUrl) (string, error) {
	args := m.Called(ctx, url)
	m.On("GetByIdentifier", ctx, url.Identifier).Return(url, nil)
	return args.String(0), args.Error(1)
}

func (m *UrlsServiceMock) Redirect(ctx context.Context, token string) (*entites.ShortenUrl, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*entites.ShortenUrl), args.Error(1)
}
