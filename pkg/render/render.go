package render

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./Templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
