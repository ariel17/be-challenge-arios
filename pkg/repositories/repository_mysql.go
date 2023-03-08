package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ariel17/be-challenge-arios/pkg/configs"
	"github.com/ariel17/be-challenge-arios/pkg/models"
)

const (
	personsSchema = `CREATE TABLE IF NOT EXISTS persons (
    id INT unsigned,
    name VARCHAR(50),
    date_of_birth CHAR(10),
    nationality VARCHAR(20),
    PRIMARY KEY (id)
)`

	teamsSchema = `CREATE TABLE IF NOT EXISTS teams (
    tla CHAR(3),
    name VARCHAR(50),
    short_name VARCHAR(100),
    area_name VARCHAR(50),
    address VARCHAR(200),
    PRIMARY KEY (tla)
)`

	teamsPersonsSchema = `CREATE TABLE IF NOT EXISTS teams_persons (
    team_tla CHAR(3),
    person_id INT unsigned,
    position VARCHAR(20) NULL,
    CONSTRAINT uc_person_by_team UNIQUE (team_tla, person_id),
    FOREIGN KEY (team_tla) REFERENCES teams (tla),
    FOREIGN KEY (person_id) REFERENCES persons (id)
)`

	competitionsSchema = `CREATE TABLE IF NOT EXISTS competitions (
    code CHAR(4),
    name VARCHAR(50),
    area_name VARCHAR(50),
    PRIMARY KEY (code)
)`

	competitionsTeamsSchema = `CREATE TABLE IF NOT EXISTS competitions_teams (
    competition_code CHAR(4),
    team_tla CHAR(3),
    CONSTRAINT uc_team_by_competition UNIQUE (competition_code, team_tla),
    FOREIGN KEY (competition_code) REFERENCES competitions (code),
    FOREIGN KEY (team_tla) REFERENCES teams (tla)
)`
)

func NewMySQLRepository() Repository {
	return &mysqlRepository{}
}

type mysqlRepository struct {
	db *sql.DB
}

func (m *mysqlRepository) Connect() error {
	db, err := sql.Open("mysql", configs.GetDSN())
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *mysqlRepository) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

func (m *mysqlRepository) GetStatus() error {
	_, err := m.db.Query(configs.GetStatusQuery())
	return err
}

func (m *mysqlRepository) CreateSchema() error {
	for _, schema := range []string{personsSchema, teamsSchema, teamsPersonsSchema, competitionsSchema, competitionsTeamsSchema} {
		if _, err := m.db.Exec(schema); err != nil {
			return err
		}
	}
	return nil
}

func (m *mysqlRepository) AddPerson(person models.Person) error {
	query := "INSERT IGNORE INTO persons (id, name, date_of_birth, nationality) VALUES (?, ?, ?, ?)"
	_, err := m.db.Exec(query, person.ID, person.Name, person.DateOfBirth, person.Nationality)
	return err
}

func (m *mysqlRepository) AddTeam(team models.Team) error {
	query := "INSERT IGNORE INTO teams (tla, name, short_name, area_name, address) VALUES (?, ?, ?, ?, ?)"
	_, err := m.db.Exec(query, team.TLA, team.Name, team.ShortName, team.AreaName, team.Address)
	return err
}

func (m *mysqlRepository) AddCompetition(competition models.Competition) error {
	query := "INSERT IGNORE INTO competitions (code, name, area_name) VALUES (?, ?, ?)"
	_, err := m.db.Exec(query, competition.Code, competition.Name, competition.AreaName)
	return err
}

func (m *mysqlRepository) AddTeamToCompetition(team models.Team, competition models.Competition) error {
	query := "INSERT IGNORE INTO competitions_teams (competition_code, team_tla) VALUES (?, ?)"
	_, err := m.db.Exec(query, competition.Code, team.TLA)
	return err
}

func (m *mysqlRepository) AddPersonToTeam(person models.Person, team models.Team) error {
	query := "INSERT IGNORE INTO teams_persons (team_tla, person_id, position) VALUES (?, ?, ?)"
	_, err := m.db.Exec(query, team.TLA, person.ID, person.Position)
	return err
}

func (m *mysqlRepository) GetTeamByTLA(tla string) (*models.Team, error) {
	query := "SELECT name, short_name, area_name, address FROM teams WHERE tla = ?"
	result, err := m.db.Query(query, tla)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	if !result.Next() {
		return nil, nil
	}

	team := models.Team{TLA: tla}
	err = result.Scan(&team.Name, &team.ShortName, &team.AreaName, &team.Address)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (m *mysqlRepository) GetPlayersByCompetitionCode(code, teamNameToFilter string) ([]models.Person, error) {
	query := `SELECT p.id, p.name, p.date_of_birth, p.nationality, tp.position
FROM competitions c
INNER JOIN competitions_teams ct ON (c.code=ct.competition_code)
INNER JOIN teams t ON (ct.team_tla=t.tla)
INNER JOIN teams_persons tp ON (t.tla=tp.team_tla)
INNER JOIN persons p ON (tp.person_id=p.id)
WHERE c.code = ? AND tp.position IS NOT NULL`

	var (
		result *sql.Rows
		err    error
	)
	if teamNameToFilter != "" {
		query += " AND t.name LIKE ?"
		teamNameToFilter = "%" + teamNameToFilter + "%"
		result, err = m.db.Query(query, code, teamNameToFilter)
	} else {
		result, err = m.db.Query(query, code)
	}
	if err != nil {
		return nil, err
	}

	defer result.Close()
	persons := []models.Person{}
	for result.Next() {
		var (
			position sql.NullString
			person   models.Person
		)
		err = result.Scan(&person.ID, &person.Name, &person.DateOfBirth, &person.Nationality, &position)
		if err != nil {
			return nil, err
		}
		if position.Valid {
			person.Position = &position.String
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (m *mysqlRepository) GetPersonsByTeamTLA(tla string) ([]models.Person, error) {
	query := `SELECT p.id, p.name, p.date_of_birth, p.nationality, tp.position
FROM teams t
INNER JOIN teams_persons tp ON (t.tla=tp.team_tla)
INNER JOIN persons p ON (tp.person_id=p.id)
WHERE t.tla = ?`
	result, err := m.db.Query(query, tla)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	persons := []models.Person{}
	for result.Next() {
		var (
			position sql.NullString
			person   models.Person
		)
		err = result.Scan(&person.ID, &person.Name, &person.DateOfBirth, &person.Nationality, &position)
		if err != nil {
			return nil, err
		}
		if position.Valid {
			person.Position = &position.String
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (m *mysqlRepository) CompetitionExists(code string) (bool, error) {
	query := "SELECT COUNT(*) FROM competitions WHERE code = ?"
	result, err := m.db.Query(query, code)
	if err != nil {
		return false, err
	}

	defer result.Close()
	var count int
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return false, err
		}
	}

	return count > 0, nil
}