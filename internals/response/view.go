package response

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

func (r *Response) View(w io.Writer, name string, data any) error {
	tmpl, ok := r.Templates[name]
	if !ok {
		// TODO: Error Helper Wrapper
		msg := fmt.Sprintf("Template not found: %s\n", name)
		w.Write([]byte("Something went wrong"))
		return errors.New(msg)
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		// TODO: Error Helper Wrapper
		msg := fmt.Sprintf("Error executing template: %s\nReason: %s", name, err.Error())
		w.Write([]byte("Something went wrong"))
		return errors.New(msg)
	}

	buf.WriteTo(w)
	return nil
}
