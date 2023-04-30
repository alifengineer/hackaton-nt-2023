package service

import (
	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/internal/service/match"
	"github.com/dilmurodov/xakaton_nt/internal/service/news"
	"github.com/dilmurodov/xakaton_nt/internal/service/sms"
	"github.com/dilmurodov/xakaton_nt/internal/service/user"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/storage"
)

type ServiceManagerI interface {
	UserService() user.ServiceI
	SmsService() sms.ServiceI
	MatchService() match.ServiceI
	NewsService() news.ServiceI
	// PaymentService() payment.ServiceI
}

type serviceManager struct {
	userService  user.ServiceI
	smsService   sms.ServiceI
	matchService match.ServiceI
	newsService  news.ServiceI
}

func NewServiceManager(cfg config.Config, log logger.LoggerI, strg storage.StorageI) ServiceManagerI {

	userService := user.NewService(cfg, log, strg)
	smsService := sms.NewService(cfg, log, strg)
	matchService := match.NewService(cfg, log, strg)
	newsService := news.NewService(cfg, log, strg)

	return &serviceManager{
		userService:  userService,
		smsService:   smsService,
		matchService: matchService,
		newsService:  newsService,
	}
}

func (s *serviceManager) UserService() user.ServiceI {
	return s.userService
}

func (s *serviceManager) SmsService() sms.ServiceI {
	return s.smsService
}

func (s *serviceManager) MatchService() match.ServiceI {
	return s.matchService
}

func (s *serviceManager) NewsService() news.ServiceI {
	return s.newsService
}
