package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/exvimmer/lets_go/snippetbox/internal/models"
)

// The holding struct for any dynamic data that we want to pass to our HTML
// templates.
type templateData struct {
	CurrentYear int
	Snippets    []*models.Snippet
	Snippet     *models.Snippet
}

// returns a nicely formatted string representation of a time.Time object
func humanDate(t time.Time) string {
	// NOTE: the structure of layout string is very important; check the doc
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // e.g. home.tmpl.html
		/*  NOTE:
		 * The template.FuncMap must be registered with the template set before you
		 * call the ParseFiles() method. This means we have to use template.New()
		 * to create an empty template set, use the Funcs() method to register the
		 * template.FuncMap, and then parse the file as normal.
		 */
		ts, err := template.New(name).Funcs(functions).ParseFiles(
			"./ui/html/base.tmpl.html",
		)
		if err != nil {
			return nil, err
		}

		// add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
