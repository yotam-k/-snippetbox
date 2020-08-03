package main

import (
	"html/template"
	"path/filepath"
	"time"

	"yotam-snippetbox/pkg/forms"
	"yotam-snippetbox/pkg/models"
)


type templateData struct {
	CurrentYear int
	Form		*forms.Form
	Snippet 	*models.Snippet
	Snippets 	[]*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Helper to avoid duplicate logic when loading the home and showSnippet pages
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Gets all Page templates
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// Look at each page
	for _, page := range pages {
		name := filepath.Base(page)

		// Pass the page into a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add all Layout templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add all Partial templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}