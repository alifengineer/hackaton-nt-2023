package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/pkg/security"
)

func (self *Service) GetUserByID(ctx context.Context, req *models.GetUserByIDRequest) (resp *models.GetUserByIDResponse, err error) {
	self.log.Info("---GetUserByID--->", logger.Any("req", req))

	resp, err = self.strg.User().GetUserByID(ctx, req)
	if err != nil && err == sql.ErrNoRows {
		self.log.Error("---GetUserByID--->GetUser", logger.Error(err))
		return nil, fmt.Errorf(config.RECORD_NOT_FOUND)
	} else if err != nil {
		self.log.Error("---GetUserByID--->GetUser", logger.Error(err))
		return nil, fmt.Errorf(`%s, %w`, config.SYSTEM_ERROR, err)
	}

	return resp, err
}

func (self *Service) CreateUser(ctx context.Context, req *models.CreateUserRequest) (resp *models.User, err error) {
	self.log.Info("---CreateUser--->", logger.Any("req", req))

	resp, err = self.strg.User().CreateUser(ctx, req)
	if err != nil && err.Error() == config.DUBLICATE_PHONE {
		return nil, fmt.Errorf(config.USER_ALREADY_EXISTS)
	}
	if err != nil {
		self.log.Error("---CreateUser--->", logger.Error(err))
		return nil, fmt.Errorf(`%s, %w`, config.SYSTEM_ERROR, err)
	}

	return resp, nil
}

func (self *Service) GetUserByCredentials(ctx context.Context, req *models.GetByCredentialsRequest) (*models.User, error) {
	self.log.Info("---GetByCredentials--->", logger.Any("req", req))

	if len(req.Password) < 6 {
		err := fmt.Errorf("password must not be less than 6 characters")
		self.log.Error("---GetByCredentials--->", logger.Error(err))
		return nil, err
	}
	user, err := self.strg.User().GetUserPasswordByEmail(ctx, req.Email)
	if err != nil && err == sql.ErrNoRows {
		self.log.Error("---GetByCredentials--->GetUser", logger.Error(err))
		return nil, fmt.Errorf(config.RECORD_NOT_FOUND)
	} else if err != nil {
		self.log.Error("---GetByCredentials--->GetUser", logger.Error(err))
		return nil, fmt.Errorf(`%s, %w`, config.SYSTEM_ERROR, err)
	}

	check, err := security.ComparePassword(user.Password, req.Password)
	if err != nil {
		self.log.Error("---GetByCredentials--->ComparePassword", logger.Error(err))
		return nil, err
	}
	if !check {
		return nil, fmt.Errorf("password or email is incorrect")
	}

	return user, nil
}

func (self *Service) GetUserPasswordByEmail(ctx context.Context, email string) (resp *models.User, err error) {
	self.log.Info("---GetUserPasswordByEmail--->", logger.Any("email", email))

	resp, err = self.strg.User().GetUserPasswordByEmail(ctx, email)
	if err != nil {
		self.log.Error("---GetUserPasswordByEmail--->GetUser", logger.Error(err))
		return nil, err
	}

	return resp, err
}
