package configs

import (
	"os"
	"strconv"
)

const (
	portKey     = "PORT"
	defaultPort = 8080
)

var (
	port int
)

// GetPort returns the port value where the server will be listening for new
// connections.
func GetPort() int {
	return port
}

func loadServerConfig() {
	if p, err := strconv.Atoi(os.Getenv(portKey)); err != nil {
		port = defaultPort
	} else {
		port = p
	}
}

func init() {
	loadServerConfig()
}