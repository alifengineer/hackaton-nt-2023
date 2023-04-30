package api

import (
	"github.com/dilmurodov/xakaton_nt/api/docs"
	"github.com/dilmurodov/xakaton_nt/api/handlers"
	"github.com/dilmurodov/xakaton_nt/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description This is online banking API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			v1.POST("/image/upload", h.UploadImage)
			// auth
			auth := v1.Group("/auth")
			{
				// авторизация
				auth.POST("/login", h.LoginHandler)
				// регистрация
				auth.POST("/register", h.RegisterHandler)
			}

			// sms := v1.Group("/sms")
			// {
			// 	// отправка смс
			// 	sms.POST("/send", h.SendSmsOTPHandler)
			// 	// проверка смс
			// 	sms.POST("/check", h.VerifySmsOTPHandler)
			// }

			email := v1.Group("/email")
			{
				// отправка email
				email.POST("/send-otp", h.SendCodeToEmail)
				// проверка email
				email.POST("/verify-email/:sms_id/:otp", h.VerifyEmailOtpHandler)
			}
			liga := v1.Group("/league")
			{
				liga.GET("/", h.GetAllLeagues)
				liga.POST("/top_teams", h.GetTopTeamsInLeague)
				liga.GET("/:id", h.GetLeagueByID)
				liga.GET("/:id/seasons", h.GetAllSeasons)
			}
			match := v1.Group("/match")
			{
				match.POST("/", h.CreateMatch)
				match.POST("/tur", h.GetMatchesByTUR)
				match.GET("/", h.GetAllMatches)
				match.GET("/:id", h.GetMatchByID)
			}

			news := v1.Group("/news")
			{
				news.GET("/", h.GetAllNews)
				news.GET("/:id", h.GetNewsByID)
				news.POST("/", h.CreateNews)
				// news.PUT("/:id", h.UpdateNews)
				// news.DELETE("/:id", h.DeleteNews)
			}

			team := v1.Group("/team")
			{
				team.GET("/", h.GetAllTeams)
				team.GET("/:id", h.GetTeamByID)
				// team.POST("/", h.CreateTeam)
				// team.PUT("/:id", h.UpdateTeam)
				// team.DELETE("/:id", h.DeleteTeam)
			}
		}
	}
	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
