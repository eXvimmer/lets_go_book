package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/exvimmer/lets_go/snippetbox/internal/models"
	"github.com/exvimmer/lets_go/snippetbox/ui"
)

// The holding struct for any dynamic data that we want to pass to our HTML
// templates.
type templateData struct {
	CurrentYear     int
	Snippets        []*models.Snippet
	Snippet         *models.Snippet
	Form            any // to pass validation errors and previously submitted data
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	User            *models.User
}

// returns a nicely formatted UTC string representation of a time.Time object
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // e.g. home.tmpl.html
		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
