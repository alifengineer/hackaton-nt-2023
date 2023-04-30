package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
	"github.com/google/uuid"
)

type leagueRepo struct {
	db *sql.DB
}

func NewLeagueRepo(db *sql.DB) storage.LeagueRepoI {
	return &leagueRepo{
		db: db,
	}
}

func (s *leagueRepo) CreateLeague(ctx context.Context, req *models.League) (resp *models.League, err error) {

	id := uuid.NewString()

	query := `
		INSERT INTO League(
			id, 
			LeagueName,
		) VALUES (
			$1,
			$2,
		) RETURNING id, LeagueName`
	err = s.db.QueryRowContext(ctx, query,
		id,
		req.Name).Scan(
		&resp.ID,
		&resp.Name,
	)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (s *leagueRepo) GetLeagueByID(ctx context.Context, req *models.GetLeagueByIDRequest) (resp *models.GetLeagueByIDResponse, err error) {
	query := `
	WITH tur (id, current_tur, soni) AS (
		SELECT
			id,
			current_tur,
			soni
		FROM tur
		WHERE league_id = $1
	), season (id, season_year_from, season_year_to) AS (
		SELECT
			id,
			season_year_from,
			season_year_to
		FROM season
		WHERE league_id = $1 AND date_part('year', CURRENT_TIMESTAMP) >= season_year_from AND date_part('year', CURRENT_TIMESTAMP) < season_year_to
	) SELECT
		l.id,
		l.league_name,
		COALESCE(l.image, '') as image,
		s.id,
		s.season_year_from,
		s.season_year_to,
		t.current_tur,
		t.soni
	FROM league l, tur t, season s
	WHERE l.id = $1
	`

	var (
		rID             sql.NullString
		rName           sql.NullString
		rImage          sql.NullString
		rSeasonYearFrom sql.NullInt64
		rSeasonYearTo   sql.NullInt64
		rCurrentTur     sql.NullInt64
		rSoni           sql.NullInt64
		rSeasonID       sql.NullString
	)

	err = s.db.QueryRowContext(ctx, query, req.ID).Scan(
		&rID,
		&rName,
		&rImage,
		&rSeasonID,
		&rSeasonYearFrom,
		&rSeasonYearTo,
		&rCurrentTur,
		&rSoni,
	)
	if err != nil {
		return nil, err
	}

	resp = &models.GetLeagueByIDResponse{
		ID:    rID.String,
		Name:  rName.String,
		Image: rImage.String,
		Season: &models.Season{
			ID:         rSeasonID.String,
			SeasonFrom: int(rSeasonYearFrom.Int64),
			SeasonTo:   int(rSeasonYearTo.Int64),
		},
		Tur: &models.Tur{
			CurrentTur: int(rCurrentTur.Int64),
			Soni:       int(rSoni.Int64),
		},
	}

	return resp, err
}

func (s *leagueRepo) GetAllLeagues(ctx context.Context) (resp []*models.League, err error) {
	resp = []*models.League{}
	query := `
		SELECT
			l.id,
			l.league_name,
			COALESCE(l.image, '') as image,
			COALESCE(t.current_tur, 0) as prev_tur,
			COALESCE(t.current_tur, 0) + 1 as next_tur
		FROM league l
		JOIN tur t ON t.league_id = l.id
		`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows != nil && rows.Next() {
		var r models.League
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Image,
			&r.PrevTur,
			&r.NextTur,
		); err != nil {
			return nil, err
		}
		resp = append(resp, &r)
	}
	return resp, err
}

func (s *leagueRepo) GetTopTeamsInLeague(ctx context.Context, req *models.GetTopTeamsInLeagueRequest) (resp *models.GetTopTeamsInLeagueResponse, err error) {

	resp = &models.GetTopTeamsInLeagueResponse{
		Teams: []*models.Team{},
	}

	query := `
		WITH league_teams(team_id, season_id) AS  (
			SELECT
				t.id AS team_id,
				tls.season_id AS season_id
			FROM team_league_season tls
			JOIN team t ON t.id = tls.team_id
			WHERE league_id = $1
			AND season_id = (
				SELECT
					id
				FROM season
				WHERE date_part('year', CURRENT_TIMESTAMP) >= season_year_from AND date_part('year', CURRENT_TIMESTAMP) < season_year_to
				AND league_id = $1
			)
		),
		-- every league teams's number of games played
		team_games_played(team_id, games_played) AS (
			SELECT
				t.id AS team_id,
				COUNT(m.id) AS games_played
			FROM team t
			LEFT JOIN match m ON m.home_team_id = t.id OR m.away_team_id = t.id
			WHERE t.id IN (SELECT team_id FROM league_teams)
			AND m.status = 'finished'
			GROUP BY t.id
		),
		-- every league teams's score
		team_score(team_id, score, scored, missed, difference) AS (
			SELECT
				t.id AS team_id,
				SUM(
					CASE
						WHEN m.home_team_id = t.id AND m.home_score > m.away_score THEN 3
						WHEN m.away_team_id = t.id AND m.away_score > m.home_score THEN 3
						WHEN m.home_score = m.away_score THEN 1
						ELSE 0
					END
				) AS score,
				SUM(
					CASE
						WHEN m.home_team_id = t.id THEN m.home_score
						ELSE m.away_score
					END
				) AS scored,
				SUM(
					CASE
						WHEN m.home_team_id = t.id THEN m.away_score
						ELSE m.home_score
					END
				) AS missed,
				SUM(
					CASE
						WHEN m.home_team_id = t.id THEN m.home_score
						ELSE m.away_score
					END
				) - SUM(
					CASE
						WHEN m.home_team_id = t.id THEN m.away_score
						ELSE m.home_score
					END
				) AS difference
			FROM team t
			LEFT JOIN match m ON m.home_team_id = t.id OR m.away_team_id = t.id
			WHERE t.id IN (SELECT team_id FROM league_teams)
			AND m.status = 'finished'
			GROUP BY t.id
		)
		SELECT
			t.id AS team_id,
			t.team_name AS team_name,
			t.image AS team_image,
			COALESCE(ts.score, 0) AS team_score,
			COALESCE(ts.scored, 0) AS team_scored,
			COALESCE(ts.missed, 0) AS team_missed,
			COALESCE(ts.difference, 0) AS team_difference,
			COALESCE(tgp.games_played, 0) AS team_games_played
		FROM team t
		LEFT JOIN team_score ts ON ts.team_id = t.id
		LEFT JOIN team_games_played tgp ON tgp.team_id = t.id
		WHERE t.id IN (SELECT team_id FROM league_teams)
		ORDER BY team_score DESC, team_difference DESC
	`

	rows, err := s.db.QueryContext(ctx, query, req.LeagueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		rID         sql.NullString
		rName       sql.NullString
		rImage      sql.NullString
		rScore      sql.NullInt64
		rScored     sql.NullInt64
		rMissed     sql.NullInt64
		rDifference sql.NullInt64
		rPlayed     sql.NullInt64
	)

	for rows.Next() {
		var r models.Team
		if err := rows.Scan(
			&rID,
			&rName,
			&rImage,
			&rScore,
			&rScored,
			&rMissed,
			&rDifference,
			&rPlayed,
		); err != nil {
			return nil, err
		}

		r.ID = rID.String
		r.Name = rName.String
		r.Image = rImage.String
		r.Score = int(rScore.Int64)
		r.Scored = int(rScored.Int64)
		r.Missed = int(rMissed.Int64)
		r.Difference = int(rDifference.Int64)
		r.Played = int(rPlayed.Int64)

		resp.Teams = append(resp.Teams, &r)
	}

	return resp, err
}

func (s *leagueRepo) GetTURByLeagueID(ctx context.Context, req *models.GetTURByLeagueIDRequest) (resp *models.GetTURByLeagueIDResponse, err error) {

	resp = &models.GetTURByLeagueIDResponse{
		TUR: &models.Tur{},
	}

	query := `
		SELECT
			id,
			league_id,
			soni,
			current_tur
		FROM tur
		WHERE league_id = $1
	`

	var (
		rID         sql.NullString
		rLeagueID   sql.NullString
		rSoni       sql.NullInt64
		rCurrentTUR sql.NullInt64
	)

	if err := s.db.QueryRowContext(ctx, query, req.LeagueID).Scan(
		&rID,
		&rLeagueID,
		&rSoni,
		&rCurrentTUR,
	); err != nil {
		if err == sql.ErrNoRows {
			return resp, nil
		}

		return nil, err
	}

	resp.TUR = &models.Tur{
		ID:         rID.String,
		LeagueID:   rLeagueID.String,
		CurrentTur: int(rCurrentTUR.Int64),
		Soni:       int(rSoni.Int64),
	}

	return resp, err
}
