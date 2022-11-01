package main

import (
	"bytes"
	"fmt"
	"github.com/felipedavid/sushi_roll/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type templateData struct {
	Games []*models.Game
}

func HumanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateData() *templateData {
	return &templateData{}
}

type templateCache map[string]*template.Template

func newTemplateCache() (templateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[filename] = ts
	}

	return cache, nil
}

func (a *app) render(w http.ResponseWriter, status int, page string, data *templateData) {
	buf := new(bytes.Buffer)

	// If we are in a development environment, don't use the template cache
	if a.env == "development" {
		files := []string{"./ui/html/base.tmpl"}

		partials, err := filepath.Glob("./ui/html/partials/*.tmpl")
		if err != nil {
			a.serverError(w, err)
			return
		}

		files = append(files, partials...)
		files = append(files, fmt.Sprintf("./ui/html/pages/%s", page))

		ts, err := template.ParseFiles(files...)
		if err != nil {
			a.serverError(w, err)
			return
		}

		err = ts.Execute(buf, data)
		if err != nil {
			a.serverError(w, err)
		}

	} else {
		ts, ok := a.templateCache[page]
		if !ok {
			err := fmt.Errorf("the template %s does not exist", page)
			a.serverError(w, err)
			return
		}

		err := ts.ExecuteTemplate(buf, "base", data)
		if err != nil {
			a.serverError(w, err)
		}
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
