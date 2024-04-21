package response

import (
	"html/template"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/DannyFestor/go-template-web.git/internals/templates"
)

func TestView(t *testing.T) {
	tmplMap := make(map[string]*template.Template)

	tmpl, err := template.New("base").Parse("")
	if err != nil {
		t.Fatalf("%s", err)
	}

	tmplMap["ok"] = tmpl

	response := &Response{
		Templates: tmplMap,
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	err = response.View(w, r, "ok", &templates.Data{})
	if err != nil {
		t.Fatalf("%s", err)
	}

	err = response.View(w, r, "fail", &templates.Data{})
	if err == nil {
		t.Fatalf("Successfully rendered an unavailable template...")
	}
}
