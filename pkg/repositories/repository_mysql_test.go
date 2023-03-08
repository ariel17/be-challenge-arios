package repositories

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/be-challenge-arios/pkg/models"
)

func TestMysqlRepository_AddPerson(t *testing.T) {
	person := models.Person{
		ID:          1,
		Name:        "Ariel",
		DateOfBirth: "1983-02-17",
		Nationality: "Argentinian",
	}

	testCases := []struct {
		name      string
		isSuccess bool
	}{
		{"ok", true},
		{"failed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedExec := mock.ExpectExec("INSERT IGNORE INTO persons").
				WithArgs(person.ID, person.Name, person.DateOfBirth, person.Nationality)

			if tc.isSuccess {
				expectedExec.WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				expectedExec.WillReturnError(errors.New("some error"))
			}

			err := r.AddPerson(person)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_AddTeam(t *testing.T) {
	team := models.Team{
		TLA:       "ABC",
		Name:      "Test team",
		ShortName: "testteam",
		AreaName:  "Argentina",
		Address:   "Evergreen 123",
	}

	testCases := []struct {
		name      string
		isSuccess bool
	}{
		{"ok", true},
		{"failed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedExec := mock.ExpectExec("INSERT IGNORE INTO teams").
				WithArgs(team.TLA, team.Name, team.ShortName, team.AreaName, team.Address)

			if tc.isSuccess {
				expectedExec.WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				expectedExec.WillReturnError(errors.New("some error"))
			}

			err := r.AddTeam(team)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_AddCompetition(t *testing.T) {
	competition := models.Competition{
		Code:     "ABC",
		Name:     "Test Competition",
		AreaName: "Argentina",
	}

	testCases := []struct {
		name      string
		isSuccess bool
	}{
		{"ok", true},
		{"failed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedExec := mock.ExpectExec("INSERT IGNORE INTO competitions").
				WithArgs(competition.Code, competition.Name, competition.AreaName)

			if tc.isSuccess {
				expectedExec.WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				expectedExec.WillReturnError(errors.New("some error"))
			}

			err := r.AddCompetition(competition)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_AddTeamToCompetition(t *testing.T) {
	team := models.Team{
		TLA: "ABC",
	}
	competition := models.Competition{
		Code: "ABC",
	}

	testCases := []struct {
		name      string
		isSuccess bool
	}{
		{"ok", true},
		{"failed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedExec := mock.ExpectExec("INSERT IGNORE INTO competitions_teams").
				WithArgs(competition.Code, team.TLA)

			if tc.isSuccess {
				expectedExec.WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				expectedExec.WillReturnError(errors.New("some error"))
			}

			err := r.AddTeamToCompetition(team, competition)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_AddPersonToTeam(t *testing.T) {
	person1 := models.Person{
		ID: 1,
	}
	position := "Center"
	person2 := models.Person{
		ID:       1,
		Position: &position,
	}
	team := models.Team{
		TLA: "ABC",
	}

	testCases := []struct {
		name      string
		person    models.Person
		isSuccess bool
	}{
		{"ok without position", person1, true},
		{"ok with position", person2, true},
		{"failed", person1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedExec := mock.ExpectExec("INSERT IGNORE INTO teams_persons").
				WithArgs(team.TLA, tc.person.ID, tc.person.Position)

			if tc.isSuccess {
				expectedExec.WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				expectedExec.WillReturnError(errors.New("some error"))
			}

			err := r.AddPersonToTeam(tc.person, team)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_GetTeamByTLA(t *testing.T) {
	team := models.Team{
		TLA:       "ABC",
		Name:      "Test Team",
		ShortName: "TestTeam",
		AreaName:  "Argentina",
		Address:   "Evergreen 123",
	}

	testCases := []struct {
		name      string
		isSuccess bool
		isFound   bool
	}{
		{"ok", true, true},
		{"not found", true, false},
		{"failed", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedQuery := mock.ExpectQuery("SELECT name, short_name, area_name, address FROM teams WHERE tla = ").
				WithArgs(team.TLA)

			if tc.isSuccess {
				rows := sqlmock.NewRows([]string{"name", "short_name", "area_name", "address"})
				if tc.isFound {
					rows.AddRow(team.Name, team.ShortName, team.AreaName, team.Address)
				}
				expectedQuery.WillReturnRows(rows)
			} else {
				expectedQuery.WillReturnError(errors.New("some error"))
			}

			result, err := r.GetTeamByTLA("ABC")
			assert.Equal(t, result != nil, tc.isSuccess && tc.isFound)
			assert.Equal(t, err == nil, tc.isSuccess)
			if tc.isSuccess && tc.isFound {
				b1, _ := json.Marshal(team)
				b2, _ := json.Marshal(result)
				assert.Equal(t, b1, b2)
			}

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_GetPersonsByCompetitionCode(t *testing.T) {
	competition := models.Competition{
		Code: "ABC",
	}
	person1 := models.Person{
		ID:          1,
		Name:        "Person 1",
		DateOfBirth: "1980-01-01",
		Nationality: "Argentinian",
	}
	position := "Center"
	person2 := models.Person{
		ID:          2,
		Name:        "Person 2",
		DateOfBirth: "1980-01-01",
		Nationality: "Argentinian",
		Position:    &position,
	}

	testCases := []struct {
		name      string
		isSuccess bool
		isFound   bool
	}{
		{"ok", true, true},
		{"not found", true, false},
		{"failed", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedQuery := mock.ExpectQuery("SELECT p.id, p.name, p.date_of_birth, p.nationality, tp.position ").
				WithArgs(competition.Code)

			if tc.isSuccess {
				rows := sqlmock.NewRows([]string{"p.id", "p.name", "p.date_of_birth", "p.nationality", "tp.position"})
				if tc.isFound {
					rows.AddRow(person1.ID, person1.Name, person1.DateOfBirth, person1.Nationality, person1.Position).
						AddRow(person2.ID, person2.Name, person2.DateOfBirth, person2.Nationality, person2.Position)
				}
				expectedQuery.WillReturnRows(rows)
			} else {
				expectedQuery.WillReturnError(errors.New("some error"))
			}

			persons, err := r.GetPlayersByCompetitionCode(competition.Code, "")
			assert.Equal(t, len(persons) == 2, tc.isSuccess && tc.isFound)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_GetPersonsByTeamTLA(t *testing.T) {
	team := models.Team{
		TLA: "ABC",
	}
	person1 := models.Person{
		ID:          1,
		Name:        "Person 1",
		DateOfBirth: "1980-01-01",
		Nationality: "Argentinian",
	}
	position := "Center"
	person2 := models.Person{
		ID:          2,
		Name:        "Person 2",
		DateOfBirth: "1980-01-01",
		Nationality: "Argentinian",
		Position:    &position,
	}

	testCases := []struct {
		name      string
		isSuccess bool
		isFound   bool
	}{
		{"ok", true, true},
		{"not found", true, false},
		{"failed", false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedQuery := mock.ExpectQuery("SELECT p.id, p.name, p.date_of_birth, p.nationality, tp.position ").
				WithArgs(team.TLA)

			if tc.isSuccess {
				rows := sqlmock.NewRows([]string{"p.id", "p.name", "p.date_of_birth", "p.nationality", "tp.position"})
				if tc.isFound {
					rows.AddRow(person1.ID, person1.Name, person1.DateOfBirth, person1.Nationality, person1.Position).
						AddRow(person2.ID, person2.Name, person2.DateOfBirth, person2.Nationality, person2.Position)
				}
				expectedQuery.WillReturnRows(rows)
			} else {
				expectedQuery.WillReturnError(errors.New("some error"))
			}

			persons, err := r.GetPersonsByTeamTLA(team.TLA)
			assert.Equal(t, len(persons) == 2, tc.isSuccess && tc.isFound)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}

func TestMysqlRepository_CompetitionExists(t *testing.T) {
	code := "ABC"
	testCases := []struct {
		name      string
		isSuccess bool
		found     int
	}{
		{"ok", true, 1},
		{"not found", true, 0},
		{"failed", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			r := &mysqlRepository{db: db}

			expectedQuery := mock.ExpectQuery("SELECT COUNT").
				WithArgs(code)

			if tc.isSuccess {
				rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(tc.found)
				expectedQuery.WillReturnRows(rows)
			} else {
				expectedQuery.WillReturnError(errors.New("some error"))
			}

			exists, err := r.CompetitionExists(code)
			assert.Equal(t, exists, tc.isSuccess && tc.found > 0)
			assert.Equal(t, err == nil, tc.isSuccess)

			err = mock.ExpectationsWereMet()
			assert.Nil(t, err)
		})
	}
}