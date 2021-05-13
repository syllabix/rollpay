package home

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed service_page.html
var static embed.FS

type Page struct {
	tmpl *template.Template
}

// ServeHTTP handles http requests
func (c Page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.tmpl.Execute(w, c)
}

// NewPage returns a useful instance of a Page
func NewPage() (Page, error) {
	tmpl, err := template.ParseFS(static, "service_page.html")
	if err != nil {
		return Page{}, fmt.Errorf("failed to pase service page html: %w", err)
	}
	return Page{
		tmpl: tmpl,
	}, nil
}
