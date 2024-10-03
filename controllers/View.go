package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func InitView(r *mux.Router) {
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

}

func parseTemplate(path string) (*template.Template, error) {
	return template.ParseFiles("templates/"+path, "templates/navbar.html")
}
