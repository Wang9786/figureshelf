package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"strings"

	"figureshelf-backend/internal/config"
	"figureshelf-backend/internal/database"
	"figureshelf-backend/internal/handler"
	"figureshelf-backend/internal/middleware"
	"figureshelf-backend/internal/repository"
	"figureshelf-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	defer db.Close()

	redisClient := database.NewRedis(cfg.RedisAddr, cfg.RedisURL)
	defer redisClient.Close()

	userRepo := repository.NewUserRepository(db)
	figureRepo := repository.NewFigureRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	dashboardService := service.NewDashboardService(figureRepo, redisClient)
	figureService := service.NewFigureService(figureRepo, dashboardService)

	authHandler := handler.NewAuthHandler(authService)
	figureHandler := handler.NewFigureHandler(figureService)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	router := gin.Default()

	allowedOrigins := strings.Split(cfg.AllowedOrigins, ",")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		dbStatus := "ok"
		if err := db.PingContext(ctx); err != nil {
			dbStatus = "error"
		}

		redisStatus := "ok"
		if err := redisClient.Ping(ctx).Err(); err != nil {
			redisStatus = "error"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"service":  "figureshelf-backend",
			"database": dbStatus,
			"redis":    redisStatus,
		})
	})

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			protected.GET("/me", authHandler.Me)

			figures := protected.Group("/figures")
			{
				figures.POST("", figureHandler.Create)
				figures.GET("", figureHandler.List)

				figures.GET("/upcoming-payments", figureHandler.ListUpcomingPayments)
				figures.GET("/upcoming-releases", figureHandler.ListUpcomingReleases)

				figures.GET("/:id", figureHandler.GetByID)
				figures.PUT("/:id", figureHandler.Update)
				figures.DELETE("/:id", figureHandler.Delete)
			}

			dashboard := protected.Group("/dashboard")
			{
				dashboard.GET("/summary", dashboardHandler.GetSummary)
			}
		}
	}

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}