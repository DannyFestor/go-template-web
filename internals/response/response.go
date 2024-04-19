package response

import (
	"html/template"
	"log/slog"
	"path/filepath"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type Response struct {
	Logger    *slog.Logger
	Templates map[string]*template.Template
}

func NewResponse(logger *slog.Logger) (*Response, error) {
	templateCache, err := templates.NewTemplateCatche(filepath.Join("resources", "views"))
	if err != nil {
		return nil, err
	}

	response := &Response{
		Templates: templateCache,
		Logger:    logger,
	}

	return response, nil
}
