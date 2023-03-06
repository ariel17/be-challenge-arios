package configs

import "os"

const (
	FOOTBALL_APIKEY_KEY = "FOOTBALL_APIKEY"
)

var (
	apiKey string
)

func GetFootballAPIKey() string {
	return apiKey
}

func init() {
	apiKey = os.Getenv(FOOTBALL_APIKEY_KEY)
}