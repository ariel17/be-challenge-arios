package clients

import (
	"github.com/stretchr/testify/mock"
)

type MockAPIClient struct {
	mock.Mock
}

func (m *MockAPIClient) GetLeagueByCode(code string) (*League, error) {
	args := m.Called(code)
	league, ok := args.Get(0).(*League)
	if !ok {
		return nil, args.Error(1)
	}
	return league, args.Error(1)
}

func (m *MockAPIClient) GetTeamsByLeagueCode(code string) ([]Team, error) {
	args := m.Called(code)
	teams, ok := args.Get(0).([]Team)
	if !ok {
		return nil, args.Error(1)
	}
	return teams, args.Error(1)
}

func (m *MockAPIClient) GetTeamByID(id int64) (*Team, error) {
	args := m.Called(id)
	team, ok := args.Get(0).(*Team)
	if !ok {
		return nil, args.Error(1)
	}
	return team, args.Error(1)
}

func (m *MockAPIClient) GetPersonByID(id int64) (*Person, error) {
	args := m.Called(id)
	person, ok := args.Get(0).(*Person)
	if !ok {
		return nil, args.Error(1)
	}
	return person, args.Error(1)
}