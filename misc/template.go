package misc

import (
	"html/template"
	"strings"
)

type Template struct {
	tmpl *template.Template
}

func LoadTemplate(pattern string) (Template, error) {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return Template{}, err
	}
	return Template{tmpl: tmpl}, nil
}

func (tmpl Template) Render(template string, value interface{}) (string, error) {
	buffer := new(strings.Builder)
	err := tmpl.tmpl.ExecuteTemplate(buffer, template, value)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
