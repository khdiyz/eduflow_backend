package main

import (
	"context"
	"eduflow/cmd/server"
	"eduflow/config"
	"eduflow/internal/handler"
	"eduflow/internal/repository"
	"eduflow/internal/service"
	"eduflow/internal/storage"
	"eduflow/pkg/logger"
	"eduflow/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

// @title EduFlow CRM API
// @version 1.0
// @description API Server for Application

// @host localhost:4040
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	log := logger.GetLogger()

	db, err := utils.SetupPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	minio, err := utils.SetupMinioConnection(cfg, log)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(db, *log)
	newStorage := storage.NewStorage(minio, cfg, log)
	services := service.NewService(*repos, *newStorage, *log, *cfg)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Host, cfg.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Warn("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}
}
