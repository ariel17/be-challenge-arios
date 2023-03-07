package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/ariel17/be-challenge-arios/pkg/configs"
)

const BASE_API_URL = "http://api.football-data.org/v4/"

var (
	client      *http.Client
	rateLimiter *rate.Limiter
)

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

type Person struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	DateOfBirth string `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
}

type Team struct {
	Area      Area     `json:"area"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	ShortName string   `json:"shortName"`
	TLA       string   `json:"tla"`
	Address   string   `json:"address"`
	Coach     Person   `json:"coach"`
	Squad     []Person `json:"squad"`
}

type teams struct {
	Teams []Team `json:"teams"`
}

// FootballAPIClient is the behavior contract that every implementation must
// comply. It offers access to football-data.org data with handy methods. It is
// NOT a full client implementation but access to required resources.
type FootballAPIClient interface {
	GetLeagueByCode(code string) (*League, error)
	GetTeamsByLeagueCode(code string) ([]Team, error)
	GetTeamByID(id int64) (*Team, error)
	GetPersonByID(id int64) (*Person, error)
}

// NewFootballAPIClient creates a new instance of real API client.
func NewFootballAPIClient() FootballAPIClient {
	return &realAPIClient{
		baseURL:     BASE_API_URL,
		client:      client,
		apiKey:      configs.GetFootballAPIKey(),
		rateLimiter: rateLimiter,
	}
}

type realAPIClient struct {
	baseURL     string
	client      *http.Client
	apiKey      string
	rateLimiter *rate.Limiter
}

func (r *realAPIClient) get(path string) ([]byte, error) {
	if r.rateLimiter != nil {
		ctx := context.Background()
		if err := r.rateLimiter.Wait(ctx); err != nil {
			return nil, err
		}
	}

	url := r.baseURL + path
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

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve content: " + string(body))
	}

	return body, nil
}

func (r *realAPIClient) GetLeagueByCode(code string) (*League, error) {
	url := fmt.Sprintf("/competitions/%s", code)
	body, err := r.get(url)
	if err != nil {
		return nil, err
	}
	league := League{}
	err = json.Unmarshal(body, &league)
	if err != nil {
		return nil, err
	}
	return &league, nil
}

func (r *realAPIClient) GetTeamsByLeagueCode(code string) ([]Team, error) {
	url := fmt.Sprintf("/competitions/%s/teams", code)
	body, err := r.get(url)
	if err != nil {
		return nil, err
	}

	t := teams{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return t.Teams, nil
}

func (r *realAPIClient) GetTeamByID(id int64) (*Team, error) {
	url := fmt.Sprintf("/teams/%d", id)
	body, err := r.get(url)
	if err != nil {
		return nil, err
	}
	team := Team{}
	err = json.Unmarshal(body, &team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *realAPIClient) GetPersonByID(id int64) (*Person, error) {
	url := fmt.Sprintf("/persons/%d", id)
	body, err := r.get(url)
	if err != nil {
		return nil, err
	}
	person := Person{}
	err = json.Unmarshal(body, &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func init() {
	rateLimiter = rate.NewLimiter(rate.Every(time.Minute), configs.GetFootballMaxRequestsPerMinute())
	client = &http.Client{
		Timeout: time.Second,
	}
}