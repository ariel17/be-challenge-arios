package clients

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFootballAPIClient(t *testing.T) {
	t.Run("fails without api key", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		NewFootballAPIClient("")
	})

	t.Run("ok", func(t *testing.T) {
		NewFootballAPIClient("abc123")
	})
}

func TestGetLeagueByCode(t *testing.T) {
	testCases := []struct {
		name       string
		code       string
		statusCode int
		isSuccess  bool
	}{
		{"ok", "PL", 200, true},
		{"not found", "XXX", 404, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiContent := loadGoldenFile(t.Name())
			server := newTestServer("/competitions/"+tc.code, tc.statusCode, apiContent)
			defer server.Close()

			c := &realAPIClient{
				baseURL: server.URL,
				client: &http.Client{
					Timeout: time.Second,
				},
				apiKey: "abc123",
			}
			response, err := c.GetLeagueByCode(tc.code)
			assert.Equal(t, err == nil, tc.isSuccess)
			assert.Equal(t, response != nil, tc.isSuccess)

			if tc.isSuccess {
				assert.Equal(t, "Premier League", response.Name)
				assert.Equal(t, "PL", response.Code)
				assert.Equal(t, "England", response.Area.Name)
			} else {
				assert.True(t, strings.Contains(err.Error(), "failed to retrieve content:"))
			}
		})
	}
}

// loadGoldenFiles uses the test name to load a JSON value as expected result.
func loadGoldenFile(testName string) []byte {
	testName = strings.ReplaceAll(testName, " ", "_")
	content, err := os.ReadFile("./golden/" + testName + ".json")
	if err != nil {
		panic(err)
	}
	return content
}

func newTestServer(url string, statusCode int, responseBody []byte) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	})
	return httptest.NewServer(mux)
}