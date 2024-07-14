package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT           string
	AUTH_SERVICE_PORT   string
	ORDER_SERVICE_PORT  string
	ACCESS_SIGNING_KEY  string
	REFRESH_SIGNING_KEY string
	LOG_PATH            string
	APP_PASSWORD        string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error while loading .env file")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":23456"))
	config.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":50053"))
	config.ACCESS_SIGNING_KEY = cast.ToString(coalesce("ACCESS_SIGNING_KEY", "root"))
	config.REFRESH_SIGNING_KEY = cast.ToString(coalesce("REFRESH_SIGNING_KEY", "root"))
	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "areyouinterested.log"))
	config.APP_PASSWORD = cast.ToString(coalesce("APP_PASSWORD", "COMMONMAN"))

	return &config
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
