package news

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type ServiceI interface {
	GetNewsByID(ctx context.Context, req *models.GetNewsByIDRequest) (
		resp *models.GetNewsByIDResponse, err error)
	GetAllNews(ctx context.Context, req *models.GetAllNewRequest) (resp *models.GetAllNewsResponse, err error)
	CreateNews(ctx context.Context, req *models.CreateNewsRequest) (resp *models.News, err error)
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
