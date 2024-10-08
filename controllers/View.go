package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type View struct {
	Title   string
	Chapter string
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

func SetHeader(header, value string, handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header, value)
		handle.ServeHTTP(w, r)
	})
}
