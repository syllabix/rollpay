package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/plaid/plaid-go/plaid"
	"github.com/syllabix/rollpay/backend/config/plaidenv"
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
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Load all application settings, from either flags or a .env file
// See the .env.example for available environment variables, or run the
// the application with -h to see what flags are available
func Load() (ServerSettings, PlaidSettings) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("unable to load .env file: reason %v", err))
	}

	return ServerSettings{
			Host:         os.Getenv("HOST"),
			Port:         os.Getenv("PORT"),
			ProfilerPort: os.Getenv("PPROF"),
			ReadTimeout:  getEnvAsDur("READ_TIMEOUT"),
			WriteTimeout: getEnvAsDur("WRITE_TIMEOUT"),
		},
		PlaidSettings{
			ClientID:     os.Getenv("PLAID_CLIENT_ID"),
			Secret:       os.Getenv("PLAID_SECRET"),
			Environment:  plaidenv.Get(os.Getenv("PLAID_ENV")),
			Products:     os.Getenv("PLAID_PRODUCTS"),
			CountryCodes: os.Getenv("PLAID_COUNTRY_CODES"),
			RedirectURI:  os.Getenv("PLAID_REDIRECT_URI"),
		}
}

// getEnvAsDur returns a duration value for and environment variable key
// if the key is empty or value is empty (or not in valid duration syntax)
// this func returns a default of 5 seconds
func getEnvAsDur(key string) time.Duration {
	val := os.Getenv(key)
	dur, err := time.ParseDuration(val)
	if err != nil {
		return time.Duration(time.Second * 5)
	}
	return dur
}
