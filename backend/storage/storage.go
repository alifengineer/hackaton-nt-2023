package storage

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Sms() SmsRepoI
	Email() EmailRepoI
	League() LeagueRepoI
	Season() SeasonRepoI
	Team() TeamRepoI
	Match() MatchRepoI
	News() NewsRepoI
}

type UserRepoI interface {
	GetUserByID(context.Context, *models.GetUserByIDRequest) (*models.GetUserByIDResponse, error)
	CreateUser(context.Context, *models.CreateUserRequest) (*models.User, error)
	GetUserPasswordByEmail(ctx context.Context, email string) (resp *models.User, err error)
}

type SmsRepoI interface {
	UpdateSMSAsSent(ctx context.Context, SMS_ID string) error
	IncrementSendCount(ctx context.Context, ID string) error
	SendOtp(ctx context.Context, req *models.Sms) (*models.SendOtpResponse, error)
	ConfirmOtp(ctx context.Context, req *models.ConfirmOtpRequest) (err error)
	GetNotSent(ctx context.Context) ([]*models.Sms, error)
}

type EmailRepoI interface {
	Create(ctx context.Context, input *models.Email) (*models.Email, error)
	GetByPK(ctx context.Context, pKey *models.GetEmailOtpByPK) (res *models.Email, err error)
}

type MatchRepoI interface {
	GetMatchesByTURStatus(ctx context.Context, req *models.GetMatchesByTURStatusRequest) (resp *models.GetMatchesByTURStatusResponse, err error)
	CreateMatch(ctx context.Context, req *models.Match) (resp *models.Match, err error)
	GetAllMatch(ctx context.Context, req *models.GetAllMatchRequest) (resp *models.GetAllMatchResponse, err error)
	GetMatchByID(ctx context.Context, req *models.GetMatchByIDRequest) (
		resp *models.GetMatchByIDResponse, err error)
}

type TeamRepoI interface {
	CreateTeam(ctx context.Context, input *models.Team) (*models.Team, error)
	GetTeamByID(ctx context.Context, req *models.GetTeamByIDRequest) (resp *models.Team, err error)
	GetAllTeams(ctx context.Context, req *models.GetAllTeamsRequest) (resp *models.GetAllTeamsResponse, err error)
}

type LeagueRepoI interface {
	CreateLeague(ctx context.Context, input *models.League) (*models.League, error)
	GetLeagueByID(ctx context.Context, req *models.GetLeagueByIDRequest) (resp *models.GetLeagueByIDResponse, err error)
	GetAllLeagues(ctx context.Context) ([]*models.League, error)
	GetTopTeamsInLeague(ctx context.Context, req *models.GetTopTeamsInLeagueRequest) (resp *models.GetTopTeamsInLeagueResponse, err error)
}

type SeasonRepoI interface {
	CreateSeason(ctx context.Context, req *models.Season) (resp *models.Season, err error)
	GetSeasonByID(ctx context.Context, req *models.Season) (*models.Season, error)
	GetAllSeasons(ctx context.Context, req *models.GetAllSeasonsRequest) (resp []*models.Season, err error)
}

type NewsRepoI interface {
	GetNewsByID(ctx context.Context, req *models.GetNewsByIDRequest) (
		resp *models.GetNewsByIDResponse, err error)
	GetAllNews(ctx context.Context, req *models.GetAllNewRequest) (resp *models.GetAllNewsResponse, err error)
	CreateNews(ctx context.Context, req *models.CreateNewsRequest) (resp *models.News, err error)
}
