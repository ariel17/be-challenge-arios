package repositories

import (
	"github.com/stretchr/testify/mock"

	"github.com/ariel17/be-challenge-arios/pkg/models"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) GetStatus() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) CreateSchema() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) AddPerson(person models.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockRepository) AddTeam(team models.Team) error {
	args := m.Called(team)
	return args.Error(0)
}

func (m *MockRepository) AddCompetition(competition models.Competition) error {
	args := m.Called(competition)
	return args.Error(0)
}

func (m *MockRepository) AddTeamToCompetition(team models.Team, competition models.Competition) error {
	args := m.Called(team, competition)
	return args.Error(0)
}

func (m *MockRepository) AddPersonToTeam(person models.Person, team models.Team) error {
	args := m.Called(person, team)
	return args.Error(0)
}

func (m *MockRepository) GetTeamByTLA(tla string) (*models.Team, error) {
	args := m.Called(tla)
	team, ok := args.Get(0).(*models.Team)
	if !ok {
		return nil, args.Error(1)
	}
	return team, args.Error(1)
}

func (m *MockRepository) GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, error) {
	args := m.Called(code, teamNameToFilter)
	players, ok := args.Get(0).([]models.Person)
	if !ok {
		return nil, args.Error(1)
	}
	return players, nil
}

func (m *MockRepository) GetPersonsByTeamTLA(tla string) ([]models.Person, error) {
	args := m.Called(tla)
	persons, ok := args.Get(0).([]models.Person)
	if !ok {
		return nil, args.Error(1)
	}
	return persons, args.Error(1)
}

func (m *MockRepository) CompetitionExists(code string) (bool, error) {
	args := m.Called(code)
	return args.Bool(0), args.Error(1)
}