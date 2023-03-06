package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	httpClient = &http.Client{
		Timeout: time.Second,
	}
	apiKey = "abc123"
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
		NewFootballAPIClient(apiKey)
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
				client:  httpClient,
				apiKey:  apiKey,
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

func TestGetTeamByID(t *testing.T) {
	client = &http.Client{
		Timeout: time.Second,
	}

	testCases := []struct {
		name       string
		id         int64
		statusCode int
		isSuccess  bool
	}{
		{"ok", 2061, 200, true},
		{"not found", 999, 404, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiContent := loadGoldenFile(t.Name())
			server := newTestServer(fmt.Sprintf("/teams/%d", tc.id), tc.statusCode, apiContent)
			defer server.Close()

			c := &realAPIClient{
				baseURL: server.URL,
				client:  httpClient,
				apiKey:  apiKey,
			}
			response, err := c.GetTeamByID(tc.id)
			assert.Equal(t, err == nil, tc.isSuccess)
			assert.Equal(t, response != nil, tc.isSuccess)

			if tc.isSuccess {
				assert.Equal(t, "CA Boca Juniors", response.Name)
				assert.Equal(t, 2061, response.ID)
				assert.Equal(t, "Boca Juniors", response.ShortName)
				assert.Equal(t, "BOC", response.TLA)
				assert.Equal(t, "Brandsen 805, La Boca Buenos Aires, Buenos Aires 1161", response.Address)
				assert.Equal(t, "Argentina", response.Area.Name)
			} else {
				assert.True(t, strings.Contains(err.Error(), "failed to retrieve content:"))
			}
		})
	}
}

// loadGoldenFiles uses the test name to load a JSON value as expected result.
func loadGoldenFile(testName string) []byte {
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