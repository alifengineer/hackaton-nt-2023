package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dilmurodov/xakaton_nt/pkg/helper"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type seasonRepo struct {
	db *sql.DB
}

func NewSeasonRepo(db *sql.DB) storage.SeasonRepoI {
	return &seasonRepo{
		db: db,
	}
}

func (s *seasonRepo) CreateSeason(ctx context.Context, req *models.Season) (resp *models.Season, err error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	query := `
		INSERT INTO season(
			season_year_from,
			season_year_to,
			league_id,
		) VALUES (
			$1,
			$2,
			$3
		) RETURNING id, season_year_from, season_year_to, league_id`

	err = tx.QueryRowContext(ctx, query,
		time.Now().Year(),
		time.Now().Year()+1,
		req.LeagueID,
		true).Scan(
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

func (s *seasonRepo) GetSeasonByID(ctx context.Context, req *models.Season) (resp *models.Season, err error) {

	resp = &models.Season{}

	query := `
		SELECT
			id,
			season_year_from,
			season_year_to,
			league_id
		FROM season
		WHERE id = $1`
	err = s.db.QueryRowContext(ctx, query, req.ID).Scan(
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

func (s *seasonRepo) GetAllSeasons(ctx context.Context, req *models.GetAllSeasonsRequest) (resp []*models.Season, err error) {

	resp = []*models.Season{}

	query := `
		SELECT
			id,
			season_year_from,
			season_year_to,
			league_id
		FROM season`

	filter := ` WHERE 1=1`
	params := map[string]interface{}{}

	if req.LeagueID != "" {
		filter += ` AND league_id = :league_id`
		params["league_id"] = req.LeagueID
	}

	query += filter

	query, args := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		r := &models.Season{}
		err = rows.Scan(
			&r.ID,
			&r.SeasonFrom,
			&r.SeasonTo,
			&r.LeagueID,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}

	return resp, err
}
