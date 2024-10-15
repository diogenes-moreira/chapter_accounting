package controllers

import (
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const chargeTypePath = "/api/charge-types"
const chargeTypePathWithID = chargeTypePath + "/{id:[0-9]+}"

func RegisterChargeTypesOn(r *mux.Router) {
	r.HandleFunc(chargeTypePath, getChargeTypes).Methods("GET")
	r.HandleFunc(chargeTypePathWithID, getChargeType).Methods("GET")
	r.HandleFunc(chargeTypePath, chargeTypesCreate).Methods("POST")
	r.HandleFunc(chargeTypePathWithID, chargeTypesUpdate).Methods("PUT")
	r.HandleFunc(chargeTypePathWithID, chargeTypesDelete).Methods("DELETE")
}

func getChargeTypes(writer http.ResponseWriter, request *http.Request) {
	types, err := services.GetChargeTypes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(types)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func getChargeType(writer http.ResponseWriter, request *http.Request) {
	//TODO: Implement
}

func chargeTypesCreate(writer http.ResponseWriter, request *http.Request) {
	//TODO: Implement
}

func chargeTypesUpdate(writer http.ResponseWriter, request *http.Request) {
	//TODO: Implement
}

func chargeTypesDelete(writer http.ResponseWriter, request *http.Request) {
	//TODO: Implement
}
