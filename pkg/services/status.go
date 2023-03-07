package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

const (
	okStatus    = "ok"
	errorStatus = "error"
)

var (
	repository repositories.Repository
)

type Status struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func (s Status) IsError() bool {
	return s.Status == errorStatus
}

// GetStatus checks the application's health and returns and object describing
// it.
func GetStatus() Status {
	if err := repository.GetStatus(); err != nil {
		return Status{
			Status: errorStatus,
			Detail: err.Error(),
		}
	}
	return Status{
		Status: okStatus,
	}
}

func init() {
	repository = repositories.NewMySQLRepository()
}