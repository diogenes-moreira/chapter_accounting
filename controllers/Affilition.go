package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

const affiliationPath = "/api/affiliations"
const affiliationPathId = affiliationPath + "/{id}"

func RegisterAffiliationRoutesOn(r *mux.Router) {
	r.HandleFunc("/affiliations/view", GetAffiliationsView).Methods("GET")
	r.HandleFunc(affiliationPath, CreateAffiliation).Methods("POST")
	r.HandleFunc(affiliationPath, GetAffiliations).Methods("GET")
	r.HandleFunc(affiliationPathId, GetAffiliation).Methods("GET")
	r.HandleFunc(affiliationPathId, UpdateAffiliation).Methods("PUT")
	r.HandleFunc(affiliationPathId, DeleteAffiliation).Methods("DELETE")

}

func DeleteAffiliation(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func UpdateAffiliation(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func GetAffiliationsView(w http.ResponseWriter, r *http.Request) {
	templateAffiliations, err := parseTemplate("affiliations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templateAffiliations.Execute(w, View{Title: "Afiliaciones"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateAffiliation(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func GetAffiliations(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func GetAffiliation(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
