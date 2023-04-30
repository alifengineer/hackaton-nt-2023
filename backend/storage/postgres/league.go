package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/dilmurodov/xakaton_nt/pkg/helper"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (s *leagueRepo) GetLeagueSeasonTeams(ctx context.Context, req *models.GetLeagueSeasonTeamsRequest) (resp *models.GetLeagueSeasonTeamsResponse, err error) {
	resp = &models.GetLeagueSeasonTeamsResponse{
		Teams: []*models.Team{},
	}
	query := `
		SELECT
			tls.id,
			tls.league_id,
			tls.season_id,
			t.id AS team_id,
			t.team_name AS team_name,
			t.image AS team_image
		FROM team_league_season tls
		INNER JOIN team t ON t.id = tls.team_id
		WHERE tls.league_id = $1 AND tls.season_id = $2
	`

	rows, err := s.db.QueryContext(ctx, query, req.LeagueID, req.SeasonID)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, nil
		} else {
			return nil, err
		}
	}

	defer rows.Close()

	var (
		rID       sql.NullString
		rLeagueID sql.NullString
		rSeasonID sql.NullString
		rTeamID   sql.NullString
		rTeamName sql.NullString
		rTeamImg  sql.NullString
	)

	for rows.Next() {
		var r models.Team
		if err := rows.Scan(
			&rID,
			&rLeagueID,
			&rSeasonID,
			&rTeamID,
			&rTeamName,
			&rTeamImg,
		); err != nil {
			return nil, err
		}
		r.ID = rTeamID.String
		r.Name = rTeamName.String
		r.Image = rTeamImg.String
		resp.Teams = append(resp.Teams, &r)
	}

	return resp, err
}

func (s *leagueRepo) CreateLeagueSeasonTeams(ctx context.Context, req *models.CreateLeagueSeasonTeamsRequest) (err error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	if _, err := uuid.Parse(req.SeasonID); err != nil {
		season, err := createSeason(tx, ctx, &models.Season{
			SeasonFrom: time.Now().Year(),
			SeasonTo:   time.Now().Year() + 1,
			LeagueID:   req.LeagueID,
		})
		if err != nil {
			return err
		}
		req.SeasonID = season.ID
	}

	query := `
		INSERT INTO team_league_season (
			league_id, 
			season_id, 
			team_id
		) VALUES `

	var values []interface{}

	for _, team := range req.Teams {
		query += ` (?, ?, ?),`
		values = append(values, req.LeagueID, req.SeasonID, team.ID)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	if _, err := tx.ExecContext(ctx, query, values...); err != nil {
		return errors.Wrap(err, "error while inserting team_league_season")
	}

	return err
}
