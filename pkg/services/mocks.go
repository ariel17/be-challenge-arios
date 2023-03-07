package services

import (
	"github.com/stretchr/testify/mock"
)

type MockImporterService struct {
	mock.Mock
}

func (m *MockImporterService) ImportDataByCompetitionCode(code string) error {
	args := m.Called(code)
	return args.Error(0)
}

type MockStatusService struct {
	mock.Mock
}

// GetStatus checks the application's health and returns and object describing
// it.
func (m *MockStatusService) GetStatus() Status {
	args := m.Called()
	return args.Get(0).(Status)
}