package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

const PasswordSalt = "diyorbek"

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	cfg := &Config{}
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	//postgres://postgres:2001@localhost:5432/postgres?sslmode=disable
	//cfg.DBUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name) for postgres
	cfg.DBUrl = fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"
					   libDir="C:\\Oracle\\instantclient_21_12\\bin"`, user, pass, host, port, name)
	return cfg
}
