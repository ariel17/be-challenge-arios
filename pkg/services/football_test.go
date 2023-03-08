package services

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ariel17/be-challenge-arios/pkg/models"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

func TestRealFootballService_GetPlayersByCompetitionCode(t *testing.T) {
	code := "ABC"
	testCases := []struct {
		name                       string
		competitionExistsIsSuccess bool
		competitionExists          bool
		hasPlayersIsSuccess        bool
		hasPlayers                 bool
		isSuccess                  bool
	}{
		{"ok when competition exists and has players", true, true, true, true, true},
		{"ok when competition exists and has no players", true, true, true, false, true},
		{"ok when competition does not exist", true, false, false, false, true},
		{"failed when competition existence check fails", false, false, false, false, false},
		{"failed when fetching players from team fails", true, true, false, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repositories.MockRepository{}
			service := realFootballService{
				Repository: &mockRepository,
			}

			on := mockRepository.On("CompetitionExists", code)
			if tc.competitionExistsIsSuccess {
				on.Return(tc.competitionExists, nil)

				on = mockRepository.On("GetPlayersByCompetitionCode", code, "")
				if tc.hasPlayersIsSuccess {
					if tc.hasPlayers {
						position := "Defender"
						players := []models.Person{
							{1, "Ariel", &position, "1980-01-01", "Argentinian"},
						}
						on.Return(players, nil)
					} else {
						on.Return([]models.Person{}, nil)
					}
				} else {
					on.Return(nil, errors.New("mocked error"))
				}

			} else {
				on.Return(false, errors.New("mocked error"))
			}

			players, competitionExists, err := service.GetPlayersByCompetitionCode(code, "")
			assert.Equal(t, err == nil, tc.isSuccess)
			assert.Equal(t, len(players) > 0, tc.hasPlayers)
			assert.Equal(t, competitionExists, tc.competitionExists && tc.hasPlayersIsSuccess)
		})
	}
}

func TestRealFootballService_GetTeamByTLA(t *testing.T) {
	tla := "ABC"
	testCases := []struct {
		name                string
		getTeamIsSuccess    bool
		getPersonsIsSuccess bool
		withPlayers         bool
		withCoach           bool
		isSuccess           bool
	}{
		{"ok team with players and coach", true, true, true, true, true},
		{"ok team with players and no coach", true, true, true, false, true},
		{"ok team without players and no coach", true, true, false, false, true},
		{"failed when get team fails", false, false, false, false, false},
		{"failed when get players fails", true, false, true, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repositories.MockRepository{}
			service := realFootballService{
				Repository: &mockRepository,
			}

			on := mockRepository.On("GetTeamByTLA", tla)
			if tc.getTeamIsSuccess {
				team := models.Team{
					TLA:       tla,
					Name:      "Test Team",
					ShortName: "TestTeam",
					AreaName:  "Argentina",
					Address:   "Evergreen 123",
				}
				on.Return(&team, nil)

				on = mockRepository.On("GetPersonsByTeamTLA", tla)
				if tc.getPersonsIsSuccess {
					persons := []models.Person{}
					if tc.withPlayers {
						position := "Defender"
						persons = append(persons, models.Person{Position: &position})
					}
					if tc.withCoach {
						persons = append(persons, models.Person{})
					}
					on.Return(persons, nil)

				} else {
					on.Return(nil, errors.New("mocked error"))
				}

			} else {
				on.Return(nil, errors.New("mocked error"))
			}

			result, err := service.GetTeamByTLA(tla, tc.withPlayers)
			assert.Equal(t, err == nil, tc.isSuccess)
			if tc.isSuccess {
				assert.Equal(t, len(result.Players) > 0, tc.withPlayers)
				assert.Equal(t, result.Coach != nil, tc.withCoach)
			}
		})
	}
}