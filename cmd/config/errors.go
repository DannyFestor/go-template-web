package config

import "net/http"

func (app *Application) TemplateNotFoundError(path string) {
	app.Logger.Error("Template not found", "path", path)
}

func (app *Application) SomethingWentWrong(w http.ResponseWriter, path string) {
	app.Logger.Error("Template not found", "path", path)
}
