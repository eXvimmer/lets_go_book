package main

import (
	"html/template"
	"path/filepath"

	"github.com/exvimmer/lets_go/snippetbox/internal/models"
)

// The holding struct for any dynamic data that we want to pass to our HTML
// templates.
type templateData struct {
	Snippets []*models.Snippet
	Snippet  *models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // e.g. home.tmpl.html
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
