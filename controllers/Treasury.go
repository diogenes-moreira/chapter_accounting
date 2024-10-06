package controllers

import (
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const treasuryPath = "/treasury"

// RegisterTreasuryRoutesOn registers the treasury routes on the given router
func RegisterTreasuryRoutesOn(r *mux.Router) {
	r.HandleFunc(treasuryPath+"/{id}", GetTreasury).Methods("GET")
	r.HandleFunc(treasuryPath, RenderView).Methods("GET")

}

func RenderView(w http.ResponseWriter, request *http.Request) {
	template, err := parseTemplate("treasury.html")
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

func GetTreasury(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chapter, err := services.GetChapter(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(chapter.TreasurerRollingBalance)

}
