package response

import (
	"html/template"
	"io"

	// "io"
	"testing"
)

func TestView(t *testing.T) {
	templates := make(map[string]*template.Template)

	tmpl, err := template.New("base").Parse("")
	if err != nil {
		t.Fatalf("%s", err)
	}

	templates["ok"] = tmpl

	response := &Response{
		Templates: templates,
	}

	type data struct{}
	var d data
	w := io.Discard

	err = response.View(w, "ok", d)
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = response.View(w, "fail", d)
	if err == nil {
		t.Fatalf("Successfully rendered an unavailable template...")
	}
}
