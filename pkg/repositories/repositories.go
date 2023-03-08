package repositories

import "github.com/ariel17/be-challenge-arios/pkg/models"

type Repository interface {
	Connect() error
	Close() error
	GetStatus() error
	CreateSchema() error

	AddPerson(person models.Person) error
	AddTeam(team models.Team) error
	AddCompetition(competition models.Competition) error
	AddTeamToCompetition(team models.Team, competition models.Competition) error
	AddPersonToTeam(player models.Person, team models.Team) error

	GetTeamByTLA(tla string) (*models.Team, error)
	GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, error)
	GetPersonsByTeamTLA(tla string) ([]models.Person, error)
	CompetitionExists(code string) (bool, error)
}