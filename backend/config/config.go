package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBPort         int
	DBUser         string
	DBPassword     string
	DBName         string
	MigrationsPath string
	SchemaPath     string
	ServerAddress  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error parsing DB_PORT")
	}

	return &Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         port,
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
		SchemaPath:     os.Getenv("SCHEMA_PATH"),
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
	}
}
