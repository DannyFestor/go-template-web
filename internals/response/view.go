package response

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func (rs *Response) View(w io.Writer, rq *http.Request, templName string, data *templates.Data) error {
	tmpl, ok := rs.Templates[templName]
	if !ok {
		// TODO: Error Helper Wrapper
		msg := fmt.Sprintf("Template not found: %s\n", templName)
		return errors.New(msg)
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

	err = errors.New("test error")
	if err != nil {
		// TODO: Error Helper Wrapper
		w.Write([]byte("Something went wrong"))
		msg := fmt.Sprintf("Error executing template: %s\nReason: %s\n", templName, err.Error())
		return errors.New(msg)
	}

	buf.WriteTo(w)
	return nil
}
