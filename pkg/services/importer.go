package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/clients"
	"github.com/ariel17/be-challenge-arios/pkg/models"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

type ImporterService interface {
	// ImportDataByCompetitionCode
	// TODO create and associate a ticket ID for future checks on process status
	ImportDataByCompetitionCode(code string) error
}

func NewImporterService(apiClient clients.FootballAPIClient, repository repositories.Repository) ImporterService {
	return &realImporterService{
		Client:     apiClient,
		Repository: repository,
	}
}

type realImporterService struct {
	Client     clients.FootballAPIClient
	Repository repositories.Repository
}

func (r *realImporterService) ImportDataByCompetitionCode(code string) error {
	rawCompetition, err := r.Client.GetLeagueByCode(code)
	if err != nil {
		return err
	}
	competition := models.Competition{
		Code:     rawCompetition.Code,
		Name:     rawCompetition.Name,
		AreaName: rawCompetition.Area.Name,
	}
	if err := r.Repository.AddCompetition(competition); err != nil {
		return err
	}

	rawTeams, err := r.Client.GetTeamsByLeagueCode(code)
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
		if err := r.Repository.AddTeam(team); err != nil {
			return err
		}
		if err := r.Repository.AddTeamToCompetition(team, competition); err != nil {
			return err
		}

		if len(rawTeam.Squad) > 0 {
			for _, rawPerson := range rawTeam.Squad {
				player := models.Person{
					ID:          rawPerson.ID,
					Name:        rawPerson.Name,
					Position:    &rawPerson.Position,
					DateOfBirth: rawPerson.DateOfBirth,
					Nationality: rawPerson.Nationality,
				}
				if err := r.Repository.AddPerson(player); err != nil {
					return err
				}
				if err := r.Repository.AddPersonToTeam(player, team); err != nil {
					return err
				}
			}
		} else {
			coach := models.Person{
				ID:          rawTeam.Coach.ID,
				Name:        rawTeam.Coach.Name,
				DateOfBirth: rawTeam.Coach.DateOfBirth,
				Nationality: rawTeam.Coach.Nationality,
			}
			if err := r.Repository.AddPerson(coach); err != nil {
				return err
			}
			if err := r.Repository.AddPersonToTeam(coach, team); err != nil {
				return err
			}
		}
	}
	return nil
}