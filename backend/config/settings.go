package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/plaid/plaid-go/plaid"
	"github.com/syllabix/rollpay/backend/config/plaidenv"
	"github.com/syllabix/rollpay/backend/db"
)

type PlaidSettings struct {
	ClientID     string
	Secret       string
	Environment  plaid.Environment
	Products     string
	CountryCodes string
	RedirectURI  string
}

type ServerSettings struct {
	Host         string
	Port         string
	ProfilerPort string
	DocsURL      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Load all application settings, from either flags or a .env file
// See the .env.example for available environment variables, or run the
// the application with -h to see what flags are available
func Load() (ServerSettings, PlaidSettings, db.Settings) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("unable to load .env file: reason %v", err))
	}

	return ServerSettings{
			Host:         os.Getenv("HOST"),
			Port:         os.Getenv("PORT"),
			DocsURL:      os.Getenv("DOCS_PATH"),
			ProfilerPort: os.Getenv("PPROF"),
			ReadTimeout:  getEnvAsDur("READ_TIMEOUT", time.Minute*2),
			WriteTimeout: getEnvAsDur("WRITE_TIMEOUT", time.Minute*2),
		},
		PlaidSettings{
			ClientID:     os.Getenv("PLAID_CLIENT_ID"),
			Secret:       os.Getenv("PLAID_SECRET"),
			Environment:  plaidenv.Get(os.Getenv("PLAID_ENV")),
			Products:     os.Getenv("PLAID_PRODUCTS"),
			CountryCodes: os.Getenv("PLAID_COUNTRY_CODES"),
			RedirectURI:  os.Getenv("PLAID_REDIRECT_URI"),
		},
		db.Settings{
			DBName:             os.Getenv("DB_NAME"),
			SSLMode:            os.Getenv("SSL_MODE"),
			User:               os.Getenv("DB_USER"),
			Password:           os.Getenv("DB_PASSWORD"),
			Host:               os.Getenv("DB_HOST"),
			Port:               os.Getenv("DB_PORT"),
			MaxConnections:     getEnvAsInt("MAX_CONNS", 5),
			MaxIdleConnections: getEnvAsInt("MAX_IDLE_CONNS", 5),
			MaxConnLifetime:    getEnvAsDur("MAX_CONN_LIFETIME", time.Hour),
		}
}

// getEnvAsDur returns a duration value for and environment variable key
// if the key is empty or value is empty (or not in valid duration syntax)
// this func returns the provided default
func getEnvAsDur(key string, def time.Duration) time.Duration {
	val := os.Getenv(key)
	dur, err := time.ParseDuration(val)
	if err != nil {
		return def
	}
	return dur
}

func getEnvAsInt(key string, def int) int {
	val := os.Getenv(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return num
}
