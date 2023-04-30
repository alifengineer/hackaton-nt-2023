CREATE TABLE IF NOT EXISTS "team"(
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "team_name" VARCHAR,
  "image" VARCHAR
);

CREATE TABLE IF NOT EXISTS "league"(
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "league_name" VARCHAR,
  "image" VARCHAR
);

CREATE TABLE IF NOT EXISTS "season" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "season_year_from" INTEGER,
  "season_year_to" INTEGER,
  "league_id" UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS "team_league_season" (
  "id"  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "league_id" UUID NOT NULL,
  "season_id" UUID NOT NULL,
  "team_id" UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS "match" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "home_team_id" uuid NOT NULL,
  "away_team_id" uuid NOT NULL,
  "league_id" uuid NOT NULL,
  "season_id" uuid NOT NULL,
  "home_score" INTEGER,
  "away_score" INTEGER,
  "tur" INTEGER DEFAULT 1,
  "m_date" TIMESTAMP DEFAULT NULL,
  "status" VARCHAR DEFAULT 'not_started'
  -- not_started, in_play, finished
);

CREATE TABLE IF NOT EXISTS "tur" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "soni" INTEGER,
  "league_id" UUID NOT NULL,
  "current_tur" INTEGER CHECK (current_tur >= 0 AND current_tur <= soni)
);

CREATE UNIQUE INDEX IF NOT EXISTS "tur_league_id_unique" ON tur (league_id);

CREATE UNIQUE INDEX IF NOT EXISTS "team_league_season_league_id_season_id_team_id_key" ON "team_league_season" ("league_id", "season_id", "team_id");

CREATE UNIQUE INDEX IF NOT EXISTS "tur_home_team_id_away_team_id_league_id_season_id_key" ON "match" ("tur", "home_team_id", "away_team_id", "league_id", "season_id");

ALTER TABLE "season" ADD CONSTRAINT fk_Season_LeagueID FOREIGN KEY ("league_id") REFERENCES "league" ("id");

ALTER TABLE team_league_season ADD CONSTRAINT fk_Team_League_Season_TeamID FOREIGN KEY ("team_id") REFERENCES "team" ("id");

ALTER TABLE "team_league_season" ADD CONSTRAINT fk_Team_League_Season_LeagueID FOREIGN KEY ("league_id") REFERENCES "league" ("id");

ALTER TABLE "team_league_season" ADD CONSTRAINT fk_Team_League_Season_SeasonID FOREIGN KEY ("season_id") REFERENCES "season" ("id");

ALTER TABLE match ADD CONSTRAINT fk_Match_HomeTeamID FOREIGN KEY ("home_team_id") REFERENCES "team" ("id");

ALTER TABLE match ADD CONSTRAINT fk_Match_AwayTeamID FOREIGN KEY ("away_team_id") REFERENCES "team" ("id");

ALTER TABLE match ADD CONSTRAINT fk_Match_LeagueID FOREIGN KEY ("league_id") REFERENCES "league" ("id");

ALTER TABLE match ADD CONSTRAINT fk_Match_SeasonID
FOREIGN KEY ("season_id") REFERENCES "season" ("id");
