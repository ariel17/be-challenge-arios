package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProduction(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected bool
	}{
		{"environment is production", productionEnv, true},
		{"environment not is production", "test", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Getenv(environmentKey)
			defer func() {
				os.Setenv(environmentKey, old)
			}()

			os.Setenv(environmentKey, tc.value)
			loadEnvsConfig()

			result := IsProduction()
			assert.Equal(t, tc.expected, result)
		})
	}
}