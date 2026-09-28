package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/prabhat119/backend-boilerplate/internal/app"
	"github.com/prabhat119/backend-boilerplate/internal/config"
	"github.com/prabhat119/backend-boilerplate/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// 1. Configurations
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprint("Failed to load env config : %v", err))
	}

	// 2. Logger
	logger.InitLogger(cfg.Env)
	defer logger.Log.Sync()

	// 3. Environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 4. Database connection
	db, err := sqlx.Connect("pgx", cfg.DBSource)
	if err != nil {
		logger.Log.Fatal("Database hook crash", zap.Error(err))
	}

	// 5. Dependency Injection
	injection := app.NewInjection(db, cfg)

	// 6. Router Engine
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(gin.Recovery())

	// 7. Route Registration
	app.Routes(router, injection)

	// 8. Listen & Serve
	logger.Log.Info("Server Running Successfully", zap.String("port", cfg.Port))
	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Log.Fatal("Server Crashed", zap.Error(err))
	}
}
