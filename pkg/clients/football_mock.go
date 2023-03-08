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