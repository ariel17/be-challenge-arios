package configs

import (
	"os"
)

const (
	dsnKey             = "DATABASE_DSN"
	statusQueryKey     = "DATABASE_STATUS_QUERY"
	defaultStatusQuery = "SELECT 1"
)

var (
	dsn         string
	statusQuery string
)

// GetDSN returns the DSN connection string for the MySQL database.
func GetDSN() string {
	return dsn
}

// GetStatusQuery returns the SQL query to execute when verifying application
// status.
func GetStatusQuery() string {
	return statusQuery
}

func init() {
	dsn = os.Getenv(dsnKey)

	statusQuery = os.Getenv(statusQueryKey)
	if statusQuery == "" {
		statusQuery = defaultStatusQuery
	}
}