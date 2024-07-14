package models

import (
	"api_gateway/config"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SystemConfig struct {
	Config     *config.Config
	PostgresDb *sql.DB
	RedisDb    *redis.Client
	Logger     *zap.Logger
}
