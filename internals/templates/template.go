package templates

import (
	"fmt"
	"html/template"
	"io/fs"
	"regexp"
	"strings"

	"github.com/DannyFestor/go-template-web.git/resources"
)

func NewTemplateCatche() (map[string]*template.Template, error) {
	tmplFuncs := template.FuncMap{}

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

		tmpl, err := template.New("base").Funcs(tmplFuncs).ParseFS(resources.EmbeddedFiles, page)
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

		fmt.Println(name, page)
	}

	return cache, nil
}
