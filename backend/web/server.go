package web

import (
	"net/http"

	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/rollpay/backend/web/home"
)

type Server struct {
	*http.Server
}

func NewServer(settings config.ServerSettings, home home.Page) Server {
	mux := http.NewServeMux()
	mux.Handle("/", home)

	return Server{
		Server: &http.Server{
			Addr:         settings.Host + ":" + settings.Port,
			Handler:      mux,
			ReadTimeout:  settings.ReadTimeout,
			WriteTimeout: settings.WriteTimeout,
		},
	}
}
