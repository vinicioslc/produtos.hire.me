package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig possui todas configs do .env
type AppConfig struct {
	AppName        string
	AppIsDebug     string
	AppPort        string
	AllowedOrigins string
	APIPrefix      string
	DBPort         string
	DBHost         string
	DBPass         string
	DBUser         string
	DBDatabase     string
	StaticImgHost  string
}

// LoadAppConfig irá instanciar uma nova config baseado no .env da pasta do server
func LoadAppConfig() AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := AppConfig{

		// carrega envs para objeto e retorna
		AppPort:        os.Getenv("APP_PORT"),
		AppName:        os.Getenv("APP_NAME"),
		AppIsDebug:     os.Getenv("APP_ISDEBUG"),
		AllowedOrigins: os.Getenv("API_ALLOWED_ORIGINS"),
		APIPrefix:      os.Getenv("API_PREFIX"),

		DBPort:     os.Getenv("MONGODB_PORT"),
		DBHost:     os.Getenv("MONGODB_HOST"),
		DBPass:     os.Getenv("MONGODB_PASS"),
		DBUser:     os.Getenv("MONGODB_USER"),
		DBDatabase: os.Getenv("MONGODB_DATABASE"),

		StaticImgHost: os.Getenv("STATIC_IMG_HOST"),
	}

	return config
}
