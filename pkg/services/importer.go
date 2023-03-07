package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/clients"
	"github.com/ariel17/be-challenge-arios/pkg/models"
)

var (
	client clients.FootballAPIClient
)

func ImportDataByCompetitionCode(code string) error {
	rawCompetition, err := client.GetLeagueByCode(code)
	if err != nil {
		return err
	}
	competition := models.Competition{
		Code:     rawCompetition.Code,
		Name:     rawCompetition.Name,
		AreaName: rawCompetition.Area.Name,
	}
	if err := repository.AddCompetition(competition); err != nil {
		return err
	}

	rawTeams, err := client.GetTeamsByLeagueCode(code)
	if err != nil {
		return err
	}
	for _, rawTeam := range rawTeams {
		team := models.Team{
			TLA:       rawTeam.TLA,
			Name:      rawTeam.Name,
			ShortName: rawTeam.ShortName,
			AreaName:  rawTeam.Area.Name,
			Address:   rawTeam.Address,
		}
		if err := repository.AddTeam(team); err != nil {
			return err
		}
		if err := repository.AddTeamToCompetition(team, competition); err != nil {
			return err
		}

		coach := models.Person{
			ID:          rawTeam.Coach.ID,
			Name:        rawTeam.Coach.Name,
			DateOfBirth: rawTeam.Coach.DateOfBirth,
			Nationality: rawTeam.Coach.Nationality,
		}
		if err := repository.AddPerson(coach); err != nil {
			return err
		}
		if err := repository.AddPersonToTeam(coach, team); err != nil {
			return err
		}

		for _, rawPerson := range rawTeam.Squad {
			player := models.Person{
				ID:          rawPerson.ID,
				Name:        rawPerson.Name,
				Position:    &rawPerson.Position,
				DateOfBirth: rawPerson.DateOfBirth,
				Nationality: rawPerson.Nationality,
			}
			if err := repository.AddPerson(player); err != nil {
				return err
			}
			if err := repository.AddPersonToTeam(player, team); err != nil {
				return err
			}
		}
	}
}

func init() {
	client = clients.NewFootballAPIClient()
}