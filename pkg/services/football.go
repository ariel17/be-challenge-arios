package services

import (
	"github.com/ariel17/be-challenge-arios/pkg/models"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

type FootballService interface {
	GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, bool, error)
	GetTeamByTLA(tla string, withPlayers bool) (*models.Team, error)
	GetPersonsByTeamTLA(tla string) ([]models.Person, error)
}

func NewFootballService(repository repositories.Repository) FootballService {
	return &realFootballService{
		Repository: repository,
	}
}

type realFootballService struct {
	Repository repositories.Repository
}

func (r *realFootballService) GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, bool, error) {
	exists, err := r.Repository.CompetitionExists(code)
	if err != nil {
		return nil, false, err
	}
	if !exists {
		return []models.Person{}, false, nil
	}
	players, err := r.Repository.GetPlayersByCompetitionCode(code, teamNameToFilter)
	if err != nil {
		return nil, false, err
	}
	return players, true, nil
}

func (r *realFootballService) GetTeamByTLA(tla string, withPlayers bool) (*models.Team, error) {
	team, err := r.Repository.GetTeamByTLA(tla)
	if err != nil {
		return nil, err
	}
	if withPlayers {
		persons, err := r.Repository.GetPersonsByTeamTLA(tla)
		if err != nil {
			return nil, err
		}
		team.Players = []models.Person{}
		for _, p := range persons {
			if p.IsPlayer() {
				team.Players = append(team.Players, p)
			} else {
				team.Coach = &p
			}
		}
	}
	return team, nil
}

func (r *realFootballService) GetPersonsByTeamTLA(tla string) ([]models.Person, error) {
	return r.Repository.GetPersonsByTeamTLA(tla)
}