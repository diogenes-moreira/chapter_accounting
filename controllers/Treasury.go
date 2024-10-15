package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

const treasuryPath = "/treasury"

// RegisterTreasuryRoutesOn registers the treasury routes on the given router
func RegisterTreasuryRoutesOn(r *mux.Router) {
	r.HandleFunc(treasuryPath, RenderView).Methods("GET")

}

func RenderView(w http.ResponseWriter, request *http.Request) {
	template, err := parseTemplate("treasury.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = executeTemplate(w, request, template, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
