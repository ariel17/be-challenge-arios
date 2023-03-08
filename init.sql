CREATE DATABASE challenge;
USE challenge;

CREATE TABLE persons (
    id INT unsigned,
    name VARCHAR(50),
    date_of_birth CHAR(10),
    nationality VARCHAR(20),
    PRIMARY KEY (id)
);

CREATE TABLE teams (
    tla CHAR(3),
    name VARCHAR(50),
    short_name VARCHAR(100),
    area_name VARCHAR(50),
    address VARCHAR(200),
    PRIMARY KEY (tla)
);

CREATE TABLE teams_persons (
    team_tla CHAR(3),
    person_id INT unsigned,
    position VARCHAR(20) NULL,
    CONSTRAINT uc_person_by_team UNIQUE (team_tla, person_id),
    FOREIGN KEY (team_tla) REFERENCES teams (tla),
    FOREIGN KEY (person_id) REFERENCES persons (id)
);

CREATE TABLE competitions (
    code CHAR(4),
    name VARCHAR(50),
    area_name VARCHAR(50),
    PRIMARY KEY (code)
);

CREATE TABLE competitions_teams (
    competition_code CHAR(4),
    team_tla CHAR(3),
    CONSTRAINT uc_team_by_competition UNIQUE (competition_code, team_tla),
    FOREIGN KEY (competition_code) REFERENCES competitions (code),
    FOREIGN KEY (team_tla) REFERENCES teams (tla)
);