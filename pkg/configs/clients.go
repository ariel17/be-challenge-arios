package configs

import (
	"os"
	"strconv"
)

const (
	FootballApikeyKey               = "FOOTBALL_APIKEY"
	FootballMaxRequestsPerMinuteKey = "FOOTBALL_MAX_REQUESTS_PER_MINUTE"
	DefaultMaxRequestsPerMinute     = 10
)

var (
	apiKey               string
	maxRequestsPerMinute int
)

func GetFootballAPIKey() string {
	return apiKey
}

func GetFootballMaxRequestsPerMinute() int {
	return maxRequestsPerMinute
}

func init() {
	apiKey = os.Getenv(FootballApikeyKey)

	temp := os.Getenv(FootballMaxRequestsPerMinuteKey)
	max, err := strconv.Atoi(temp)
	if err != nil {
		max = DefaultMaxRequestsPerMinute
	} else {
		maxRequestsPerMinute = max
	}
}