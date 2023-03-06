package services

import (
	"github.com/ariel17/golang-base/pkg/configs"
)

const (
	okStatus    = "ok"
	errorStatus = "error"
)

type Status struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

// GetStatus checks the application's health and returns and object describing
// it.
func GetStatus() (Status, error) {
	if tx := configs.GetDB().Raw(configs.GetStatusQuery()); tx.Error != nil {
		return Status{
			Status: errorStatus,
			Detail: tx.Error.Error(),
		}, tx.Error
	}
	return Status{
		Status: okStatus,
	}, nil
}