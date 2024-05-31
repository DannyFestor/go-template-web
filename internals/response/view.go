package response

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func (rs *Response) View(w http.ResponseWriter, rq *http.Request, templName string, data *templates.Data) {
	tmpl, ok := rs.Templates[templName]
	if !ok {
		// TODO: Error Helper Wrapper
		rs.SomethingWentWrong(w, fmt.Errorf("internals/response/view: Template not found: %s", templName))
		return
	}

	executedTemplate := rq.Context().Value("block").(string)
	if executedTemplate == "" {
		executedTemplate = "base"
	}

	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(
		buf,                                // buffer written to
		executedTemplate,                   // rendered block, will be set by isHtmxRequest middleware or overwritten controller
		templates.AddDefaultData(data, rq), // template that will be executed
	)

	if err != nil {
		// TODO: Error Helper Wrapper
		rs.SomethingWentWrong(w, fmt.Errorf("internals/response/view: Error executing template: %s", templName))
		return
	}

	buf.WriteTo(w)
}
