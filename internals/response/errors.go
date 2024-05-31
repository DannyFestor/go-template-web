package response

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func (rs *Response) Error(w http.ResponseWriter, err error, message string, code int) {
	rs.Logger.Error(err.Error())
	http.Error(w, message, code)
}

func (rs *Response) SomethingWentWrong(w http.ResponseWriter, err error) {
	rs.Error(w, err, "Something went wrong", http.StatusInternalServerError)
}

func (rs *Response) TemplateNotFound(w http.ResponseWriter, err error) {
	rs.Error(w, err, "Template not found", http.StatusInternalServerError)
}

func (rs *Response) NotFound(w http.ResponseWriter, r *http.Request, templateName string, data *templates.Data) {
	rs.Logger.Error("Not Found", "route", r.RequestURI, "rendering template", templateName)
	rs.View(w, r, templateName, data)
}
