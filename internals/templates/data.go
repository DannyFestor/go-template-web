package templates

import "net/http"

type Data struct {
	StatusCode int
	Route      string
}

func AddDefaultData(td *Data, r *http.Request) *Data {
	if td == nil {
		td = &Data{
			StatusCode: 200,
			Route:      "/",
		}
	}

	td.Route = r.URL.Path

	return td
}
