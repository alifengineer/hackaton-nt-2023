package match

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
)

func (s *Service) GetMatchesByTURStatus(ctx context.Context, req *models.GetMatchesByTURStatusRequest) (resp *models.GetMatchesByTURStatusResponse, err error) {

	resp, err = s.strg.Match().GetMatchesByTURStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetAllLeagues(ctx context.Context) (resp []*models.League, err error) {

	resp, err = s.strg.League().GetAllLeagues(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetTopTeamsInLeague(ctx context.Context, req *models.GetTopTeamsInLeagueRequest) (resp *models.GetTopTeamsInLeagueResponse, err error) {

	resp, err = s.strg.League().GetTopTeamsInLeague(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) CreateMatch(ctx context.Context, req *models.Match) (resp *models.Match, err error) {

	resp, err = s.strg.Match().CreateMatch(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetAllMatch(ctx context.Context, req *models.GetAllMatchRequest) (resp *models.GetAllMatchResponse, err error) {

	resp, err = s.strg.Match().GetAllMatch(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetMatchByID(ctx context.Context, req *models.GetMatchByIDRequest) (resp *models.GetMatchByIDResponse, err error) {

	resp, err = s.strg.Match().GetMatchByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetLeagueByID(ctx context.Context, req *models.GetLeagueByIDRequest) (resp *models.GetLeagueByIDResponse, err error) {

	resp, err = s.strg.League().GetLeagueByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetAllTeams(ctx context.Context, req *models.GetAllTeamsRequest) (resp *models.GetAllTeamsResponse, err error) {

	resp, err = s.strg.Team().GetAllTeams(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetTeamByID(ctx context.Context, req *models.GetTeamByIDRequest) (resp *models.Team, err error) {

	resp, err = s.strg.Team().GetTeamByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetAllSeasons(ctx context.Context, req *models.GetAllSeasonsRequest) (resp []*models.Season, err error) {

	resp, err = s.strg.Season().GetAllSeasons(ctx, req)
	if err != nil {
		return nil, err
	}

	return
}
