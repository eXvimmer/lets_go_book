package main

import "github.com/exvimmer/lets_go/snippetbox/internal/models"

// The holding struct for any dynamic data that we want to pass to our HTML
// templates.
type templateData struct {
	Snippet *models.Snippet
}
