package configs

import "os"

const (
	FootballApikeyKey = "FOOTBALL_APIKEY"
)

var (
	apiKey string
)

func GetFootballAPIKey() string {
	return apiKey
}

func init() {
	apiKey = os.Getenv(FootballApikeyKey)
}