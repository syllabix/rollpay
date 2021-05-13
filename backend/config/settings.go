package config

import (
	"flag"
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

	port := flag.String("port", "8080", "set the port the server will listen on")
	host := flag.String("host", "localhost", "set the server host")
	pprofPort := flag.String("pprof-port", "6060", "sets port that pprof will listen on (used for profiling)")
	readtimeout := flag.Duration("read-timeout", time.Second*5, "set the server read timeout")
	writetimeout := flag.Duration("write-timeout", time.Second*5, "set the server write timeout")
	flag.Parse()

	return ServerSettings{
			Port:         *port,
			Host:         *host,
			ProfilerPort: *pprofPort,
			ReadTimeout:  *readtimeout,
			WriteTimeout: *writetimeout,
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
