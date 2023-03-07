package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ariel17/be-challenge-arios/pkg/clients"
	"github.com/ariel17/be-challenge-arios/pkg/models"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

func TestRealImporterService_ImportDataByCompetitionCode(t *testing.T) {
	code := "ABC"
	teamWithPlayers := clients.Team{
		Name:      "Test Team With Players",
		ShortName: "TestTeamPlayers",
		TLA:       "TTP",
		Address:   "Evergreen 123",
		Squad: []clients.Person{
			{
				ID:          1,
				Name:        "Player",
				Position:    "Forward",
				DateOfBirth: "1980-01-01",
				Nationality: "Argentinian",
			},
		},
	}
	teamWithoutPlayers := clients.Team{
		Name:      "Test Team No Players",
		ShortName: "TestTeamNoPlayers",
		TLA:       "TTNP",
		Address:   "Evergreen 123",
		Coach: clients.Person{
			ID:          2,
			Name:        "Coach",
			DateOfBirth: "1980-01-01",
			Nationality: "Argentinian",
		},
	}

	testCases := []struct {
		name                          string
		getCompetitionIsSuccess       bool
		addCompetitionIsSuccess       bool
		getTeamsByLeagueCodeIsSuccess bool
		addTeamIsSuccess              bool
		addTeamToCompetitionIsSuccess bool
		team                          *clients.Team
		addPersonIsSuccess            bool
		addPersonToTeamIsSuccess      bool
		isSuccess                     bool
	}{
		{"ok team with players", true, true, true, true, true, &teamWithPlayers, true, true, true},
		{"ok team without players", true, true, true, true, true, &teamWithoutPlayers, true, true, true},
		{"failed to get competitions by code", false, false, false, false, false, nil, false, false, false},
		{"failed to save competition", true, false, false, false, false, nil, false, false, false},
		{"failed to get teams by competition code", true, true, false, false, false, nil, false, false, false},
		{"failed to save team", true, true, true, false, false, &teamWithPlayers, false, false, false},
		{"failed to add team to competition", true, true, true, true, false, &teamWithPlayers, false, false, false},
		{"failed to save person", true, true, true, true, true, &teamWithPlayers, false, false, false},
		{"failed to add person to team", true, true, true, true, true, &teamWithPlayers, true, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAPIClient := clients.MockAPIClient{}
			mockRepository := repositories.MockRepository{}
			service := realImporterService{
				Client:     &mockAPIClient,
				Repository: &mockRepository,
			}

			// scenario set up
			on := mockAPIClient.On("GetLeagueByCode", code)
			if tc.getCompetitionIsSuccess {
				competition := clients.League{
					Code: code,
					Name: "Test Competition",
					Area: clients.Area{
						Name: "Argentina",
					},
				}
				on.Return(&competition, nil)
			} else {
				on.Return(nil, errors.New("mocked error"))
			}

			on = mockRepository.On("AddCompetition", mock.AnythingOfType("models.Competition"))
			if tc.addCompetitionIsSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			on = mockAPIClient.On("GetTeamsByLeagueCode", code)
			if tc.getTeamsByLeagueCodeIsSuccess {
				teams := []clients.Team{*tc.team}
				on.Return(teams, nil)
			} else {
				on.Return(nil, errors.New("mocked error"))
			}

			on = mockRepository.On("AddTeam", mock.AnythingOfType("models.Team"))
			if tc.addTeamIsSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			on = mockRepository.On("AddTeamToCompetition", mock.AnythingOfType("models.Team"), mock.AnythingOfType("models.Competition"))
			if tc.addTeamToCompetitionIsSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			on = mockRepository.On("AddPerson", mock.AnythingOfType("models.Person"))
			if tc.addPersonIsSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			on = mockRepository.On("AddPersonToTeam", mock.AnythingOfType("models.Person"), mock.AnythingOfType("models.Team"))
			if tc.addPersonToTeamIsSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			// do & assert
			err := service.ImportDataByCompetitionCode(code)
			assert.Equal(t, err == nil, tc.isSuccess)

			if tc.team != nil && tc.addPersonIsSuccess {
				matcher := mock.MatchedBy(func(person models.Person) bool {
					if len(tc.team.Squad) > 0 {
						return person.Name == "Player"
					}
					return person.Name == "Coach"
				})
				mockRepository.AssertCalled(t, "AddPerson", matcher)
			}
		})
	}
}