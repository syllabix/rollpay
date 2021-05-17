package home

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/syllabix/rollpay/backend/config"
)

//go:embed service_page.html
var static embed.FS

type Page struct {
	tmpl *template.Template

	DocsURL string
}

// ServeHTTP handles http requests
func (c Page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.tmpl.Execute(w, c)
}

// NewPage returns a useful instance of a Page
func NewPage(settings config.ServerSettings) (Page, error) {
	tmpl, err := template.ParseFS(static, "service_page.html")
	if err != nil {
		return Page{}, fmt.Errorf("failed to pase service page html: %w", err)
	}

	return Page{
		tmpl:    tmpl,
		DocsURL: settings.DocsURL,
	}, nil
}
