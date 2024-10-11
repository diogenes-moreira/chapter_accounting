package controllers

import (
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const affiliationPath = "/api/affiliations"
const affiliationPathId = affiliationPath + "/{id}"

func RegisterAffiliationRoutesOn(r *mux.Router) {
	r.HandleFunc("/affiliations", GetAffiliationsView).Methods("GET")
	r.HandleFunc(affiliationPath+"/payment", CreatePayment).Methods("POST")
	r.HandleFunc(affiliationPath, CreateAffiliation).Methods("POST")
	r.HandleFunc(affiliationPath, GetAffiliations).Methods("GET")
	r.HandleFunc(affiliationPathId, GetAffiliation).Methods("GET")
	r.HandleFunc(affiliationPathId, UpdateAffiliation).Methods("PUT")
	r.HandleFunc(affiliationPathId, DeleteAffiliation).Methods("DELETE")

}

type Payment struct {
	Amount        float64 `json:"amount"`
	Receipt       string  `json:"receipt"`
	Date          ISODate `json:"date"`
	AffiliationId uint    `json:"affiliation_id"`
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	payment := &Payment{}
	if err := json.NewDecoder(r.Body).Decode(payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreatePayment(payment.AffiliationId, payment.Amount, payment.Receipt, payment.Date.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{ "status":"Payment created" }`))
	if err != nil {
		return
	}
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
