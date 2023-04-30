package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/helper"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type teamRepo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) storage.TeamRepoI {
	return &teamRepo{
		db: db,
	}
}

func (s *teamRepo) CreateTeam(ctx context.Context, req *models.Team) (resp *models.Team, err error) {
	query := `
		INSERT INTO team(
			team_name,
			image
		) VALUES (
			$1,
			$2
		) RETURNING id, team_name, image`
	err = s.db.QueryRowContext(ctx, query,
		req.Name,
		req.Image).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Image,
	)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *teamRepo) GetTeamByID(ctx context.Context, req *models.GetTeamByIDRequest) (resp *models.Team, err error) {

	resp = &models.Team{}

	query := `
		SELECT
			id,
			team_name,
			COALESCE(image, '') as image
		FROM team
		WHERE id = $1`
	err = s.db.QueryRowContext(ctx, query, req.ID).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Image,
	)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *teamRepo) GetAllTeams(ctx context.Context, req *models.GetAllTeamsRequest) (resp *models.GetAllTeamsResponse, err error) {

	resp = &models.GetAllTeamsResponse{
		Teams: []*models.Team{},
	}

	query := `
		SELECT
			id,
			team_name,
			COALESCE(image, '') as image
		FROM team`

	filter := ` WHERE 1=1 `
	params := map[string]interface{}{}

	if req.LeagueID != "" {
		filter += ` AND league_id = :league_id `
		params["league_id"] = req.LeagueID
	}
	if req.Search != "" {
		filter += ` AND team_name LIKE :search `
		params["search"] = "%" + req.Search + "%"
	}

	query += filter

	query, args := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		team := &models.Team{}
		err = rows.Scan(
			&team.ID,
			&team.Name,
			&team.Image,
		)
		if err != nil {
			return nil, err
		}
		resp.Teams = append(resp.Teams, team)
	}

	return resp, err
}
