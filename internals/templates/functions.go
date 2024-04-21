package templates

import (
	"errors"
	"html/template"
)

var functions = template.FuncMap{
	"data": data,
}

func data(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict error")
	}

	data := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}

		data[key] = values[i+1]
	}

	return data, nil
}
