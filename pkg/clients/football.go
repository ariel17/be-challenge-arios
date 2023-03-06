package clients

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

const BASE_API_URL = "http://api.football-data.org/v4/"

type Area struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type League struct {
	Area Area   `json:"area"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`
}

// FootballAPIClient is the behavior contract that every implementation must
// comply. It offers access to football-data.org data with handy methods. It is
// NOT a full client implementation but access to required resources.
type FootballAPIClient interface {

	// GetLeagueByCode retrieves league data by its code.
	GetLeagueByCode(code string) (*League, error)
}

// NewFootballAPIClient creates a new instance of real API client.
func NewFootballAPIClient(apiKey string) FootballAPIClient {
	if apiKey == "" {
		panic("cannot work without a key")
	}
	return &realAPIClient{
		baseURL: BASE_API_URL,
		client: &http.Client{
			Timeout: time.Second,
		},
		apiKey: apiKey,
	}
}

type realAPIClient struct {
	baseURL string
	client  *http.Client
	apiKey  string
}

func (r *realAPIClient) GetLeagueByCode(code string) (*League, error) {
	url := r.baseURL + "/competitions/" + code

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("X-Auth-Token", r.apiKey)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	var body []byte
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("failed to retrieve content: " + string(body))
	}

	league := League{}
	err = json.Unmarshal(body, &league)
	if err != nil {
		return nil, err
	}
	return &league, nil
}