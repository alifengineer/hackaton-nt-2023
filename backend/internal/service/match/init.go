package match

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type ServiceI interface {
	GetTopTeamsInLeague(ctx context.Context, req *models.GetTopTeamsInLeagueRequest) (resp *models.GetTopTeamsInLeagueResponse, err error)
	GetAllLeagues(ctx context.Context) (resp []*models.League, err error)
	GetMatchesByTURStatus(ctx context.Context, req *models.GetMatchesByTURStatusRequest) (resp *models.GetMatchesByTURStatusResponse, err error)
	CreateMatch(ctx context.Context, req *models.Match) (resp *models.Match, err error)
	GetAllMatch(ctx context.Context, req *models.GetAllMatchRequest) (resp *models.GetAllMatchResponse, err error)
	GetMatchByID(ctx context.Context, req *models.GetMatchByIDRequest) (resp *models.GetMatchByIDResponse, err error)
	GetLeagueByID(ctx context.Context, req *models.GetLeagueByIDRequest) (resp *models.GetLeagueByIDResponse, err error)
	GetAllTeams(ctx context.Context, req *models.GetAllTeamsRequest) (resp *models.GetAllTeamsResponse, err error)
	GetTeamByID(ctx context.Context, req *models.GetTeamByIDRequest) (resp *models.Team, err error) 
	GetAllSeasons(ctx context.Context, req *models.GetAllSeasonsRequest) (resp []*models.Season, err error)
}

type Service struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
}

func NewService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *Service {
	return &Service{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}
