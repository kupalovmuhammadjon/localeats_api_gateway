package main

import (
	"api_gateway/api"
	"api_gateway/config"
	"api_gateway/models"
	logger "api_gateway/pkg/logger"


	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log, err := logger.New("debug", "development", cfg.LOG_PATH)
	if err != nil {
		panic(err)
	}


	systemConfig := &models.SystemConfig{
		Config:     cfg,
		Logger:     log,
	}

	router := api.NewRouter(systemConfig)

	err = router.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatal("Server failed to run ", zap.Error(err))
		return
	}
}
