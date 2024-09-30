package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template

func InitView(r *mux.Router) {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

}
