package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
}

func getEnv(key string, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return (fallback)
	}

}

func initConfig() Config {
	godotenv.Load()
	return Config{JWTSecret: getEnv("JWTSECRET", "Shrek")}
}

var Env = initConfig()
