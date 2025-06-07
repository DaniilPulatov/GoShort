package urls

import (
	"context"
	"github.com/stretchr/testify/mock"
	"url-shortener/internal/domain/entites"
)

type UrlRepositoryMock struct {
	mock.Mock
}

func NewUrlRepositoryMock() *UrlRepositoryMock {
	return &UrlRepositoryMock{}
}

func (m *UrlRepositoryMock) Create(ctx context.Context, url *entites.ShortenUrl) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

func (m *UrlRepositoryMock) UpdateUsage(ctx context.Context, url *entites.ShortenUrl) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

func (m *UrlRepositoryMock) GetByIdentifier(ctx context.Context, identifier string) (*entites.ShortenUrl, error) {
	args := m.Called(ctx, identifier)
	if res := args.Get(0); res != nil {
		return res.(*entites.ShortenUrl), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UrlRepositoryMock) GetByUrl(ctx context.Context, url string) (*entites.ShortenUrl, error) {
	args := m.Called(ctx, url)
	if res := args.Get(0); res != nil {
		return res.(*entites.ShortenUrl), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UrlRepositoryMock) Delete(ctx context.Context, url string) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}
