package main

import (
	"context"

	"github.com/dilmurodov/xakaton_nt/api"
	"github.com/dilmurodov/xakaton_nt/api/handlers"
	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/internal/service"
	"github.com/dilmurodov/xakaton_nt/pkg/logger"
	"github.com/dilmurodov/xakaton_nt/pkg/sms"
	"github.com/dilmurodov/xakaton_nt/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}
	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() { _ = logger.Cleanup(log) }()

	// Connect to postgres
	strg, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer strg.CloseDB()

	svcs := service.NewServiceManager(cfg, log, strg)

	h := handlers.NewHandler(cfg, log, svcs)

	smsDaemon := sms.Daemon{Conf: cfg, Strg: strg}
	go smsDaemon.Init()

	r := api.SetUpRouter(h, cfg)
	if err := r.Run(cfg.HTTPPort); err != nil {
		log.Panic("r.Run", logger.Error(err))
	}
}
