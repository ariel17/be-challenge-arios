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
	"golang.org/x/time/rate"
)

var (
	httpClient = &http.Client{
		Timeout: time.Second,
	}
	apiKey = "abc123"
)

func TestGet(t *testing.T) {
	t.Run("rate limit applied", func(t *testing.T) {
		apiContent := loadGoldenFile(t.Name())
		server := newTestServer("/", 200, apiContent)
		defer server.Close()

		c := &realAPIClient{
			baseURL:     server.URL,
			client:      httpClient,
			apiKey:      apiKey,
			rateLimiter: rate.NewLimiter(rate.Every(time.Second), 1),
		}

		start := time.Now()
		for i := 0; i < 2; i++ {
			response, err := c.get("/")
			assert.NotNil(t, response)
			assert.Nil(t, err)
		}
		elapsed := time.Since(start)

		assert.True(t, elapsed.Seconds() >= float64(1))
	})

	t.Run("not found", func(t *testing.T) {
		apiContent := loadGoldenFile(t.Name())
		server := newTestServer("/", 404, apiContent)
		defer server.Close()

		c := &realAPIClient{
			baseURL: server.URL,
			client:  httpClient,
			apiKey:  apiKey,
		}
		response, err := c.get("/")
		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve content:")
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

func TestGetTeamsByLeagueCode(t *testing.T) {
	testCases := []struct {
		name       string
		code       string
		statusCode int
		isSuccess  bool
	}{
		{"ok", "WC", 200, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiContent := loadGoldenFile(t.Name())
			server := newTestServer("/competitions/"+tc.code+"/teams", tc.statusCode, apiContent)
			defer server.Close()

			c := &realAPIClient{
				baseURL: server.URL,
				client:  httpClient,
				apiKey:  apiKey,
			}
			response, err := c.GetTeamsByLeagueCode(tc.code)
			assert.Equal(t, err == nil, tc.isSuccess)
			assert.Equal(t, response != nil, tc.isSuccess)

			if tc.isSuccess {
				assert.Equal(t, 32, len(response))
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
				assert.Equal(t, int64(2061), response.ID)
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

func TestGetPersonByID(t *testing.T) {
	client = &http.Client{
		Timeout: time.Second,
	}

	testCases := []struct {
		name       string
		id         int64
		statusCode int
		isSuccess  bool
	}{
		{"ok", 44, 200, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiContent := loadGoldenFile(t.Name())
			server := newTestServer(fmt.Sprintf("/persons/%d", tc.id), tc.statusCode, apiContent)
			defer server.Close()

			c := &realAPIClient{
				baseURL: server.URL,
				client:  httpClient,
				apiKey:  apiKey,
			}
			response, err := c.GetPersonByID(tc.id)
			assert.Equal(t, err == nil, tc.isSuccess)
			assert.Equal(t, response != nil, tc.isSuccess)

			if tc.isSuccess {
				assert.Equal(t, "Cristiano Ronaldo", response.Name)
				assert.Equal(t, int64(44), response.ID)
				assert.Equal(t, "Centre-Forward", response.Position)
				assert.Equal(t, "1985-02-05", response.DateOfBirth)
				assert.Equal(t, "Portugal", response.Nationality)
			} else {
				assert.True(t, strings.Contains(err.Error(), "failed to retrieve content:"))
			}
		})
	}
}

// loadGoldenFiles uses the test name to load a JSON value as expected result.
func loadGoldenFile(name string) []byte {
	content, err := os.ReadFile("./golden/" + name + ".json")
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