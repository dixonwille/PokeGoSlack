﻿CREATE SCHEMA IF NOT EXISTS system;
GRANT USAGE ON SCHEMA system TO current_user;
GRANT CREATE ON SCHEMA system TO current_user;
CREATE TABLE IF NOT EXISTS system.Team(
	TeamId		VARCHAR(50) PRIMARY KEY,
	TeamName	VARCHAR(50) NOT NULL,
	PokeTeam 	SMALLINT DEFAULT 0,
	AccessToken VARCHAR(100) UNIQUE
);
CREATE TABLE IF NOT EXISTS system.Trainer(
	TrainerId 	VARCHAR(50) PRIMARY KEY,
	TeamId		VARCHAR(50) NOT NULL REFERENCES system.Team(TeamId),
	UserName  	VARCHAR(21) NOT NULL,
	VerifiedTeam 	SMALLINT DEFAULT 0,
	VerifiedBy	VARCHAR(50) REFERENCES system.Trainer(TrainerId),
	Updated		TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE TABLE IF NOT EXISTS system.Gym (
	GymId		BIGSERIAL PRIMARY KEY,
	TeamId		VARCHAR(50) NOT NULL REFERENCES system.Team(TeamId),
	GymName		VARCHAR(50) NOT NULL,
	PokeTeam 	SMALLINT DEFAULT 0,
	GymLevel	SMALLINT DEFAULT 0,
	UpdatedBy	VARCHAR(50) REFERENCES system.Trainer(TrainerId),
	Updated		TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc')
);
