package response

import (
	"html/template"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type Response struct {
	Templates map[string]*template.Template
}

func NewResponse() (*Response, error) {
	templateCache, err := templates.NewTemplateCatche()
	if err != nil {
		return nil, err
	}

	response := &Response{
		Templates: templateCache,
	}

	return response, nil
}
