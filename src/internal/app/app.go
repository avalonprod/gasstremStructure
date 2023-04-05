package app

import (
	"context"
	"errors"
	"github.com/avalonprod/gasstrem/src/internal/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avalonprod/gasstrem/src/internal/server"
	"github.com/gin-gonic/gin"

	"github.com/avalonprod/gasstrem/src/pkg/database/mongodb"
	"github.com/avalonprod/gasstrem/src/pkg/logger"

	"github.com/avalonprod/gasstrem/src/internal/config"
)

const configsDir string = "configs"

func Run() {
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Error(err)
		return
	}
	mongoClient, err := mongodb.NewClient(cfg.MongoConfig.URI, cfg.MongoConfig.Username, cfg.MongoConfig.Password)
	if err != nil {
		logger.Error(err)
		return
	}
	mongodb := mongoClient.Database(cfg.MongoConfig.DBName)

	r := gin.Default()

	api := r.Group("/api")

	user.Init(user.UserDeps{

		Database: mongodb,
	}, api)

	srv := server.NewServer(cfg, r)
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server %s\n", err.Error())
		}
	}()

	logger.Infof("Server started on PORT: %s", cfg.HTTPConfig.Port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	const timeout time.Duration = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)

	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}
