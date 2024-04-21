package response

import (
	"html/template"
	"log/slog"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

type Response struct {
	Logger    *slog.Logger
	Templates map[string]*template.Template

	HtmxKey string
}

func NewResponse(logger *slog.Logger) (*Response, error) {
	templateCache, err := templates.NewTemplateCatche()
	if err != nil {
		return nil, err
	}

	response := &Response{
		Templates: templateCache,
		Logger:    logger,

		HtmxKey: "HtmxRequest",
	}

	return response, nil
}
