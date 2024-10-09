package controllers

import (
	"argentina-tresury/model"
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const movementTypePath = "/api/movement-types"
const movementTypePathId = movementTypePath + "/{id}"

func RegisterMovementTypeRoutesOn(r *mux.Router) {

	r.HandleFunc(movementTypePath, CreateMovementType).Methods("POST")
	r.HandleFunc(movementTypePath, GetMovementTypes).Methods("GET")
	r.HandleFunc("/manual-types", GetManualTypes).Methods("GET")
	r.HandleFunc(movementTypePathId, GetMovementType).Methods("GET")
	r.HandleFunc(movementTypePathId, UpdateMovementType).Methods("PUT")
	r.HandleFunc(movementTypePathId, DeleteMovementType).Methods("DELETE")
}

func GetManualTypes(writer http.ResponseWriter, request *http.Request) {
	movementTypes, err := services.GetManualMovementTypes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(movementTypes)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func GetMovementTypes(writer http.ResponseWriter, request *http.Request) {
	movementTypes, err := services.GetMovementTypes()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(movementTypes)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func GetMovementType(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	movementType, err := services.GetMovementType(uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(movementType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func CreateMovementType(writer http.ResponseWriter, request *http.Request) {
	var movementType model.MovementType
	if err := json.NewDecoder(request.Body).Decode(&movementType); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err := services.CreateMovementType(&movementType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(movementType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateMovementType(writer http.ResponseWriter, request *http.Request) {
	var movementType model.MovementType
	if err := json.NewDecoder(request.Body).Decode(&movementType); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err := services.UpdateMovementType(&movementType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(writer).Encode(movementType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteMovementType(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	err = services.DeleteMovementType(uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
