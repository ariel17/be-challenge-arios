package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

type Status struct {
	OK     bool   `json:"ok"`
	Detail string `json:"detail"`
}

type StatusService interface {
	GetStatus() Status
}

func NewStatusService(repository repositories.Repository) StatusService {
	return &realStatusService{
		Repository: repository,
	}
}

type realStatusService struct {
	Repository repositories.Repository
}

// GetStatus checks the application's health and returns and object describing
// it.
func (r *realStatusService) GetStatus() Status {
	if err := r.Repository.GetStatus(); err != nil {
		return Status{
			OK:     false,
			Detail: err.Error(),
		}
	}
	return Status{
		OK: true,
	}
}