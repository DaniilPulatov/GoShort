package urls

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
	"url-shortener/internal/domain/entites"
	urlRepo "url-shortener/internal/repository/urls"
)

func TestService_Shorten(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo := new(urlRepo.UrlRepositoryMock)
		s := &service{urlRepo: repo}
		ctx := context.Background()

		input := &entites.InputUrl{
			RealUrl: "https://example.com",
		}

		repo.On("GetByUrl", context.Background(), input.RealUrl).
			Return(nil, pgx.ErrNoRows)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(su *entites.ShortenUrl) bool {
			return su.RealUrl == input.RealUrl && su.Identifier != ""
		})).Return(nil)

		os.Setenv("BASE_URL", "https://short.ly")

		short, err := s.Shorten(ctx, input)
		require.NoError(t, err)
		require.Contains(t, short, "https://short.ly/")
	})

	t.Run("exist but not expired", func(t *testing.T) {
		repo := new(urlRepo.UrlRepositoryMock)
		s := &service{urlRepo: repo}
		ctx := context.Background()

		input := &entites.InputUrl{
			RealUrl: "https://example.com",
		}
		repo.On("GetByUrl", context.Background(), input.RealUrl).Return(&entites.ShortenUrl{
			RealUrl:   input.RealUrl,
			ExpiresAt: time.Now().Local().Add(time.Hour),
		}, nil)

		res, err := s.Shorten(ctx, input)
		require.Error(t, err)
		require.Equal(t, res, "")
	})

	t.Run("exist and expired", func(t *testing.T) {
		repo := new(urlRepo.UrlRepositoryMock)
		s := &service{urlRepo: repo}
		ctx := context.Background()

		input := &entites.InputUrl{
			RealUrl: "https://example.com",
		}
		repo.On("GetByUrl", context.Background(), input.RealUrl).Return(&entites.ShortenUrl{
			RealUrl:   input.RealUrl,
			ExpiresAt: time.Now().Local(),
		}, nil)

		repo.On("Delete", ctx, input.RealUrl).Return(nil)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(su *entites.ShortenUrl) bool {
			return su.RealUrl == input.RealUrl && su.Identifier != ""
		})).Return(nil)

		os.Setenv("BASE_URL", "https://mu")

		res, err := s.Shorten(ctx, input)
		require.NoError(t, err)
		require.Contains(t, res, "https://mu/")
	})
}
