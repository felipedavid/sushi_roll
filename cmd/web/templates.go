package main

import (
	"bytes"
	"fmt"
	"github.com/felipedavid/sushi_roll/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
)

type templateData struct {
	Games []*models.Game
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
	ts, ok := a.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		a.serverError(w, err)
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(w, "base.tmpl", data)
	if err != nil {
		a.serverError(w, err)
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}
