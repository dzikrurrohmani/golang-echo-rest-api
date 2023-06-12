package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"time"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type TokenConfig struct {
	Secret      string
	Time        uint32
	Memory      uint32
	KeyLen      uint32
	Parallelism uint8
	SignKey     *rsa.PrivateKey
	AccessExp   time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

func isEnvExist(envValue, defaultValue string) string {
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func (c *Config) readConfig() {
	api := isEnvExist(os.Getenv("API_URL"), "SOME_URL")

	dbHost := isEnvExist(os.Getenv("DB_HOST"), "SOME_HOST")
	dbPort := isEnvExist(os.Getenv("DB_PORT"), "SOME_PORT")
	dbUser := isEnvExist(os.Getenv("DB_USER"), "SOME_USER")
	dbPass := isEnvExist(os.Getenv("DB_PASS"), "SOME_PASS")
	dbName := isEnvExist(os.Getenv("DB_NAME"), "SOME_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort)

	c.ApiConfig = ApiConfig{Url: api}
	c.DbConfig = DbConfig{DataSourceName: dsn}

	secret := isEnvExist(os.Getenv("SIGN_KEY"), "SOME_SIGN_KEY")
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	c.TokenConfig = TokenConfig{
		Secret:      secret,
		Time:        0,
		Memory:      64 * 1024,
		KeyLen:      4,
		Parallelism: 32,
		SignKey:     signKey,
		AccessExp:   60 * time.Second,
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
