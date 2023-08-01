CREATE SCHEMA IF NOT EXISTS gamescoring;
-- CREATE TABLE
DROP TABLE IF EXISTS gamescoring.games;
CREATE TABLE gamescoring.games (
    id VARCHAR NOT NULL,
    start_time timestamptz NOT NULL,
    end_time timestamptz NOT NULL,
    arrive timestamptz NOT NULL,
    created_on timestamptz NOT NULL DEFAULT now(),
	updated_on timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT games_pkey PRIMARY KEY (id)
);

DROP TABLE IF EXISTS gamescoring.scoring_events;
CREATE TABLE gamescoring.scoring_events (
    id VARCHAR NOT NULL,
    game_id VARCHAR NOT NULL,
    timestamp timestamptz NOT NULL,
    "data" jsonb NOT NULL,
    created_on timestamptz NOT NULL DEFAULT now(),
	updated_on timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT scoring_events_pkey PRIMARY KEY (id)
);