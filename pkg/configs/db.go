package configs

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnKey = "DATABASE_DSN"
	statusQueryKey = "DATABASE_STATUS_QUERY"
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

// GetDB returns a GORM object connected to database, if everything goes well.
func GetDB() *gorm.DB {
	log.Printf("Connecting to database with DSN: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func loadDBConfig() {
	dsn = os.Getenv(dsnKey)
	statusQuery = os.Getenv(statusQueryKey)
}

func init() {
	loadDBConfig()
}