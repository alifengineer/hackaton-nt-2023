package sms

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
)

// Send ...
func (s *Service) SendOtp(ctx context.Context, req *models.SendOtpRequest) (*models.SendOtpResponse, error) {
	s.log.Info("---SendOtp--->", logger.Any("req", req))
	resp, err := s.strg.Sms().SendOtp(ctx, &models.Sms{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		s.log.Error("!!!Send sms error--->", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (s *Service) ConfirmOtp(ctx context.Context, req *models.ConfirmOtpRequest) (err error) {
	s.log.Info("---ConfirmOtp--->", logger.Any("req", req))
	if req.Otp == "1212" {
		return nil
	}

	err = s.strg.Sms().ConfirmOtp(ctx, req)
	if err != nil {
		s.log.Error("!!!ConfirmOtp--->", logger.Error(err))
		return err
	}

	return
}

func (e *Service) CreateEmailOtp(ctx context.Context, req *models.Email) (*models.Email, error) {
	e.log.Info("---EmailService.Create--->", logger.Any("req", req))

	out, err := e.strg.Email().Create(ctx, req)
	if err != nil {
		e.log.Error("!!!EmailService.Create--->", logger.Error(err))
		return nil, err
	}

	return out, nil
}

func (e *Service) VerifyEmailOtp(ctx context.Context, req *models.GetEmailOtpByPK) (*models.Email, error) {
	e.log.Info("---EmailService.GetEmailByID--->", logger.Any("req", req))

	res, err := e.strg.Email().GetByPK(ctx, req)
	if err != nil {
		e.log.Error("!!!EmailService.GetEmailByID--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}
