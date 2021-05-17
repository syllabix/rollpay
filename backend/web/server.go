package web

import (
	"net/http"

	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/rollpay/backend/web/docs"
	"github.com/syllabix/rollpay/backend/web/home"
	"github.com/syllabix/rollpay/backend/web/rest"
)

type Server struct {
	*http.Server
}

func NewServer(settings config.ServerSettings, home home.Page, rest rest.Server, docs docs.Server) Server {
	mux := http.NewServeMux()
	mux.Handle("/", root(rest, home))
	mux.Handle(settings.DocsURL, docs)

	return Server{
		Server: &http.Server{
			Addr:         settings.Host + ":" + settings.Port,
			Handler:      mux,
			ReadTimeout:  settings.ReadTimeout,
			WriteTimeout: settings.WriteTimeout,
		},
	}
}

func root(api, home http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			home.ServeHTTP(w, r)
			return
		}
		api.ServeHTTP(w, r)
	})
}
