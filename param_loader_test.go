package dagster_pipes

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvVarLoader(t *testing.T) {
	t.Parallel()
	encoded := "eJwVwdEJgCAQANBV4ha4SDNsjhYQNf0wFTtEjXaP3nsgK/KwT4BVFTxTMLbc2DakIFwfUS516MJWkrw5En3+uYwH0pWDZzpyW1GnSLYRvB9CZRtp"
	expected := map[string]json.RawMessage{
		"path": json.RawMessage([]byte(`"/var/folders/x7/tl6gyzn92vzcr35t94xgt6y00000gp/T/tmplh3cn4ev/context"`)),
	}

	loader := NewEnvVarLoader()

	t.Run("LoadContextParams", func(t *testing.T) {
		t.Parallel()
		os.Setenv(DAGSTER_PIPES_CONTEXT_ENV_VAR, encoded)

		result, err := loader.LoadContextParams()
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("LoadMessageParams", func(t *testing.T) {
		t.Parallel()
		os.Setenv(DAGSTER_PIPES_MESSAGES_ENV_VAR, encoded)

		result, err := loader.LoadMessageParams()
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})
}
