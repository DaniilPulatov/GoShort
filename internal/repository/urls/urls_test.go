package urls

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"url-shortener/internal/domain/entites"
	"url-shortener/pkg/postgresDB"
)

/*
url := &entites.ShortenUrl{
			ExpiresAt:  time.Now().Add(time.Hour),
			CreatedAt:  time.Now(),
			Identifier: "identifier",
			Usages:     0,
			ID:         uuid.New().ID(),
		}
err := pool.Create(context.Background(), url)
		assert.Nil(t, err)
*/

func TestUrlRepo_Create(t *testing.T) {
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockRow := new(postgresDB.MockRow)
		defer mockPool.AssertExpectations(t)
		defer mockRow.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("QueryRow", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRow)
		mockRow.On("Scan", mock.Anything).Return(nil)
		err := pool.Create(ctx, &entites.ShortenUrl{})
		assert.NoError(t, err)
	})
	t.Run("Error in scan", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockRow := new(postgresDB.MockRow)
		defer mockPool.AssertExpectations(t)
		defer mockRow.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("QueryRow", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRow)
		mockRow.On("Scan", mock.Anything).Return(errors.New("error"))
		err := pool.Create(ctx, &entites.ShortenUrl{})
		assert.Error(t, err)
	})
}

func TestUrlRepo_UpdateUsage(t *testing.T) {
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockPool.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("Exec", ctx, mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
		err := pool.UpdateUsage(ctx, &entites.ShortenUrl{})
		assert.NoError(t, err)
	})
	t.Run("Error in exec", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockPool.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("Exec", ctx, mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("error"))
		err := pool.UpdateUsage(ctx, &entites.ShortenUrl{})
		assert.Error(t, err)
	})
}

func TestUrlRepo_Delete(t *testing.T) {
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockPool.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("Exec", ctx, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

		err := pool.Delete(ctx, mock.Anything)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockPool.AssertExpectations(t)

		pool := NewUrlRepo(mockPool)
		mockPool.On("Exec", ctx, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("error"))

		err := pool.Delete(ctx, mock.Anything)
		assert.Error(t, err)
	})
}

func TestUrlRepo_GetByIdentifier(t *testing.T) {
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockRow := new(postgresDB.MockRow)
		defer mockPool.AssertExpectations(t)
		defer mockRow.AssertExpectations(t)
		pool := NewUrlRepo(mockPool)
		mockPool.On("QueryRow", ctx, mock.Anything, mock.Anything).Return(mockRow)
		mockRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, err := pool.GetByIdentifier(ctx, mock.Anything)
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockRow := new(postgresDB.MockRow)
		defer mockPool.AssertExpectations(t)
		defer mockRow.AssertExpectations(t)
		pool := NewUrlRepo(mockPool)
		mockPool.On("QueryRow", ctx, mock.Anything, mock.Anything).Return(mockRow)
		mockRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

		_, err := pool.GetByIdentifier(ctx, mock.Anything)
		assert.Error(t, err)
	})
}
func TestUrlRepo_GetByUrl(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockPool := new(postgresDB.MockPool)
		mockRow := new(postgresDB.MockRow)
		defer mockPool.AssertExpectations(t)
		defer mockRow.AssertExpectations(t)
		pool := NewUrlRepo(mockPool)
		mockPool.On("QueryRow", context.Background(), mock.Anything, mock.Anything).Return(mockRow)
		mockRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, err := pool.GetByUrl(context.Background(), mock.Anything)
		assert.NoError(t, err)
	})
}
