package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type View struct {
	Title       string
	Chapter     string
	IsAdmin     bool
	IsTreasurer bool
	ReadOnly    bool
}

func InitView(r *mux.Router) {
	fsCss := http.FileServer(http.Dir("./static/css/"))
	fsJs := http.FileServer(http.Dir("./static/js/"))
	r.PathPrefix("/css/").Handler(SetHeader("Content-Type", "text/css", http.StripPrefix("/css/", fsCss)))
	r.PathPrefix("/js/").Handler(SetHeader("Content-Type", "text/javascript", http.StripPrefix("/js/", fsJs)))

}

func parseTemplate(path string) (*template.Template, error) {
	return template.ParseFiles("templates/"+path, "templates/navbar.html")
}

func executeTemplate(w http.ResponseWriter, r *http.Request, template *template.Template, view *View) error {
	if view == nil {
		view = &View{}
	}
	view.IsAdmin = false
	view.IsTreasurer = false
	view.ReadOnly = true
	if view.Title == "" {
		view.Title = "Argentina Treasury"
	}
	view.Chapter = r.Header.Get("Chapter")
	if r.Header.Get("profile") == "admin" {
		view.IsAdmin = true
		view.IsTreasurer = true
		view.ReadOnly = false
	} else if r.Header.Get("profile") == "treasurer" {
		view.IsTreasurer = true
		view.ReadOnly = false
	} else if r.Header.Get("profile") == "principal" {
		view.IsTreasurer = true
	}
	return template.Execute(w, view)
}

func SetHeader(header, value string, handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header, value)
		handle.ServeHTTP(w, r)
	})
}
