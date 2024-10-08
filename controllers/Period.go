package controllers

import (
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const periodPath = "/api/periods"
const periodPathId = periodPath + "/{id}"

func RegisterPeriodRoutesOn(r *mux.Router) {
	r.HandleFunc(periodPath, CreatePeriod).Methods("POST")
	r.HandleFunc(periodPath, GetPeriods).Methods("GET")
	r.HandleFunc(periodPathId, GetPeriod).Methods("GET")
	r.HandleFunc(periodPathId, UpdatePeriod).Methods("PUT")
	r.HandleFunc(periodPathId, DeletePeriod).Methods("DELETE")
}

func CreatePeriod(writer http.ResponseWriter, request *http.Request) {
	//TODO Implement
}

func GetPeriods(writer http.ResponseWriter, request *http.Request) {
	//TODO Implement
}

func GetPeriod(writer http.ResponseWriter, request *http.Request) {
	var periodId = mux.Vars(request)["id"]
	id, err := strconv.Atoi(periodId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	period, err := services.GetPeriod(uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(period)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdatePeriod(writer http.ResponseWriter, request *http.Request) {
	//TODO Implement
}

func DeletePeriod(writer http.ResponseWriter, request *http.Request) {
	//TODO Implement
}
