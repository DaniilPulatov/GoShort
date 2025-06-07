package di

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestNewModule(t *testing.T) {
	t.Run("module is running", func(t *testing.T) {
		err := fx.ValidateApp(NewModule())
		require.NoError(t, err)
	})
}
