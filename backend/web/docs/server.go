package docs

import (
	"net/http"

	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
	"github.com/syllabix/swagserver/theme"
)

// Server serves the api swagger ui page
type Server struct {
	http.Handler
}

// NewServer returns a server that serves the
// the swagger api documentation page
func NewServer(settings config.ServerSettings) Server {
	handler := swagserver.NewHandler(
		option.SwaggerSpecURL("/swagger.json"),
		option.Theme(theme.Muted),
		option.Path(settings.DocsURL),
	)
	return Server{handler}
}
