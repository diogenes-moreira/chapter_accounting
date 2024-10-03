package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterIndex(r *mux.Router) {
	r.HandleFunc("/", Index).Methods("GET")
}

func Index(w http.ResponseWriter, r *http.Request) {
	template, err := parseTemplate("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
