package handlers

import (
	"net/http"

	"github.com/KanurkarPrateek/Golang_web_app/pkg/render"
	//"text/template"
)

//home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

//about is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
