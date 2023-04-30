package sms

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type ServiceI interface {
	SendOtp(ctx context.Context, req *models.SendOtpRequest) (*models.SendOtpResponse, error)
	ConfirmOtp(ctx context.Context, req *models.ConfirmOtpRequest) (err error)
	CreateEmailOtp(ctx context.Context, req *models.Email) (*models.Email, error)
	VerifyEmailOtp(ctx context.Context, req *models.GetEmailOtpByPK) (*models.Email, error)
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
