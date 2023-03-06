package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected int
	}{
		{"port is not present", "", defaultPort},
		{"port is invalid", "text", defaultPort},
		{"port is valid", "9090", 9090},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Getenv(portKey)
			defer func() {
				os.Setenv(portKey, old)
			}()

			os.Setenv(portKey, tc.value)
			loadServerConfig()

			result := GetPort()
			assert.Equal(t, tc.expected, result)
		})
	}
}