package services

import (
	"github.com/stretchr/testify/mock"

	"github.com/ariel17/be-challenge-arios/pkg/models"
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

func (m *MockStatusService) GetStatus() Status {
	args := m.Called()
	return args.Get(0).(Status)
}

type MockFootballService struct {
	mock.Mock
}

func (m *MockStatusService) GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, bool, error) {
	args := m.Called(code, teamNameToFilter)
	return args.Get(0).([]models.Person), args.Bool(1), args.Error(2)
}

func (m *MockStatusService) GetTeamByTLA(tla string, withPlayers bool) (*models.Team, error) {
	args := m.Called(tla, withPlayers)
	return args.Get(0).(*models.Team), args.Error(1)
}

func (m *MockStatusService) GetPersonsByTeamTLA(tla string) ([]models.Person, error) {
	args := m.Called(tla)
	return args.Get(0).([]models.Person), args.Error(1)
}