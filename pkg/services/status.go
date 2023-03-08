package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

type StatusService interface {
	GetStatus() error
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
func (r *realStatusService) GetStatus() error {
	return r.Repository.GetStatus()
}