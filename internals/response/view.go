package response

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func (rs *Response) View(w io.Writer, rq *http.Request, name string, data *templates.Data) error {
	tmpl, ok := rs.Templates[name]
	if !ok {
		// TODO: Error Helper Wrapper
		msg := fmt.Sprintf("Template not found: %s\n", name)
		w.Write([]byte("Something went wrong"))
		return errors.New(msg)
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, templates.AddDefaultData(data, rq))
	if err != nil {
		// TODO: Error Helper Wrapper
		msg := fmt.Sprintf("Error executing template: %s\nReason: %s", name, err.Error())
		w.Write([]byte("Something went wrong"))
		return errors.New(msg)
	}

	buf.WriteTo(w)
	return nil
}
