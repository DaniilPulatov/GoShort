package shortening

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Run("output is not empty", func(t *testing.T) {
		res := Shorten(uuid.New().ID())
		assert.NotEmpty(t, res, "output should not be empty")
	})

	t.Run("output is empty", func(t *testing.T) {
		res := Shorten(0)
		assert.Empty(t, res, "output should be empty")
	})

	t.Run("same output for same input", func(t *testing.T) {
		id := uuid.New().ID()
		res := Shorten(id)
		for range 100 {
			temp := Shorten(id)
			assert.Equal(t, temp, res, "for same input output should be the same")
		}
	})
}

func TestAddBaseUrl(t *testing.T) {
	t.Run("base url with https", func(t *testing.T) {
		identifier := "GhDr_DaAS"
		baseUrl := "http://baseUrl"
		res, err := AddBaseUrl(baseUrl, identifier)
		require.NoError(t, err)
		assert.Equal(t, "http://baseUrl/GhDr_DaAS", res)
	})

	t.Run("base url as a plain word result an incorrect result", func(t *testing.T) {
		identifier := "GhDr_DaAS"
		baseUrl := "plainWord"
		res, err := AddBaseUrl(baseUrl, identifier)
		require.NoError(t, err)
		assert.NotEqual(t, "plainWord/GhDr_DaAS", res, "output should be without base url")
		assert.Equal(t, "GhDr_DaAS", res)
	})

	t.Run("base url and identifier are empty", func(t *testing.T) {
		identifier := ""
		baseUrl := ""
		res, err := AddBaseUrl(baseUrl, identifier)
		require.NoError(t, err)
		assert.Equal(t, "", res)
	})
}
