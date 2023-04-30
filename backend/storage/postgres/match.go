package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dilmurodov/xakaton_nt/pkg/helper"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type matchRepo struct {
	db *sql.DB
}

func NewMatchRepo(db *sql.DB) storage.MatchRepoI {
	return &matchRepo{
		db: db,
	}
}

func (s *matchRepo) GetMatchesByTURStatus(ctx context.Context, req *models.GetMatchesByTURStatusRequest) (resp *models.GetMatchesByTURStatusResponse, err error) {

	query := `
	WITH matches (id, home_team_id, away_team_id, m_date, home_score, away_score) AS (
		SELECT
			m.id,
			m.home_team_id,
			m.away_team_id,
			m.m_date,
			m.home_score,
			m.away_score
		FROM match m
		WHERE m.league_id = $1 AND m.tur = $2
	),
	teams (id, team_name, image) AS (
		SELECT
			t.id,
			t.team_name,
			t.image
		FROM team t
	)
	SELECT
		m.id,
		m.m_date,
		m.home_score,
		m.away_score,
		ht.id,
		ht.team_name,
		ht.image,
		at.id,
		at.team_name,
		at.image,
		CASE
			WHEN m.home_score > m.away_score THEN ht.id
			WHEN m.home_score < m.away_score THEN at.id
			ELSE NULL
		END AS winner
	FROM matches m
	INNER JOIN teams ht ON ht.id = m.home_team_id
	INNER JOIN teams at ON at.id = m.away_team_id
	ORDER BY m.m_date ASC
	`

	rows, err := s.db.QueryContext(ctx, query,
		req.LeagueID,
		req.Tur,
	)
	if err != nil {
		return nil, err
	}

	resp = &models.GetMatchesByTURStatusResponse{
		Matches: make([]*models.TeamsWithResults, 0),
	}
	var (
		mID     sql.NullString
		mDate   sql.NullString
		htScore sql.NullInt64
		atScore sql.NullInt64
		htID    sql.NullString
		htName  sql.NullString
		htImage sql.NullString
		atID    sql.NullString
		atName  sql.NullString
		atImage sql.NullString
		winner  sql.NullString
	)

	for rows.Next() {
		var m models.TeamsWithResults
		var ht models.MatchTeam
		var at models.MatchTeam

		err = rows.Scan(

			&mID,
			&mDate,
			&htScore,
			&atScore,
			&htID,
			&htName,
			&htImage,
			&atID,
			&atName,
			&atImage,
			&winner,
		)
		if err != nil {
			return nil, err
		}

		m.ID = mID.String
		m.MDate = mDate.String
		ht.Score = int(htScore.Int64)
		at.Score = int(atScore.Int64)
		ht.ID = htID.String
		ht.Name = htName.String
		ht.Image = htImage.String
		at.ID = atID.String
		at.Name = atName.String
		at.Image = atImage.String
		m.Winner = winner.String

		m.HomeTeam = &ht
		m.AwayTeam = &at

		resp.Matches = append(resp.Matches, &m)
	}

	return resp, nil
}

func (s *matchRepo) CreateMatch(ctx context.Context, req *models.Match) (resp *models.Match, err error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			id
		FROM season
		WHERE date_part('year', CURRENT_TIMESTAMP) >= season_year_from AND date_part('year', CURRENT_TIMESTAMP) < season_year_to
	`
	err = tx.QueryRowContext(ctx, query).Scan(
		&req.SeasonID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			seassonResp, err := createSeason(tx, ctx, &models.Season{
				SeasonFrom: time.Now().Year(),
				SeasonTo:   time.Now().Year() + 1,
				LeagueID:   req.LeagueID,
				IsCurrent:  true,
			})
			if err != nil {
				return nil, err
			}
			req.SeasonID = seassonResp.ID
		} else {
			return nil, err
		}
	}

	resp = &models.Match{}
	query = `
	INSERT INTO match(
		home_team_id, 
		away_team_id, 
		league_id,
		season_id,
		tur, 
		m_date
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, home_team_id, away_team_id, m_date, tur, league_id, season_id`

	err = tx.QueryRowContext(ctx, query,
		req.HomeTeamID,
		req.AwayTeamID,
		req.LeagueID,
		req.SeasonID,
		req.Tur,
		req.MDate,
	).Scan(
		&resp.ID,
		&resp.HomeTeamID,
		&resp.AwayTeamID,
		&resp.MDate,
		&resp.Tur,
		&resp.LeagueID,
		&resp.SeasonID,
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *matchRepo) GetAllMatch(ctx context.Context, req *models.GetAllMatchRequest) (resp *models.GetAllMatchResponse, err error) {

	query := `
	WITH matches (id, home_team_id, away_team_id, m_date, home_score, away_score,
		league_id, season_id, tur, league_name, season_year_from, season_year_to) AS (
		SELECT
			m.id,
			m.home_team_id,
			m.away_team_id,
			m.m_date,
			m.home_score,
			m.away_score,
			m.league_id,
			m.season_id,
			m.tur,
			l.league_name,
			s.season_year_from,
			s.season_year_to
		FROM match m
		INNER JOIN league l ON l.id = m.league_id
		INNER JOIN season s ON s.id = m.season_id
		LIMIT :limit OFFSET :offset
	),
	teams (id, team_name, image) AS (
		SELECT
			t.id,
			t.team_name,
			t.image
		FROM team t
	)
	SELECT
		m.id,
		m.m_date,
		m.home_score,
		m.away_score,
		ht.id,
		ht.team_name,
		ht.image,
		at.id,
		at.team_name,
		at.image,
		CASE
			WHEN m.home_score > m.away_score THEN ht.id
			WHEN m.home_score < m.away_score THEN at.id
			ELSE NULL
		END AS winner,
		m.tur,
		m.league_id,
		m.season_id,
		m.league_name,
		m.season_year_from,
		m.season_year_to,
		COUNT(*) OVER() AS total_count
	FROM matches m
	INNER JOIN teams ht ON ht.id = m.home_team_id
	INNER JOIN teams at ON at.id = m.away_team_id
	`

	filter := " WHERE 1=1"
	order := " ORDER BY m.m_date DESC"
	params := map[string]interface{}{
		"limit":  req.Limit,
		"offset": req.Offset,
	}

	if req.LeagueID != "" {
		filter += ` AND m.league_id = :league_id"`
		params["league_id"] = req.LeagueID
	}
	if req.SeasonID != "" {
		filter += ` AND m.season_id = :season_id"`
		params["season_id"] = req.SeasonID
	}
	if req.Tur != 0 {
		filter += ` AND m.tur = :tur"`
		params["tur"] = req.Tur
	}
	if req.MDateFrom != "" && req.MDateTo != "" {
		filter += ` AND m.m_date BETWEEN :m_date_from AND :m_date_to`
		params["m_date_from"] = req.MDateFrom
		params["m_date_to"] = req.MDateTo
	}
	if req.Status != "" {
		filter += ` AND m.status = :status`
		params["status"] = req.Status
	}
	if req.Search != "" {
		filter += ` AND (ht.team_name ILIKE :search OR at.team_name ILIKE :search 
			OR m.league_name ILIKE :search)`
		params["search"] = "%" + req.Search + "%"
	}

	query = query + filter + order

	query, args := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.QueryContext(ctx, query,
		args...,
	)
	if err != nil {
		return nil, err
	}

	resp = &models.GetAllMatchResponse{
		Matches: make([]*models.TeamsWithResults, 0),
	}
	var (
		mID         sql.NullString
		mDate       sql.NullString
		htScore     sql.NullInt64
		atScore     sql.NullInt64
		htID        sql.NullString
		htName      sql.NullString
		htImage     sql.NullString
		atID        sql.NullString
		atName      sql.NullString
		atImage     sql.NullString
		winner      sql.NullString
		mTur        sql.NullInt64
		mLeagueName sql.NullString
		mLeagueID   sql.NullString
		mSeasonID   sql.NullString
		mSeasonFrom sql.NullInt64
		mSeasonTo   sql.NullInt64
		totalCount  sql.NullInt64
	)

	for rows.Next() {
		var m models.TeamsWithResults
		var ht models.MatchTeam
		var at models.MatchTeam

		err = rows.Scan(

			&mID,
			&mDate,
			&htScore,
			&atScore,
			&htID,
			&htName,
			&htImage,
			&atID,
			&atName,
			&atImage,
			&winner,
			&mTur,
			&mLeagueID,
			&mSeasonID,
			&mLeagueName,
			&mSeasonFrom,
			&mSeasonTo,
			&totalCount,
		)
		if err != nil {
			return nil, err
		}

		m.ID = mID.String
		m.MDate = mDate.String
		ht.Score = int(htScore.Int64)
		at.Score = int(atScore.Int64)
		ht.ID = htID.String
		ht.Name = htName.String
		ht.Image = htImage.String
		at.ID = atID.String
		at.Name = atName.String
		at.Image = atImage.String
		m.Winner = winner.String
		m.Tur = int(mTur.Int64)
		m.LeagueID = mLeagueID.String
		m.SeasonID = mSeasonID.String
		m.LeagueName = mLeagueName.String
		m.SeasonFrom = int(mSeasonFrom.Int64)
		m.SeasonTo = int(mSeasonTo.Int64)
		resp.TotalCount = int(totalCount.Int64)

		m.HomeTeam = &ht
		m.AwayTeam = &at

		resp.Matches = append(resp.Matches, &m)
	}

	return resp, nil
}

func (s *matchRepo) GetMatchByID(ctx context.Context, req *models.GetMatchByIDRequest) (
	resp *models.GetMatchByIDResponse, err error) {

	resp = &models.GetMatchByIDResponse{
		Match: &models.TeamsWithResults{},
	}

	query := `
	WITH matches (id, home_team_id, away_team_id, m_date, home_score, away_score,
		league_id, season_id, tur, league_name, season_year_from, season_year_to) AS (
		SELECT
			m.id,
			m.home_team_id,
			m.away_team_id,
			m.m_date,
			m.home_score,
			m.away_score,
			m.league_id,
			m.season_id,
			m.tur,
			l.league_name,
			s.season_year_from,
			s.season_year_to
		FROM match m
		INNER JOIN league l ON l.id = m.league_id
		INNER JOIN season s ON s.id = m.season_id
		WHERE m.id = $1
	),
	teams (id, team_name, image) AS (
		SELECT
			t.id,
			t.team_name,
			t.image
		FROM team t
	)
	SELECT
		m.id,
		m.m_date,
		m.home_score,
		m.away_score,
		ht.id,
		ht.team_name,
		ht.image,
		at.id,
		at.team_name,
		at.image,
		CASE
			WHEN m.home_score > m.away_score THEN ht.id
			WHEN m.home_score < m.away_score THEN at.id
			ELSE NULL
		END AS winner,
		m.tur,
		m.league_id,
		m.season_id,
		m.league_name,
		m.season_year_from,
		m.season_year_to
	FROM matches m
	INNER JOIN teams ht ON ht.id = m.home_team_id
	INNER JOIN teams at ON at.id = m.away_team_id
	`

	var (
		mID         sql.NullString
		mDate       sql.NullString
		htScore     sql.NullInt64
		atScore     sql.NullInt64
		htID        sql.NullString
		htName      sql.NullString
		htImage     sql.NullString
		atID        sql.NullString
		atName      sql.NullString
		atImage     sql.NullString
		winner      sql.NullString
		mTur        sql.NullInt64
		mLeagueName sql.NullString
		mLeagueID   sql.NullString
		mSeasonID   sql.NullString
		mSeasonFrom sql.NullInt64
		mSeasonTo   sql.NullInt64
	)

	err = s.db.QueryRowContext(ctx, query, req.ID).Scan(
		&mID,
		&mDate,
		&htScore,
		&atScore,
		&htID,
		&htName,
		&htImage,
		&atID,
		&atName,
		&atImage,
		&winner,
		&mTur,
		&mLeagueID,
		&mSeasonID,
		&mLeagueName,
		&mSeasonFrom,
		&mSeasonTo,
	)
	if err != nil {
		return nil, err
	}
	var m models.TeamsWithResults
	var ht models.MatchTeam
	var at models.MatchTeam

	m.ID = mID.String
	m.MDate = mDate.String
	ht.Score = int(htScore.Int64)
	at.Score = int(atScore.Int64)
	ht.ID = htID.String
	ht.Name = htName.String
	ht.Image = htImage.String
	at.ID = atID.String
	at.Name = atName.String
	at.Image = atImage.String
	m.Winner = winner.String
	m.Tur = int(mTur.Int64)
	m.LeagueID = mLeagueID.String
	m.SeasonID = mSeasonID.String
	m.LeagueName = mLeagueName.String
	m.SeasonFrom = int(mSeasonFrom.Int64)
	m.SeasonTo = int(mSeasonTo.Int64)

	m.HomeTeam = &ht
	m.AwayTeam = &at

	resp.Match = &m
	return resp, nil
}

func createSeason(tx *sql.Tx, ctx context.Context, req *models.Season) (resp *models.Season, err error) {

	query := `
		INSERT INTO season(
			season_year_from,
			season_year_to,
			league_id
		) VALUES (
			$1,
			$2,
			$3
		) RETURNING id, season_year_from, season_year_to, league_id`

	err = tx.QueryRowContext(ctx, query,
		req.SeasonFrom,
		req.SeasonTo,
		req.LeagueID).Scan(
		&resp.ID,
		&resp.SeasonFrom,
		&resp.SeasonTo,
		&resp.LeagueID,
	)
	if err != nil {
		return nil, err
	}
	return resp, err
}
