package user

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type ServiceI interface {
	GetUserByID(context.Context, *models.GetUserByIDRequest) (*models.GetUserByIDResponse, error)
	CreateUser(context.Context, *models.CreateUserRequest) (*models.User, error)
	GetUserByCredentials(ctx context.Context, req *models.GetByCredentialsRequest) (*models.User, error)
	GetUserPasswordByEmail(ctx context.Context, email string) (resp *models.User, err error)
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
