package templates

import (
	"errors"
	"html/template"
	"io/fs"
	"regexp"
	"strings"

	"github.com/DannyFestor/go-template-web.git/resources"
)

func data(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict error")
	}

	data := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}

		data[key] = values[i+1]
	}

	return data, nil
}

var functions = template.FuncMap{
	"data": data,
}

func NewTemplateCatche() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages := []string{}
	pagePattern := regexp.MustCompile("^.*.page.tmpl$")

	layouts := []string{}
	layoutPattern := regexp.MustCompile("^.*.layout.tmpl$")

	partials := []string{}
	partialPattern := regexp.MustCompile("^.*.partial.tmpl")

	err := fs.WalkDir(resources.EmbeddedFiles, "views", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || err != nil {
			return err
		}

		if pagePattern.MatchString(path) {
			pages = append(pages, path)
		}

		if layoutPattern.MatchString(path) {
			layouts = append(layouts, path)
		}

		if partialPattern.MatchString(path) {
			partials = append(partials, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Normalize Page Names to remove path and suffix
		// Makes resources/views/user/dashboard.page.tmpl -> user.dashboard
		// Supports Path Nesting and Same Page Names by using directory as namespace
		name := strings.TrimSuffix(page, ".page.tmpl")
		name = strings.TrimPrefix(name, "views/")
		name = strings.ReplaceAll(name, "/", ".")

		tmpl, err := template.New("base").Funcs(functions).ParseFS(resources.EmbeddedFiles, page)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFS(resources.EmbeddedFiles, layouts...)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFS(resources.EmbeddedFiles, partials...)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
