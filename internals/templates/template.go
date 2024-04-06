package templates

import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

func NewTemplateCatche(pagesDir string) (map[string]*template.Template, error) {
	tmplFuncs := template.FuncMap{}

	cache := make(map[string]*template.Template)

	pages := []string{}
	pagePattern := regexp.MustCompile("^.*.page.tmpl$")

	layouts := []string{}
	layoutPattern := regexp.MustCompile("^.*.layout.tmpl$")

	partials := []string{}
	partialPattern := regexp.MustCompile("^.*.partial.tmpl")

	err := filepath.WalkDir(pagesDir, func(path string, d fs.DirEntry, err error) error {
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

	fmt.Println("Pages", pages)
	fmt.Println("Layouts", layouts)
	fmt.Println("Partials", partials)

	for _, page := range pages {
		// Normalize Page Names to remove path and suffix
		// Makes resources/views/user/dashboard.page.tmpl -> user.dashboard
		// Supports Path Nesting and Same Page Names by using directory as namespace
		name := strings.TrimSuffix(page, ".page.tmpl")
		name = strings.TrimPrefix(name, pagesDir+"/")
		// name := strings.TrimPrefix(page, pagesDir+"/")
		name = strings.ReplaceAll(name, "/", ".")

		tmpl, err := template.New("base").Funcs(tmplFuncs).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFiles(layouts...)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFiles(partials...)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
