package pages

import (
	"html/template"
	"io"
	"testing"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

func TestRender(t *testing.T) {
	templates := make(map[string]*template.Template)

	tmpl, err := template.New("base").Parse("")
	if err != nil {
		t.Fatalf("%s", err)
	}

	templates["ok"] = tmpl

	app := &config.Application{
		Templates: templates,
	}

	type data struct{}
	var d data
	w := io.Discard

	err = Render(app, w, "ok", d)
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = Render(app, w, "fail", d)
	if err == nil {
		t.Fatalf("Successfully rendered an unavailable template...")
	}
}
